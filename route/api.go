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
	userHandler := handler.NewUserHandler(repos.user)

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

	return r.Run(address)
}
