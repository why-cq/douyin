package pojo

import (
	"github.com/jinzhu/gorm"
)

// Comment
type Comment struct {
	gorm.Model
	UserId      int64  //评论用户id
	VideoId     int64  //视频id
	CommentText string //评论内容
	Cancel      int32  //取消评论为1，发布评论为0
}
