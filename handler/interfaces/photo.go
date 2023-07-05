package interfaces

import "github.com/gin-gonic/gin"

type PhotoHandler interface {
	CreatePhoto(*gin.Context)
	GetPhoto(*gin.Context)
	GetAllPhotos(*gin.Context)
	UpdatePhoto(*gin.Context)
	DeletePhoto(*gin.Context)
}
