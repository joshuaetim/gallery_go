package db

import (
	"fmt"

	"github.com/joshuaetim/akiraka3/domain/model"
	"github.com/joshuaetim/akiraka3/domain/repository/interfaces"
	"gorm.io/gorm"
)

type photoRepo struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) interfaces.PhotoRepository {
	return &photoRepo{
		db: db,
	}
}

func (r *photoRepo) CreatePhoto(photo model.Photo) error {
	return r.db.Create(&photo).Error
}

func (r *photoRepo) GetPhotoMap(query map[string]interface{}) ([]model.Photo, error) {
	var fields []interface{}
	var photo []model.Photo

	queryString := "1=1"
	for k, v := range query {
		queryString = fmt.Sprintf("%s AND %s=?", queryString, k)
		fields = append(fields, v)
	}

	var queryMain []interface{}
	queryMain = append(queryMain, queryString)
	queryMain = append(queryMain, fields...)

	return photo, r.db.Find(&photo, queryMain...).Error
}

func (r *photoRepo) UpdatePhoto(photo model.Photo) (model.Photo, error) {
	return photo, r.db.Model(&photo).Updates(&photo).Error
}

func (r *photoRepo) DeletePhoto(photo model.Photo) error {
	return r.db.Delete(&photo).Error
}
