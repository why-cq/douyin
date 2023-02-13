package test

import (
	"douyin/dao"
	"douyin/service"
	"fmt"
	"testing"
)

func init() {
	dao.InitMySQl()
}

func TestGetUserByUserName(t *testing.T) {
	USI := service.UserServiceImpl{}

	user := USI.GetUserByUserName("why")
	fmt.Println(user)

}
