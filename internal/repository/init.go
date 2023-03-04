package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"temp/config"
)

var DB *gorm.DB

func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		config.Conf.Database.User, config.Conf.Database.Password,
		config.Conf.Database.Host, config.Conf.Database.Port,
		config.Conf.Database.Database, config.Conf.Database.Param)
}

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(mysql.Open(getDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败")
	}
	
	if err := DB.AutoMigrate(&User{}); err != nil {
		log.Fatal("表初始化失败")
	}
}
