package platform

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"time"
)

type ChainInfo struct {
	CurrentBlockHash  string
	PreviousBlockHash string
	Height            uint64
	//Endorser          string
	//Status            int32
}

// build Block start

type TransactionInfo struct {
	CreateTime  string //交易创建时间
	ChaincodeID string //交易调用链码ID
	//ChaincodeVersion string   //交易调用链码版本
	Nonce      int64  //随机数
	CreatorMsp string //交易发起者MSPID
	//OUTypes          string   //交易发起者OU分组
	Args []string //输入参数
	//Type             int32    //交易类型
	TxID                string //交易ID
	PayloadProposalHash string
	Endorser            []string
	WriteResult         []string
	ReadResult          string
}

type Block struct {
	BlockHeight       uint64
	CurrentBlockHash  string
	PreviousBlockHash string
	DataHash          string
	TxNumber          int
	Transactions      []*TransactionInfo
}

// end

type BlockHeader struct {
	Number       int
	PreviousHash []byte
	DataHash     []byte
}

type ChannelInfo struct {
	Id          string
	AnchorPeers []string
}

type ChaincodeInfo struct {
	Name        string
	Version     string
	Endorsement string
	Validation  string
}

type Node struct {
	Name  string
	Type  string
	MspID string
}

type Asset struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

type History struct {
	Record    *Asset    `json:"record"`
	TxId      string    `json:"txId"`
	Timestamp time.Time `json:"timestamp"`
}

type Peer struct {
	PeerName string `json:"PeerName"`
	Org      string `json:"Org"`
	Type     string `json:"Type"`
}

type AnchorPeer struct {
	Org  string `json:"Org"`
	Host string `json:"Host"`
	Port int32  `json:"Port"`
}

type ChannelCfg struct {
	ID          string `json:"id"`
	BlockNumber uint64 `json:"blockNumber"`
	AnchorPeers []*AnchorPeer
	Orderers    []string `json:"orderers"`
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Chaincode struct {
	Label      string                           `json:"label"`
	PackageID  string                           `json:"packageID"`
	References map[string][]resmgmt.CCReference `json:"references"`
}
