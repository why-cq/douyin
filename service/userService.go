package service

import "douyin/dao"

type UserService interface {
	CreatUser(dao.UserDao) bool
	GetUserById(id int64) dao.UserDao
	GetUserByUserName(username string) dao.UserDao
}
