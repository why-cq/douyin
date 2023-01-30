package routers

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//r.Static("/static", "./public") 告诉模板文件引用的静态文件在哪里去找
	//r.LoadHTMLGlob("templates/*") 模板文件的位置

	// 设置一个路由组
	dyGroup := r.Group("/douyin")
	{
		//基础接口
		dyGroup.GET("/feed/", func(context *gin.Context) {

		})
		dyGroup.GET("/user/", func(context *gin.Context) {

		})
		dyGroup.POST("/user/register/", func(context *gin.Context) {

		})
		dyGroup.POST("/user/login/", func(context *gin.Context) {

		})
		dyGroup.POST("/publish/action/", func(context *gin.Context) {

		})
		dyGroup.GET("/publish/list/", func(context *gin.Context) {

		})

	}
	return r
}
