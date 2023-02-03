package service

import (
	"douyin/dao"
	"fmt"
)

type UserServiceImpl struct {
}

func (UserServiceImpl) CreatUser(userDao dao.UserDao) bool {
	ok := dao.CreatUser(userDao)
	if ok {
		fmt.Println("创建用户成功")
		return true
	}
	fmt.Println("创建用户失败")
	return false

}

func (UserServiceImpl) GetUserById(id int64) dao.UserDao {
	userDao, err := dao.GetUserByUserid(id)
	if err != nil {
		fmt.Println("获取用户失败")
		return userDao
	}
	fmt.Println("获取用户成功")
	return userDao
}

func (UserServiceImpl) GetUserByUserName(username string) dao.UserDao {
	userDao, err := dao.GetUserByUsername(username)
	if err != nil {
		fmt.Println("获取用户失败")
		return userDao
	}
	fmt.Println("获取用户成功")
	return userDao
}
