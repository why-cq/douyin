package service

import (
	"douyin/dao"
	"fmt"
)

type FollowServiceImpl struct{}

func (FollowServiceImpl) CreatFollowInf(followDao dao.FollowDao) error {
	err := dao.CreatFollowInf(followDao)
	if err != nil {
		fmt.Println("失败")
		return err
	}
	fmt.Println("成功")
	return nil
}

func (FollowServiceImpl) GetFollowersByUserId(userId int64) (int64, error) {
	count, err := dao.GetFollowersByUserId(userId)
	if err != nil {
		fmt.Println("获取粉丝数失败")
		return 0, err
	}
	fmt.Println("获取粉丝数成功")
	return count, err
}

func (FollowServiceImpl) GetFollowInf(userId int64, followId int64) (dao.FollowDao, error) {
	followInf, err := dao.GetFollowInf(userId, followId)
	if err != nil {
		fmt.Println("获取粉丝数失败")
		return followInf, err
	}
	fmt.Println("获取粉丝数成功")
	return followInf, err
}

func (FollowServiceImpl) UpdateFollow(userId int64, followId int64) (bool, error) {
	ok, err := dao.UpdateFollow(userId, followId)
	if err != nil {
		fmt.Println("更新失败")
		return ok, err
	}
	fmt.Println("更新成功")
	return ok, err
}
