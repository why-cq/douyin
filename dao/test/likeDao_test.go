package test

import (
	"douyin/dao"
	"fmt"
	"testing"
)

//
//func init() {
//	dao.InitMySQl()
//}

func TestCreateLike(t *testing.T) {
	err := dao.CreateLike(dao.LikeDao{
		UserId:  1,
		VideoId: 234,
		Cancel:  1,
	})
	fmt.Println(err)
}

func TestFindLike(t *testing.T) {
	likeDao, err := dao.GetLike(1, 234)
	fmt.Println(likeDao)
	fmt.Println(err)
}

func TestUpdateLike(t *testing.T) {
	err := dao.UpdateLike(1, 234)
	fmt.Println(err)

}
