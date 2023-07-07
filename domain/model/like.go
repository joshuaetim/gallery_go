package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID  uint  `json:"userId"`
	User    User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	PhotoID uint  `json:"photoId"`
	Photo   Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"photo"`
}

func (Like) PublicArray(likes []Like) []Like {
	var des []Like
	for _, like := range likes {
		user := (like.User).PublicUser()
		like.User = user
		des = append(des, like)
	}
	return des
}
