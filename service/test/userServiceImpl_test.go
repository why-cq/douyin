package test

import (
	"douyin/service"
	"fmt"
	"testing"
)

func TestGetUserByUserName(t *testing.T) {
	USI := service.UserServiceImpl{}

	user := USI.GetUserByUserName("why")
	fmt.Println(user)

}
