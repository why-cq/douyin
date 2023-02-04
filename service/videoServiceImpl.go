package service

import (
	"douyin/dao"
	"fmt"
)

type VideoServiceImpl struct {
}

func (VideoServiceImpl) CreateVideo(video dao.VideoDao) bool {
	ok := dao.CreateVideo(video)
	if !ok {
		fmt.Println("创建视频失败")
		return false
	}
	fmt.Println("创建视频成功")
	return true

}

func (VideoServiceImpl) GetVideoByAuthorId(authorId int64) ([]dao.VideoDao, error) {
	videoDaos, err := dao.GetVideoByAuthorId(authorId)
	if err != nil {
		fmt.Println("获取视频失败")
		return nil, err
	}
	fmt.Println("获取用户成功")
	return videoDaos, nil
}

func (VideoServiceImpl) GetVideoByVideoId(videoID int64) (dao.VideoDao, error) {
	videoDao, err := dao.GetVideoByVideoId(videoID)
	if err != nil {
		fmt.Println("获取视频失败")
		return videoDao, err
	}
	fmt.Println("获取用户成功")
	return videoDao, nil
}
