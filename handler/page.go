package handler

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuaetim/akiraka3/domain/repository/db"
)

type JsonReq struct {
	Name string `json:"name"`
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "pong"})
}

func JSONRequest(ctx *gin.Context) {
	var req JsonReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request: "+err.Error())
		return
	}
	// fmt.Println(req.Name)
	ctx.JSON(http.StatusOK, gin.H{
		"data": req.Name,
	})
}

func CheckAuth(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")
	dbConn := db.DB()
	ur := db.NewUserRepository(dbConn)
	users, err := ur.GetMap(query{"id": uint(userID)})
	user := users[0]
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
	}
	ctx.JSON(http.StatusOK, user.PublicUser())
}
