package test

import (
	"douyin/dao"
	"fmt"
	"testing"
)

func init() {
	dao.InitMySQl()
}

func TestCreatVideo(t *testing.T) {
	ok := dao.CreateVideo(dao.VideoDao{
		AuthorId:      3,
		PlayUrl:       "123",
		CoverUrl:      "123",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    true,
	})
	fmt.Println(ok)
}

func TestGetVideoByAuthorId(t *testing.T) {
	videoList, err := dao.GetVideoByAuthorId(2)
	fmt.Println(videoList)
	fmt.Println(err)
}
