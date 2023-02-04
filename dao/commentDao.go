package dao

import (
	"douyin/pojo"
	"fmt"
)

type CommentDao pojo.Comment

func (CommentDao) TableName() string {
	return "comments"
}

// 创建评论信息
func CreatComment(commentDao CommentDao) (bool, error) {
	err := DB.Create(&commentDao).Error
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

// 通过视频Id查询评论数量
func GetCommentByVideoId(videoId int64) (int64, error) {
	var count int64
	err := DB.Model(CommentDao{}).Where("video_id = ?", videoId).Count(&count).Error
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return count, nil
}

// 通过ID查询评论
func GetCommentById(id int64) (CommentDao, error) {
	var commentDao CommentDao
	err := DB.Where("id = ?", id).Find(&commentDao).Error
	if err != nil {
		fmt.Println(err)
		return commentDao, err
	}
	return commentDao, nil
}

// 通过主键删除(恢复)评论
func DeleteComment(id int64) (bool, error) {
	commentDao, err := GetCommentById(id)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	//评论
	if commentDao.Cancel == 1 {
		err := DB.Model(&commentDao).Update("cancel", 0).Error
		if err != nil {
			fmt.Println(err)
			return false, err
		}
		return true, nil
	} else {
		//删除评论
		err := DB.Model(&commentDao).Update("cancel", 1).Error
		if err != nil {
			fmt.Println(err)
			return false, err
		}
		return true, nil
	}

	return true, nil
}
