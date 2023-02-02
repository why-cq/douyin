package dao

import (
	"douyin/pojo"
	"fmt"
)

type LikeDao pojo.Like

func (LikeDao) TableName() string {
	return "likes"
}

// 创建点赞信息
func CreateLike(likeDao LikeDao) error {
	err := DB.Create(&likeDao).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// 根据userID和videoID查找点赞信息

func FindLike(userID int64, videoId int64) (LikeDao, error) {
	likeDao := LikeDao{}
	err := DB.Where("user_id = ? AND video_id = ?", userID, videoId).Find(&likeDao).Error
	if err != nil {
		fmt.Println(err.Error())
		return likeDao, err
	}
	return likeDao, nil
}

// 根据userID和videoID更新点赞信息(取消或者点赞)
func UpdateLike(userID int64, videoId int64) error {
	// 先查询是否存在点赞记录
	likeDao, err := FindLike(userID, videoId)
	if err != nil {
		fmt.Println("点赞记录不存在")
		return err
	}
	// 取消点赞
	if likeDao.Cancel == 0 {
		err := DB.Model(&likeDao).Update("cancel", 1).Error
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	} else {
		// 点赞
		err := DB.Model(&likeDao).Update("cancel", 0).Error
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}

}
