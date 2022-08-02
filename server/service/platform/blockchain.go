package platform

import (
	"encoding/hex"
	"gin/global"
	"gin/model/platform"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config/lookup"
	"log"
	"strconv"
	"strings"
)

type BlockChainService struct {
}

func (blockChainService *BlockChainService) GetHighestBlock(ctx *gin.Context) {
	res, err := global.LEDGER_CLIENT.QueryInfo()
	if err != nil {
		ctx.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}
	blockInfo := platform.ChainInfo{CurrentBlockHash: hex.EncodeToString(res.BCI.CurrentBlockHash),
		PreviousBlockHash: hex.EncodeToString(res.BCI.PreviousBlockHash),
		Height:            res.BCI.Height}
	ctx.JSON(200, platform.Response{Code: 200, Data: blockInfo, Message: "success"})
}

func (blockChainService *BlockChainService) GetBlockByNumber(ctx *gin.Context) {
	str := ctx.Param("number")
	intStr, err := strconv.Atoi(str)
	number := uint64(intStr)
	res, err := global.LEDGER_CLIENT.QueryBlock(number)
	if err != nil {
		ctx.JSON(200, platform.Response{Code: 500, Message: err.Error()})
		return
	}
	blockInfo, err := utils.ParsingBlock(res)
	if err != nil {
		ctx.JSON(200, platform.Response{Code: 500, Message: err.Error()})
		return
	}
	ctx.JSON(200, platform.Response{Code: 200, Data: blockInfo, Message: "success"})
}

func (blockChainService BlockChainService) GetBlockList(c *gin.Context) {
	res, err := global.LEDGER_CLIENT.QueryInfo()
	if err != nil {
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}
	blockInfo := platform.ChainInfo{CurrentBlockHash: hex.EncodeToString(res.BCI.CurrentBlockHash),
		PreviousBlockHash: hex.EncodeToString(res.BCI.PreviousBlockHash),
		Height:            res.BCI.Height}
	var blocks []*platform.Block
	for i := uint64(0); i < blockInfo.Height; i++ {
		res, err := global.LEDGER_CLIENT.QueryBlock(i)
		if err != nil {
			c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
			return
		}
		blockInfo, err := utils.ParsingBlock(res)
		if err != nil {
			c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
			return
		}
		blocks = append(blocks, blockInfo)
	}
	c.JSON(200, platform.Response{Code: 200, Data: blocks, Message: "success"})
}

func (blockChainService BlockChainService) GetTransactionList(c *gin.Context) {
	res, err := global.LEDGER_CLIENT.QueryInfo()
	if err != nil {
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}
	blockInfo := platform.ChainInfo{CurrentBlockHash: hex.EncodeToString(res.BCI.CurrentBlockHash),
		PreviousBlockHash: hex.EncodeToString(res.BCI.PreviousBlockHash),
		Height:            res.BCI.Height}
	var transactions []*platform.TransactionInfo
	for i := uint64(0); i < blockInfo.Height; i++ {
		res, err := global.LEDGER_CLIENT.QueryBlock(i)
		if err != nil {
			c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
			return
		}
		blockInfo, err := utils.ParsingBlock(res)
		if err != nil {
			c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
			return
		}
		transactions = append(transactions, blockInfo.Transactions...)
	}
	c.JSON(200, platform.Response{Code: 200, Data: transactions, Message: "success"})
}

func (blockChainService BlockChainService) GetChannelList(c *gin.Context) {
	configBackend, err := global.SDK.Config()
	if err != nil {
		log.Printf("Failed to get config backend from SDK: %s", err)
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}

	targets, err := orgTargetPeers([]string{"Org1"}, configBackend)
	if err != nil {
		log.Printf("Creating peers failed: %s", err)
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}

	channelQueryResponse, err := global.RC.QueryChannels(
		resmgmt.WithTargetEndpoints(targets[0]),
		resmgmt.WithRetry(retry.DefaultResMgmtOpts))

	if err != nil {
		log.Printf("QueryChannels return error: %s", err)
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}

	var channels []string

	for _, channel := range channelQueryResponse.Channels {
		channels = append(channels, channel.ChannelId)
	}

	var channelList []*platform.ChannelCfg

	// 我们需要动态的创建这个LEDGER_CLIENT
	for _, channel := range channels {
		ledgerClient, err := utils.CreateLedgerClient(channel, "Admin")
		cfg, _ := ledgerClient.QueryConfig()
		if err != nil {
			c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
			return
		}

		channelCfg, err := utils.ParsingChannelCfg(cfg)
		if err != nil {
			c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
			return
		}
		channelList = append(channelList, channelCfg)
	}

	c.JSON(200, platform.Response{Code: 200, Data: channelList, Message: "success"})
}

func (blockChainService *BlockChainService) GetBlockByHash(ctx *gin.Context) {
	str := ctx.Param("hash")
	hash, _ := hex.DecodeString(str)
	res, err := global.LEDGER_CLIENT.QueryBlockByHash(hash)
	if err != nil {
		ctx.JSON(200, platform.Response{Code: 500, Message: err.Error()})
		return
	}
	blockInfo, err := utils.ParsingBlock(res)
	if err != nil {
		ctx.JSON(200, platform.Response{Code: 500, Message: err.Error()})
		return
	}
	ctx.JSON(200, platform.Response{Code: 200, Data: blockInfo, Message: "success"})
}

func (blockChainService *BlockChainService) GetBlockByTxHash(c *gin.Context) {
	txHash := c.Param("hash")
	blockInfo, err := global.LEDGER_CLIENT.QueryBlockByTxID(fab.TransactionID(txHash))
	if err != nil {
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}
	_blockInfo, err := utils.ParsingBlock(blockInfo)
	if err != nil {
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(200, platform.Response{Code: 200, Data: _blockInfo, Message: "success"})
}

func (blockChainService *BlockChainService) GetTransactionByTxHash(c *gin.Context) {
	txHash := c.Param("hash")
	processedTransaction, err := global.LEDGER_CLIENT.QueryTransaction(fab.TransactionID(txHash))
	if err != nil {
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}
	payLoad := processedTransaction.TransactionEnvelope.Payload
	transaction, err := utils.UnmarshalTransaction(payLoad)
	if err != nil {
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(200, platform.Response{Code: 200, Data: transaction, Message: "success"})
}

// GetChannelConfig 查询通道配置
func (blockChainService *BlockChainService) GetChannelConfig(c *gin.Context) {
	cfg, err := global.LEDGER_CLIENT.QueryConfig()
	if err != nil {
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
		return
	}

	channelCfg, err := utils.ParsingChannelCfg(cfg)
	if err != nil {
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
		return
	}

	c.JSON(200, platform.Response{Code: 200, Data: channelCfg, Message: "success"})
}

// GetChannelConfigBlock 获取当前指定通道的配置块信息
func (blockChainService *BlockChainService) GetChannelConfigBlock(c *gin.Context) {
	// QueryConfigBlock返回的是common下的Block现象
	blockInfo, err := global.LEDGER_CLIENT.QueryConfigBlock()
	if err != nil {
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
		return
	}

	_blockInfo, err := utils.ParsingBlock(blockInfo)
	if err != nil {
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(200, platform.Response{Code: 200, Data: _blockInfo, Message: "success"})
}

// GetChannels Query channels
func (blockChainService *BlockChainService) GetChannels(c *gin.Context) {

	configBackend, err := global.SDK.Config()
	if err != nil {
		log.Printf("Failed to get config backend from SDK: %s", err)
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}

	targets, err := orgTargetPeers([]string{"Org1"}, configBackend)
	if err != nil {
		log.Printf("Creating peers failed: %s", err)
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}

	channelQueryResponse, err := global.RC.QueryChannels(
		resmgmt.WithTargetEndpoints(targets[0]),
		resmgmt.WithRetry(retry.DefaultResMgmtOpts))

	if err != nil {
		log.Printf("QueryChannels return error: %s", err)
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}

	var channels []string

	for _, channel := range channelQueryResponse.Channels {
		channels = append(channels, channel.ChannelId)
	}
	c.JSON(200, platform.Response{Code: 200, Data: channels, Message: "success"})
}

// GetInstalledCC Query installed chaincode
func (blockChainService *BlockChainService) GetInstalledCC(c *gin.Context) {

	configBackend, err := global.SDK.Config()
	if err != nil {
		log.Printf("Failed to get mainSDK config:%s \n", err)
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}

	targets, err := orgTargetPeers([]string{"Org1"}, configBackend)
	if err != nil {
		log.Printf("Failed to get targets:%s \n", err)
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}
	peer := targets[0]

	var chaincodeInfos []*platform.Chaincode

	installedCC, err := global.RC.LifecycleQueryInstalledCC(resmgmt.WithTargetEndpoints(peer), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	for _, cc := range installedCC {

		chaincodeInfo := &platform.Chaincode{}
		chaincodeInfo.Label = cc.Label
		chaincodeInfo.PackageID = cc.PackageID
		chaincodeInfo.References = cc.References

		chaincodeInfos = append(chaincodeInfos, chaincodeInfo)

	}
	c.JSON(200, platform.Response{Code: 200, Data: chaincodeInfos, Message: "success"})
}

func (blockChainService BlockChainService) GetPeerNameList(c *gin.Context) {

	config, err := global.LEDGER_CLIENT.QueryConfig()
	if err != nil {
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}
	var peers []string
	anchorPeers := config.AnchorPeers()
	for _, peer := range anchorPeers {
		peerSocketStr := peer.Host + ":" + strconv.Itoa(int(peer.Port))
		peers = append(peers, peerSocketStr)
	}

	orderers := config.Orderers()
	for _, orderer := range orderers {
		peers = append(peers, orderer)
	}

	c.JSON(200, platform.Response{Code: 200, Data: peers, Message: "success"})
}

func (blockChainService BlockChainService) GetPeerList(c *gin.Context) {

	config, err := global.LEDGER_CLIENT.QueryConfig()
	if err != nil {
		c.JSON(200, platform.Response{Code: 500, Message: err.Error()})
	}
	var peers []*platform.Peer
	anchorPeers := config.AnchorPeers()
	for _, peer := range anchorPeers {
		peerInfo := &platform.Peer{}
		peerInfo.Org = peer.Org
		peerInfo.Type = "PEER"
		peerSocketStr := peer.Host + ":" + strconv.Itoa(int(peer.Port))
		peerInfo.PeerName = peerSocketStr
		peers = append(peers, peerInfo)
	}

	orderers := config.Orderers()
	for _, orderer := range orderers {
		peerInfo := &platform.Peer{}
		peerInfo.PeerName = orderer
		peerInfo.Org = "OrdererMSP"
		peerInfo.Type = "ORDERER"
		peers = append(peers, peerInfo)
	}

	c.JSON(200, platform.Response{Code: 200, Data: peers, Message: "success"})
}

func orgTargetPeers(orgs []string, configBackend ...core.ConfigBackend) ([]string, error) {

	networkConfig := fab.NetworkConfig{}

	err := lookup.New(configBackend...).UnmarshalKey("organizations", &networkConfig.Organizations)
	if err != nil {
		return nil, err
	}

	var peers []string
	for _, org := range orgs {
		orgConfig, ok := networkConfig.Organizations[strings.ToLower(org)]
		if !ok {
			continue
		}
		peers = append(peers, orgConfig.Peers...)
	}
	return peers, nil
}
