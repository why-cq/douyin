package test

import (
	"douyin/controller"
	"github.com/gin-gonic/gin"
	"testing"
)

var R *gin.Engine

func init() {
	R = gin.Default()
}

func TestRegister(t *testing.T) {
	R.POST("/register", controller.Register)
	R.Run()
}
