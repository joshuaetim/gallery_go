package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID  uint  `json:"userId"`
	User    User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	PhotoID uint  `json:"photoId"`
	Photo   Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"photo"`
}
