package main

import (
	"douyin/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	//r.Static("/static", "./public")

	//第一层陆路由地址
	apiRouter := r.Group("/douyin")
	{
		//基础接口
		apiRouter.GET("/feed/", controller.Feed)
		apiRouter.GET("/user/", controller.UserInfo)
		apiRouter.POST("/user/register/", controller.Register)
		apiRouter.POST("/user/login/", controller.Login)
		apiRouter.POST("/publish/action/", func(context *gin.Context) {

		})
		apiRouter.GET("/publish/list/", func(context *gin.Context) {

		})

	}

}
