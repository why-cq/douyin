package pojo

import "github.com/jinzhu/gorm"

// Like 表的结构。
type Like struct {
	gorm.Model
	UserId  int64 //点赞用户id
	VideoId int64 //视频id
	Cancel  int   //是否点赞，0为点赞，1为取消赞
}
