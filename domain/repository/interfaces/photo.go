package interfaces

import "github.com/joshuaetim/akiraka3/domain/model"

type PhotoRepository interface {
	CreatePhoto(model.Photo) error
	GetPhotoMap(map[string]interface{}) ([]model.Photo, error)
	UpdatePhoto(model.Photo) (model.Photo, error)
	DeletePhoto(model.Photo) error
}
