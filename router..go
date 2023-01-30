package main

import "github.com/gin-gonic/gin"

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	//r.Static("/static", "./public")

	//第一层陆路由地址
	apiRouter := r.Group("/douyin")
	{
		//基础接口
		apiRouter.GET("/feed/", func(context *gin.Context) {

		})
		apiRouter.GET("/user/", func(context *gin.Context) {

		})
		apiRouter.POST("/user/register/", func(context *gin.Context) {

		})
		apiRouter.POST("/user/login/", func(context *gin.Context) {

		})
		apiRouter.POST("/publish/action/", func(context *gin.Context) {

		})
		apiRouter.GET("/publish/list/", func(context *gin.Context) {

		})

	}

}
