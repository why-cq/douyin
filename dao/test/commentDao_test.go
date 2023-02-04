package test

import (
	"douyin/dao"
	"fmt"
	"testing"
)

func TestCreatComment(t *testing.T) {
	dao.CreatComment(dao.CommentDao{
		UserId:     3,
		VideoId:    1,
		Content:    "这个视频很好",
		CreateDate: "",
		Cancel:     0,
	})

}

func TestFindCommentByVideoId(t *testing.T) {
	count, err := dao.GetCommentByVideoId(1)
	fmt.Println(count)
	fmt.Println(err)
}

func TestFindCommentById(t *testing.T) {
	commentDao, err := dao.GetCommentById(1)
	fmt.Println(commentDao)
	fmt.Println(err)
}

func TestDeleteComment(t *testing.T) {
	ok, err := dao.DeleteComment(1)
	fmt.Println(ok)
	fmt.Println(err)
}
