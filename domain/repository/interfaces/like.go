package interfaces

import "github.com/joshuaetim/gallery_go/domain/model"

type LikeRepository interface {
	CreateLike(model.Like) (model.Like, error)
	GetLikesMap(map[string]interface{}) ([]model.Like, error)
	DeleteLike(model.Like) error
}
