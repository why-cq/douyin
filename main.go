package main

import (
	"douyin/dao"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//1.连接数据库
	err := dao.InitMySQl()
	if err != nil {
		log.Printf("初始化数据库失败, err:%v\n", err)
		return
	}

	// 程序退出关闭数据库连接
	defer dao.Close()

	// 模型绑定(运行后数据库中的表自动创建出来)
	//dao.DB.AutoMigrate(
	//	&pojo.User{},
	//	&pojo.Comment{},
	//	&pojo.Like{},
	//	&pojo.Follow{},
	//	&pojo.Video{},
	//)

	// 注册路由
	r := gin.Default()
	initRouter(r)

	// 默认端口8080
	r.Run()

}
