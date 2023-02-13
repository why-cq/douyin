package test

import (
	"douyin/controller"
	"douyin/dao"
	"github.com/gin-gonic/gin"
	"testing"
)

var R *gin.Engine

func init() {
	dao.InitMySQl()
	R = gin.Default()
}

func TestRegister(t *testing.T) {
	R.POST("/douyin/user/register", controller.Register)
	R.Run()
}

func TestLogin(t *testing.T) {
	R.POST("/login", controller.Login)
	R.Run()

}

func TestUserInfo(t *testing.T) {
	R.GET("/user", controller.UserInfo)
	R.Run()

}
