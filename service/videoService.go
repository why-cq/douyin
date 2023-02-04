package service

import "douyin/dao"

type VideoService interface {
	CreateVideo(video dao.VideoDao) bool
	GetVideoByAuthorId(authorId int64) ([]dao.VideoDao, error)
	GetVideoByVideoId(videoID int64) (dao.VideoDao, error)
}
