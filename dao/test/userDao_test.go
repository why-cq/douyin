package test

import (
	"douyin/dao"
	"fmt"
	"testing"
)

func init() {
	dao.InitMySQl()
}

func TestCreatUsers(t *testing.T) {
	ok := dao.CreatUser(dao.UserDao{
		UserName: "xixi",
		Password: "123456",
	})
	if ok {
		fmt.Println("插入数据成功")
	} else {
		fmt.Println("插入数据失败")
	}

}

func TestGetUsers(t *testing.T) {

	users, err := dao.GetUsers()
	fmt.Println(users)
	fmt.Println(err)

}

func TestGetUserByid(t *testing.T) {

	user, err := dao.GetUserByUserid(2)
	fmt.Println(user)
	fmt.Println(err)
}

func TestGetUserByUsername(t *testing.T) {
	user, err := dao.GetUserByUsername("why")
	fmt.Println(user)
	fmt.Println(err)

}
