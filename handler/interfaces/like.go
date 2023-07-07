package interfaces

import "github.com/gin-gonic/gin"

type LikeHandler interface {
	CreateLike(*gin.Context)
	GetLikes(*gin.Context)
	GetPhotoLikes(*gin.Context)
	DeleteLike(*gin.Context)
}
