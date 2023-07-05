package model

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Link   string `json:"link"`
	Title  string `json:"title"`
	Story  string `json:"story"`
	UserID uint   `json:"userId"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user"`
}

// type photosArray []Photo
func (p Photo) PublicArray(photos []Photo) []Photo {
	var des []Photo
	for _, photo := range photos {
		user := (photo.User).PublicUser()
		photo.User = user
		des = append(des, photo)
	}
	return des
}
func (Photo) TableName() string {
	return "photos"
}
