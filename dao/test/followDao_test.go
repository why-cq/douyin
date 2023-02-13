package test

import (
	"douyin/dao"
	"fmt"
	"testing"
)

//func init() {
//	dao.InitMySQl()
//}

func TestCreatFollowInf(t *testing.T) {
	err := dao.CreatFollowInf(dao.FollowDao{
		UserId:     2,
		FollowerId: 6,
		Cancel:     0,
	})
	fmt.Println(err)
}

func TestGetFansByUserId(t *testing.T) {
	count, err := dao.GetFansByUserId(1)
	fmt.Println(count)
	fmt.Println(err)

}

func TestGetFollowersByUserId(t *testing.T) {
	count, err := dao.GetFansByUserId(2)
	fmt.Println(count)
	fmt.Println(err)
}

func TestFindFollowInf(t *testing.T) {
	followDao, err := dao.GetFollowInf(1, 8)
	fmt.Println(followDao)
	fmt.Println(err)
}

func TestUpdateFollow(t *testing.T) {
	ok, err := dao.UpdateFollow(1, 8)
	fmt.Println(ok)
	fmt.Println(err)

}
