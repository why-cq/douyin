package service

import "douyin/dao"

type LikeService interface {
	CreateLike(likeDao dao.LikeDao) error
	GetLike(userID int64, videoId int64) (dao.LikeDao, error)
	UpdateLike(userID int64, videoId int64) error
}
