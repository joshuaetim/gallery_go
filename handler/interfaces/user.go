package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {
	CreateUser(*gin.Context)
	SignInUser(*gin.Context)
	GetUser(*gin.Context)
	GetCurrentUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
	GetUsers(*gin.Context)
}
