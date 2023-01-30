package pojo

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Video struct {
	gorm.Model
	AuthorId    int64
	PlayUrl     string `json:"play_url"`
	CoverUrl    string `json:"cover_url"`
	PublishTime time.Time
	Title       string `json:"title"`
}
