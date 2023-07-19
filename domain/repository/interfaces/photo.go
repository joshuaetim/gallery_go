package interfaces

import "github.com/joshuaetim/gallery_go/domain/model"

type PhotoRepository interface {
	CreatePhoto(model.Photo) error
	GetPhotoMap(map[string]interface{}) ([]model.Photo, error)
	SearchPhotos(string) ([]model.Photo, error)
	UpdatePhoto(model.Photo) (model.Photo, error)
	DeletePhoto(model.Photo) error
}
