package service

import (
	"douyin/dao"
	"fmt"
)

type LikeServiceImpl struct{}

func (LikeServiceImpl) CreateLike(likeDao dao.LikeDao) error {
	err := dao.CreateLike(likeDao)
	if err != nil {
		fmt.Println("创建like信息失败")
		return err
	}
	fmt.Println("创建like信息成功")
	return nil
}

func (LikeServiceImpl) GetLike(userID int64, videoId int64) (dao.LikeDao, error) {
	like, err := dao.GetLike(userID, videoId)
	if err != nil {
		fmt.Println("查询like信息sb")
		return like, err
	}
	fmt.Println("查询like信息成功")
	return like, nil
}

func (LikeServiceImpl) UpdateLike(userID int64, videoId int64) error {
	err := dao.UpdateLike(userID, videoId)
	if err != nil {
		fmt.Println("失败")
		return err
	}
	fmt.Println("成功")
	return nil

}
