package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuaetim/akiraka3/domain/repository/db"
	"github.com/joshuaetim/akiraka3/handler"
	"github.com/joshuaetim/akiraka3/middleware"
)

func RunAPI(address string) error {
	db := db.DB()
	repos := InitRepos(db)
	userHandler := handler.NewUserHandler(repos)
	photoHandler := handler.NewPhotoHandler(repos)
	likeHandler := handler.NewLikeHandler(repos)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Gallery")
	})
	apiRoutes := r.Group("/api")

	apiRoutes.GET("/checkauth", middleware.AuthorizeJWT(), handler.CheckAuth)

	userRoutes := apiRoutes.Group("/auth")
	userRoutes.POST("/register", userHandler.CreateUser)
	userRoutes.POST("/login", userHandler.SignInUser)

	adminRoutes := apiRoutes.Group("/admin/users", middleware.AuthorizeJWT())
	adminRoutes.GET("/", userHandler.GetUsers)
	adminRoutes.GET("/:id", userHandler.GetUser)
	adminRoutes.PUT("/", userHandler.UpdateUser)
	// adminRoutes.GET("/", userHandler.GetCurrentUser)

	photoRoutes := apiRoutes.Group("/photos")
	photoRoutes.GET("/", photoHandler.GetAllPhotos)
	photoRoutes.GET("/:id", photoHandler.GetPhoto)

	photoProtectedRoutes := photoRoutes.Group("", middleware.AuthorizeJWT())
	photoProtectedRoutes.POST("/", photoHandler.CreatePhoto)
	photoProtectedRoutes.PATCH("/:id", photoHandler.UpdatePhoto)
	photoProtectedRoutes.DELETE("/:id", photoHandler.DeletePhoto)

	likeRoutes := apiRoutes.Group("/likes", middleware.AuthorizeJWT())
	likeRoutes.POST("/", likeHandler.CreateLike)
	likeRoutes.GET("/", likeHandler.GetLikes)
	likeRoutes.GET("/:photo", likeHandler.GetPhotoLikes)
	likeRoutes.DELETE("/:id", likeHandler.DeleteLike)

	return r.Run(address)
}
