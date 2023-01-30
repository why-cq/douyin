package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name     string
	PassWord string
}

func main() {
	// 连接数据库
	db, _ := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&loc=Local")

	defer db.Close()

	// 把模型和数据库中的表对应起来
	db.AutoMigrate(&User{})

	//创建实例对象
	//u1 := User{Name: "ccc", PassWord: "123456"}
	//u2 := User{Name: "ddd", PassWord: "123456"}
	//db.Create(&u1)
	//db.Create(&u2)

	// 查询数据库
	//var user1 User
	user1 := new(User)
	db.First(&user1) //查询主键的第一条记录
	fmt.Printf("user:%#v\n", user1)
	fmt.Printf("-------------------")
	// 查询所有记录

	users := new([]User)
	db.Find(users)

}
