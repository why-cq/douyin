package service

import (
	"douyin/dao"
	"fmt"
)

type CommentServiceImpl struct{}

func (CommentServiceImpl) CreatComment(commentDao dao.CommentDao) (bool, error) {
	ok, err := dao.CreatComment(commentDao)
	if err != nil {
		fmt.Println("创建评论失败")
		return ok, err
	}
	fmt.Println("创建成功")
	return ok, nil
}

func (CommentServiceImpl) GetCommentByVideoId(videoId int64) (int64, error) {
	count, err := dao.GetCommentByVideoId(videoId)
	if err != nil {
		fmt.Println("查询失败")
		return 0, err
	}
	fmt.Println("查询成功")
	return count, nil
}

func (CommentServiceImpl) GetCommentById(id int64) (dao.CommentDao, error) {
	commentDao, err := dao.GetCommentById(id)
	if err != nil {
		fmt.Println("查询失败")
		return commentDao, err
	}
	fmt.Println("查询成功")
	return commentDao, nil
}

func (CommentServiceImpl) DeleteComment(id int64) (bool, error) {
	ok, err := dao.DeleteComment(id)
	if err != nil {
		fmt.Println("更新评论失败")
		return ok, err
	}
	fmt.Println("更新成功")
	return ok, nil
}
