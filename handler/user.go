package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshuaetim/akiraka3/domain/model"
	domain "github.com/joshuaetim/akiraka3/domain/repository/interfaces"
	"github.com/joshuaetim/akiraka3/handler/interfaces"

	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	repo domain.UserRepository
}

func NewUserHandler(repos *domain.Repositories) interfaces.UserHandler {
	return &userHandler{
		repo: repos.User,
	}
}

type query map[string]interface{}

func (uh *userHandler) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.Password = hashPassword(user.Password)
	user, err := uh.repo.AddUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": user.PublicUser(),
	})
}

func errorResponse(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, gin.H{
		"error": err.Error(),
	})
}
func (uh *userHandler) signinUser(ctx *gin.Context) (int, error) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return http.StatusUnprocessableEntity, err
	}

	users, err := uh.repo.GetMap(query{"email": user.Email})
	if err != nil || len(users) == 0 {
		return http.StatusUnauthorized, errors.New("details incorrect")
	}
	dbUser := users[0]

	if !comparePassword(dbUser.Password, user.Password) {
		return http.StatusUnauthorized, errors.New("details incorrect")
	}

	token, err := GenerateToken(dbUser.ID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	ctx.JSON(http.StatusOK, gin.H{"data": dbUser.PublicUser(), "token": token})
	return http.StatusOK, nil
}
func (uh *userHandler) SignInUser(ctx *gin.Context) {
	if code, err := uh.signinUser(ctx); err != nil {
		errorResponse(ctx, code, err)
	}
}

func (uh *userHandler) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id parameter",
		})
		return
	}
	users, err := uh.repo.GetMap(query{"id": uint(id)})
	if err != nil || len(users) == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error fetching user",
		})
		return
	}
	user := users[0]

	ctx.JSON(http.StatusOK, gin.H{"data": user.PublicUser()})
}

func (uh *userHandler) UpdateUser(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "binding error: please check your input data: " + err.Error()})
		return
	}

	user.ID = uint(userID)
	user, err := uh.repo.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user, "msg": "user updated"})
}

func (uh *userHandler) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users, err := uh.repo.GetMap(query{"id": uint(id)})
	user := users[0]

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch user"})
		return
	}
	err = uh.repo.DeleteUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "problem deleting user: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "user deleted"})
}

func (uh *userHandler) GetCurrentUser(ctx *gin.Context) {
	userId := ctx.GetFloat64("userID")
	fmt.Println(userId)
	users, err := uh.repo.GetMap(query{"id": uint(userId)})
	user := users[0]

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user.PublicUser(),
	})
}

func (uh *userHandler) GetUsers(ctx *gin.Context) {
	users, err := uh.repo.GetMap(query{"\"1\"": "1"})
	fmt.Println(users)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func hashPassword(password string) string {
	bytePassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}

func comparePassword(hashed, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) == nil
}
