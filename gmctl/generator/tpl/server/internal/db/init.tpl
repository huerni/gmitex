package db

import (
	"github.com/huerni/gmitex/core/logger"
	"gmitest/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	mysqlConfig := config.Cfg.Mysql
	if mysqlConfig.HasConfig() {
		DB, err = gorm.Open(mysql.Open(mysqlConfig.DSN))
		if err != nil {
			logger.Error("mysql初始化失败: ", err)
		}
		logger.Info("Mysql初始化完成")
	}
}
