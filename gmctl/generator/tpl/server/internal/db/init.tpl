package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	{{.imports}}
)

var DB *gorm.DB

func Init(c *config.Config) {
	var err error
	DB, err = gorm.Open(mysql.Open(c.Mysql.DSN))
	if err != nil {
		fmt.Println(err)
	}

	// 根据结构自动建表

}
