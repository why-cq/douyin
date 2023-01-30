package pojo

import "github.com/jinzhu/gorm"

// Follow 用户关系结构，对应用户关系表。
type Follow struct {
	gorm.Model
	UserId     int64
	FollowerId int64
	Cancel     int8
}
