package global

import (
	"gin/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var(

	// GVA_LOG zap日志配置
	GVA_LOG *zap.Logger

	// GVA_CONFIG 服务端配置
	GVA_CONFIG config.Server

	// GVA_VP 全局viper
	GVA_VP *viper.Viper

	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB

	lock       sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
