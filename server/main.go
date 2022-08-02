package main

import (
	"gin/core"
	"gin/global"
	"gin/initialize"
	"go.uber.org/zap"
)

func main() {

	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)

	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	// 初始化所有的fabric相关对象
	initialize.FabricInit()

	core.RunWindowsServer()

}
