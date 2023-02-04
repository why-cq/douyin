package service

import "douyin/dao"

type CommentService interface {
	CreatComment(commentDao dao.CommentDao) (bool, error)
	GetCommentByVideoId(videoId int64) (int64, error)
	GetCommentById(id int64) (dao.CommentDao, error)
	DeleteComment(id int64) (bool, error)
}
