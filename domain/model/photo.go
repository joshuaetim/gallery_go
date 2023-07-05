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

func (Photo) TableName() string {
	return "photos"
}
