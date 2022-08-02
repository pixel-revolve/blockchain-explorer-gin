package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"gin/model/platform"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric/protoutil"
	"github.com/pkg/errors"
	"log"
	"regexp"
	"time"
)

func GetEnvelopeFromBlock(data []byte) (*common.Envelope, error) {
	var err error
	env := &common.Envelope{}
	if err = proto.Unmarshal(data, env); err != nil {
		return nil, errors.Wrap(err, "error unmarshaling Envelope")
	}

	return env, nil
}

func GetTransactionResult(data []byte, result *platform.TransactionInfo) *platform.TransactionInfo {
	action, _ := protoutil.GetActionFromEnvelope(data)
	flySnowRegexp := regexp.MustCompile(".{.*?}")
	res := flySnowRegexp.FindStringSubmatch(string(action.GetResults()))
	result.WriteResult = res
	if res == nil {
		result.WriteResult = make([]string, 0)
	}
	if action != nil {
		result.ReadResult = string(action.GetResponse().Payload)
	}
	return result
}

func UnmarshalTransaction(payloadRaw []byte) (*platform.TransactionInfo, error) {
	result := &platform.TransactionInfo{}
	//获取payload
	payload, err := protoutil.UnmarshalPayload(payloadRaw)
	if err != nil {
		log.Printf("payload")
		return nil, err
	}
	channelHeader, err := protoutil.UnmarshalChannelHeader(payload.Header.ChannelHeader)
	if err != nil {
		log.Printf("channelHeader")
		return nil, err
	}
	signHeader, err := protoutil.UnmarshalSignatureHeader(payload.Header.SignatureHeader)
	if err != nil {
		log.Printf("SignatureHeader")
		return nil, err
	}
	identity, err := protoutil.UnmarshalSerializedIdentity(signHeader.GetCreator())
	if err != nil {
		log.Printf("SerializedIdentity")
		return nil, err
	}
	//block, _ := pem.Decode(identity.GetIdBytes())
	//if block == nil {
	//	log.Printf("区块为空？")
	//	return nil, fmt.Errorf("identity could not be decoded from credential")
	//}
	//cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Printf("error")
		return nil, fmt.Errorf("failed to parse certificate: %s", err)
	}
	//uname := cert.Subject.CommonName
	//outypes := cert.Subject.OrganizationalUnit
	tx, err := protoutil.UnmarshalTransaction(payload.Data)
	if err != nil {
		return nil, err
	}
	chaincodeActionPayload, err := protoutil.UnmarshalChaincodeActionPayload(tx.Actions[0].Payload)
	if err != nil {
		return nil, err
	}
	proposalPayload, err := protoutil.UnmarshalChaincodeProposalPayload(chaincodeActionPayload.ChaincodeProposalPayload)
	if err != nil {
		return nil, err
	}
	chaincodeInvocationSpec, err := protoutil.UnmarshalChaincodeInvocationSpec(proposalPayload.Input)
	if err != nil {
		return nil, err
	}
	if chaincodeInvocationSpec.ChaincodeSpec != nil {
		result.ChaincodeID = chaincodeInvocationSpec.ChaincodeSpec.ChaincodeId.Name
		//result.ChaincodeVersion = chaincodeInvocationSpec.ChaincodeSpec.ChaincodeId.Version
		var args []string
		for _, v := range chaincodeInvocationSpec.ChaincodeSpec.Input.Args {
			args = append(args, string(v))
		}
		result.Args = args
	}
	buf := bytes.NewBuffer(signHeader.GetNonce())
	var nonce int64
	binary.Read(buf, binary.BigEndian, &nonce)
	prp, _ := protoutil.UnmarshalProposalResponsePayload(chaincodeActionPayload.Action.GetProposalResponsePayload())
	endorsers := chaincodeActionPayload.Action.GetEndorsements()
	for _, endorser := range endorsers {
		regexp := regexp.MustCompile("(.)*MSP")
		if endorser.GetEndorser() == nil {
			continue
		}
		result.Endorser = append(result.Endorser, regexp.FindStringSubmatch(string(endorser.GetEndorser()))[0])
	}
	//result.Args = args 在上面给了
	result.Nonce = nonce
	//result.Type = channelHeader.GetType()
	result.TxID = channelHeader.GetTxId()
	result.CreatorMsp = identity.GetMspid()
	//result.Name = uname
	//result.OUTypes = outypes[0]
	result.PayloadProposalHash = hex.EncodeToString(prp.GetProposalHash())
	result.CreateTime = time.Unix(channelHeader.Timestamp.Seconds, 0).Format("2006-01-02 15:04:05")
	return result, nil
}

func MakeTxSlice(blockData [][]byte) ([]*platform.TransactionInfo, error) {
	var txSlice []*platform.TransactionInfo
	for _, data := range blockData {
		envLope, err := GetEnvelopeFromBlock(data)
		if err != nil {
			return nil, err
		}
		txInfo, err := UnmarshalTransaction(envLope.Payload)
		if err != nil {
			return nil, err
		}
		txSlice = append(txSlice, txInfo)
	}
	return txSlice, nil
}

func GetBlockHash(height uint64, previousHash, dataHash []byte) (string, error) {
	data := struct {
		Number       uint64 `json:"number"`
		PreviousHash []byte `json:"previous_hash"`
		DataHash     []byte `json:"data_hash"`
	}{
		Number:       height,
		PreviousHash: previousHash,
		DataHash:     dataHash,
	}
	blockHeaderDer, err := asn1.Marshal(data)
	if err != nil {
		return err.Error(), nil
	}
	hash := sha256.New()
	hash.Write(blockHeaderDer)
	bytes := hash.Sum(nil)
	currentBlockHash := hex.EncodeToString(bytes)
	return currentBlockHash, nil
}

// ParsingBlock 从common的Block数据结构反序列化成自定义区块数据结构
func ParsingBlock(res *common.Block) (*platform.Block, error) {
	header := platform.BlockHeader{Number: int(res.Header.Number), PreviousHash: res.Header.PreviousHash, DataHash: res.Header.DataHash}
	blockData := res.GetData()
	data, _ := asn1.Marshal(header)
	m := sha256.New()
	m.Write(data)
	currentHash := hex.EncodeToString(m.Sum(nil))
	previousHash := hex.EncodeToString(res.Header.PreviousHash)
	dataHash := hex.EncodeToString(res.Header.DataHash)
	height := header.Number
	txNumber := len(blockData.GetData())
	// transactions := []*entity.TransactionInfo{}
	var transactions []*platform.TransactionInfo
	for _, value := range blockData.GetData() {
		env, err := GetEnvelopeFromBlock(value)
		if err != nil {
			return nil, err
		}
		result, err := UnmarshalTransaction(env.Payload)
		if err != nil {
			return nil, err
		}
		result = GetTransactionResult(value, result)
		transactions = append(transactions, result)
	}
	blockInfo := platform.Block{
		BlockHeight:       uint64(height),
		PreviousBlockHash: previousHash,
		CurrentBlockHash:  currentHash,
		DataHash:          dataHash,
		TxNumber:          txNumber,
		Transactions:      transactions,
	}
	return &blockInfo, nil
}

//ParsingChannelCfg 将fab包中的数据结构转换成我们自己的数据结构
func ParsingChannelCfg(cfg fab.ChannelCfg) (*platform.ChannelCfg, error) {
	id := cfg.ID()
	blockNumber := cfg.BlockNumber()
	orderers := cfg.Orderers()

	// 因为我们希望返回的就是唯一的对象
	var anchorPeers []*platform.AnchorPeer
	for _, value := range cfg.AnchorPeers() {
		anchorPeer := new(platform.AnchorPeer)
		*anchorPeer = platform.AnchorPeer{
			Org:  value.Org,
			Host: value.Host,
			Port: value.Port,
		}
		anchorPeers = append(anchorPeers, anchorPeer)
	}
	channelCfg := new(platform.ChannelCfg)
	*channelCfg = platform.ChannelCfg{
		ID:          id,
		BlockNumber: blockNumber,
		AnchorPeers: anchorPeers,
		Orderers:    orderers,
	}
	return channelCfg, nil
}
