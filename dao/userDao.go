package dao

import (
	"douyin/pojo"
	"fmt"
	"log"
)

/*	type UserDao struct {
	Id       int64
	Name     string
	Password string
}*/
// UserDao 对应数据库users表
type UserDao pojo.User

// TableName 修改表名映射
func (UserDao) TableName() string {
	return "users"
}

// 创建一个用户
func CreatUser(user UserDao) bool {
	err := DB.Create(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

// 获取所有的user用户
func GetUsers() ([]UserDao, error) {
	//定义一个切片来存数据
	users := []UserDao{}
	//在数据库中查找
	err := DB.Find(&users).Error
	if err != nil {
		log.Printf(err.Error())
		return users, err
	}
	return users, nil
}

// 通过用户id获取用户
func GetUserByUserid(id int64) (UserDao, error) {
	user := UserDao{Id: id}
	err := DB.First(&user).Error
	if err != nil {
		log.Printf(err.Error())
		return user, err
	}
	return user, nil
}

// 通过用户名获取用户
func GetUserByUsername(username string) (UserDao, error) {
	UserDao := UserDao{}
	err := DB.Where("name = ?", username).First(&UserDao).Error
	if err != nil {
		log.Printf(err.Error())
		return UserDao, err
	}
	return UserDao, nil
}
