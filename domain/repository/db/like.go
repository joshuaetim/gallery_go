package db

import (
	"fmt"

	"github.com/joshuaetim/akiraka3/domain/model"
	"github.com/joshuaetim/akiraka3/domain/repository/interfaces"
	"gorm.io/gorm"
)

type likeRepo struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) interfaces.LikeRepository {
	return &likeRepo{
		db: db,
	}
}

func (r *likeRepo) CreateLike(like model.Like) (model.Like, error) {
	return like, r.db.Create(&like).Error
}

func (r *likeRepo) GetLikesMap(query map[string]interface{}) ([]model.Like, error) {
	var fields []interface{}
	var like []model.Like

	queryString := "1=1"
	for k, v := range query {
		queryString = fmt.Sprintf("%s AND likes.%s=?", queryString, k)
		fields = append(fields, v)
	}

	var queryMain []interface{}
	queryMain = append(queryMain, queryString)
	queryMain = append(queryMain, fields...)

	return like, r.db.Order("created_at desc").Joins("User").Joins("Photo").Find(&like, queryMain...).Error
}

func (r *likeRepo) DeleteLike(like model.Like) error {
	return r.db.Delete(&like).Error
}
