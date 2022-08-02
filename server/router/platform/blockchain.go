package platform

import (
	"gin/service"
	"github.com/gin-gonic/gin"
)

type BlockchainRouter struct{}

func (s *BlockchainRouter) InitBlockchainRouter(Router *gin.RouterGroup) {
	blockchainRouter := Router.Group("blockchain")
	platformServiceGroup := service.GroupApp.PlatformServiceGroup
	{
		blockchainRouter.GET("getHighestBlock", platformServiceGroup.GetHighestBlock)                     // 获取最高区块信息
		blockchainRouter.GET("getBlockByNumber/:number", platformServiceGroup.GetBlockByNumber)           // 通过区块序号获取区块信息
		blockchainRouter.GET("getBlockByHash/:hash", platformServiceGroup.GetBlockByHash)                 // 通过区块hash获取区块信息
		blockchainRouter.GET("getBlockByTxHash/:hash", platformServiceGroup.GetBlockByTxHash)             /// 通过交易id获取所属区块信息
		blockchainRouter.GET("getTransactionByTxHash/:hash", platformServiceGroup.GetTransactionByTxHash) /// 通过交易id获取交易信息
		blockchainRouter.GET("getChannelConfig", platformServiceGroup.GetChannelConfig)                   /// 获取通道配置信息
		blockchainRouter.GET("getChannelConfigBlock", platformServiceGroup.GetChannelConfigBlock)         /// 获取当前指定通道配置块信息
		blockchainRouter.GET("getChannels", platformServiceGroup.GetChannels)                             /// 获取所有通道
		blockchainRouter.GET("getInstalledCC", platformServiceGroup.GetInstalledCC)                       /// 获取安装的链码
		blockchainRouter.GET("getPeerNameList", platformServiceGroup.GetPeerNameList)                     /// 获取所有的节点
		blockchainRouter.GET("getPeerList", platformServiceGroup.GetPeerList)                             /// 获取所有的节点详细信息
		blockchainRouter.GET("getTransactionList", platformServiceGroup.GetTransactionList)               /// 获取所有的交易
		blockchainRouter.GET("getBlockList", platformServiceGroup.GetBlockList)                           /// 获取所有区块
		blockchainRouter.GET("getChannelList", platformServiceGroup.GetChannelList)                       /// 获取所有通道详细信息
	}
}
