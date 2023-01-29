package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("连接成功")
	defer db.Close()
}
