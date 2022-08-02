package core

import (
	"fmt"
	"gin/global"
	"gin/initialize"
	"go.uber.org/zap"
)

type server interface {
	// ListenAndServe
	// 不定义ListenAndServe接口程序就不阻塞了
	ListenAndServe() error
}

func RunWindowsServer() {
	// 初始化路由
	Router := initialize.Routers()
	// 从全局配置中获取web服务器端口
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	// 使用地址和Router构造一个server
	s := initServer(address, Router)

	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Println("欢迎使用 blockchain-explorer-gin ！")

	global.GVA_LOG.Error(s.ListenAndServe().Error())

}
