package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id     int    `json:"id" gorm:"autoIncrement"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}
