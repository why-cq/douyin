package dao

import (
	"douyin/pojo"
	"fmt"
)

type VideoDao pojo.Video

// TableName 修改表名映射
func (VideoDao) TableName() string {
	return "videos"
}

// 创建视频信息
func CreateVideo(video VideoDao) bool {
	err := DB.Create(&video).Error
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

// 通过用户Id获取视频列表(一条或者多条视频信息)
func GetVideoByAuthorId(authorId int64) ([]VideoDao, error) {
	var videoList = []VideoDao{}
	err := DB.Where("author_id = ? ", authorId).Find(&videoList).Error
	if err != nil {
		fmt.Printf(err.Error())
		return videoList, err
	}
	return videoList, nil
}

// 通过视频ID查询video
func GetVideoByVideoId(videoID int64) (VideoDao, error) {
	videoDao := VideoDao{Id: videoID}
	err := DB.First(&videoDao).Error
	if err != nil {
		fmt.Printf(err.Error())
		return videoDao, err
	}
	return videoDao, nil
}
