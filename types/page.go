package types

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Id   int `gorm:"primaryKey"`
	Name string
}
type Daily struct {
	gorm.Model
	Id  int   `gorm:"primaryKey"`
	Tag []int `gorm:"type:integer[]"`
}
