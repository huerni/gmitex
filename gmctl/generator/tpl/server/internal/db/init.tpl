package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	{{.imports}}
)

var DB *gorm.DB

func Init(c *config.Config) error {
	var err error
	DB, err = gorm.Open(mysql.Open(c.Mysql.DSN))
	if err != nil {
		return err
	}

	// 根据结构自动建表

    return nil
}
