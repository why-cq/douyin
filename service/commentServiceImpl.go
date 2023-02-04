package service

import "douyin/dao"

type CommentServiceImpl struct{}

func (CommentServiceImpl) CreatComment(commentDao dao.CommentDao) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (CommentServiceImpl) GetCommentByVideoId(videoId int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (CommentServiceImpl) GetCommentById(id int64) (dao.CommentDao, error) {
	//TODO implement me
	panic("implement me")
}

func (CommentServiceImpl) DeleteComment(id int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}
