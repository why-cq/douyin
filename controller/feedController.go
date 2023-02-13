package controller

import (
	"douyin/dao"
	"douyin/pojo"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	pojo.Response
	VideoList []dao.VideoDao `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// 这是假数据
	DemoVideos := []dao.VideoDao{
		{
			Id:            1,
			AuthorId:      2,
			PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
			CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
		},
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response: pojo.Response{
			StatusCode: 0,
			StatusMsg:  "成功",
		},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
