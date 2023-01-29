package pojo

import "github.com/jinzhu/gorm"

type user struct {
	gorm.Model
	Name     string
	PassWord string
}
