package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	DB *gorm.DB
)

func InitMySQl() (err error) {
	dsn := "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接数据库
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		// 连接失败
		log.Panicln("err:", err.Error())
		return
	}
	log.Println("连接成功")
	// 测试连通性
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
