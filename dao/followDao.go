package dao

import (
	"douyin/pojo"
	"fmt"
)

type FollowDao pojo.Follow

func (FollowDao) TableName() string {
	return "follows"
}

// 创建关注信息
func CreatFollowInf(followDao FollowDao) error {
	err := DB.Create(&followDao).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil

}

// 给定userId查询他关注了多少人
func GetFollowersByUserId(userId int64) (int64, error) {
	// 用来记录关注的人数
	var count int64
	err := DB.Model(FollowDao{}).
		Where("user_id = ? AND cancel = ?", userId, 0).
		Count(&count).Error
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return count, nil

}

// 通过userId和FollowId查询关注信息
func GetFollowInf(userId int64, followId int64) (FollowDao, error) {
	followDao := FollowDao{UserId: userId, FollowerId: followId}
	err := DB.Where("user_id = ? AND follower_id = ?", userId, followId).
		Find(&followDao).
		Error
	if err != nil {
		fmt.Println(err)
		return followDao, err
	}
	return followDao, nil
}

// 关注和取消关注(通过userId和FollowId)
func UpdateFollow(userId int64, followId int64) (bool, error) {
	followDao, err := GetFollowInf(userId, followId)
	if err != nil {
		fmt.Println("关注信息不存在")
		return false, err
	}
	//关注
	if followDao.Cancel == 1 {
		err := DB.Model(&followDao).Update("cancel", 0).Error
		if err != nil {
			fmt.Println(err)
			return false, err
		}
		return true, nil
	} else {
		//取消关注
		err := DB.Model(&followDao).Update("cancel", 1).Error
		if err != nil {
			fmt.Println(err)
			return false, err
		}
		return true, nil
	}

	return true, nil
}
