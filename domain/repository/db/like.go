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
	var queryString string
	var fields []interface{}
	var like []model.Like
	for k, v := range query {
		if queryString != "" {
			queryString = " " + queryString + " AND "
		}
		queryString = fmt.Sprintf("%s%s = ?", queryString, k)
		fields = append(fields, v)
	}
	// fields[0]
	var queryMain []interface{}
	queryMain = append(queryMain, queryString)
	queryMain = append(queryMain, fields...)

	return like, r.db.Find(&like, queryMain...).Error
}

func (r *likeRepo) DeleteLike(like model.Like) error {
	return r.db.Delete(&like).Error
}
