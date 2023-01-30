package main

import (
	"douyin/pojo"
	"douyin/routers"
	"douyin/sql"
	"log"
)

func main() {
	//1.连接数据库
	err := sql.InitMySQl()
	if err != nil {
		log.Printf("初始化数据库失败, err:%v\n", err)
		return
	}
	// 程序退出关闭数据库连接
	//defer sql.Close()

	// 模型绑定
	sql.DB.AutoMigrate(
		&pojo.User{},
		&pojo.Comment{},
		&pojo.Like{},
		&pojo.Follow{},
		&pojo.Video{},
	)
	//// 注册路由
	r := routers.SetupRouter()
	//// 默认端口8080
	r.Run("9000")

}
