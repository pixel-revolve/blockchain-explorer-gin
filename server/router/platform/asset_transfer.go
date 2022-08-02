package platform

import (
	"gin/service"
	"github.com/gin-gonic/gin"
)

type AssetTransferRouter struct{}

func (s *AssetTransferRouter) InitAssetTransferRouter(Router *gin.RouterGroup) {
	assetTransferRouter := Router.Group("assetTransfer")
	platformServiceGroup := service.GroupApp.PlatformServiceGroup
	{
		assetTransferRouter.POST("putAsset", platformServiceGroup.PutAsset)                 // 设置资产
		assetTransferRouter.GET("getHistoryById", platformServiceGroup.ReadAssetHistory)    // 通过交易id溯源
		assetTransferRouter.GET("readAsset", platformServiceGroup.ReadAsset)                // 读取资产信息
		assetTransferRouter.POST("searchAsset", platformServiceGroup.QueryAssetByAssetData) /// 通过资产数据查询资产
	}
}
