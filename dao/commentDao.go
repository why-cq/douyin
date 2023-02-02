package dao

import "douyin/pojo"

type CommentDao pojo.Comment

func (CommentDao) TableName() string {
	return "comments"
}
