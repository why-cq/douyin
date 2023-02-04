package service

import "douyin/dao"

type FollowService interface {
	CreatFollowInf(followDao dao.FollowDao) error
	GetFollowersByUserId(userId int64) (int64, error)
	GetFollowInf(userId int64, followId int64) (dao.FollowDao, error)
	UpdateFollow(userId int64, followId int64) (bool, error)
}
