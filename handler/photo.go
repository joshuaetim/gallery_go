package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshuaetim/akiraka3/domain/model"
	domain "github.com/joshuaetim/akiraka3/domain/repository/interfaces"
	"github.com/joshuaetim/akiraka3/handler/interfaces"
)

type photoHandler struct {
	repo     domain.PhotoRepository
	userRepo domain.UserRepository
}

func NewPhotoHandler(repo domain.PhotoRepository, userRepo domain.UserRepository) interfaces.PhotoHandler {
	return &photoHandler{
		repo:     repo,
		userRepo: userRepo,
	}
}

func emptyStrings(texts ...string) bool {
	for _, s := range texts {
		if s == "" {
			return true
		}
	}
	return false
}

func (ph *photoHandler) createPhoto(ctx *gin.Context) (int, error) {
	var photo model.Photo
	if err := ctx.ShouldBindJSON(&photo); err != nil {
		return http.StatusUnprocessableEntity, err
	}
	if emptyStrings(photo.Link, photo.Title) {
		return http.StatusUnprocessableEntity, errors.New("please fill all fields")
	}
	userId := ctx.GetFloat64("userID")
	photo.UserID = uint(userId)

	err := ph.repo.CreatePhoto(photo)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": "create photo successful",
	})

	return 0, nil
}
func (ph *photoHandler) CreatePhoto(ctx *gin.Context) {
	if code, err := ph.createPhoto(ctx); err != nil {
		errorResponse(ctx, code, err)
	}
}

func (ph *photoHandler) getPhoto(ctx *gin.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return http.StatusUnprocessableEntity, err
	}
	photos, err := ph.repo.GetPhotoMap(query{"id": uint(id)})
	if err != nil || len(photos) == 0 {
		return http.StatusInternalServerError, errors.New("photo record not found")
	}
	photos = new(model.Photo).PublicArray(photos)

	ctx.JSON(http.StatusOK, gin.H{
		"photo": photos[0],
	})
	return 0, nil
}
func (ph *photoHandler) GetPhoto(ctx *gin.Context) {
	if code, err := ph.getPhoto(ctx); err != nil {
		errorResponse(ctx, code, err)
	}
}

func (ph *photoHandler) getAllPhotos(ctx *gin.Context) (int, error) {
	photos, err := ph.repo.GetPhotoMap(query{})
	if err != nil {
		return http.StatusUnprocessableEntity, err
	}
	photos = new(model.Photo).PublicArray(photos)

	ctx.JSON(http.StatusOK, gin.H{
		"photos": photos,
	})
	return 0, nil
}
func (ph *photoHandler) GetAllPhotos(ctx *gin.Context) {
	if code, err := ph.getAllPhotos(ctx); err != nil {
		errorResponse(ctx, code, err)
	}
}

func (ph *photoHandler) updatePhoto(ctx *gin.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return http.StatusUnprocessableEntity, err
	}

	photo := model.Photo{}
	if err := ctx.ShouldBindJSON(&photo); err != nil {
		return http.StatusUnprocessableEntity, err
	}
	photo.ID = uint(id)

	// photoModel, err := ph.repo.GetPhotoMap(query{"id": photo.ID})
	// if err != nil || len(photoModel) == 0 {
	// 	return http.StatusInternalServerError, err
	// }
	// fmt.Println(photoModel[0].User)

	// currentUser, _ := ph.userRepo.GetMap(query{"id": uint(ctx.GetFloat64("id"))})
	// if currentUser[0].Role != "admin" {

	// }
	// if photoModel[0].UserID != uint(ctx.GetFloat64("userID")) {
	// 	return http.StatusUnauthorized, errors.New("photo not found")
	// }

	_, err = ph.repo.UpdatePhoto(photo)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "photo updated",
	})
	return 0, nil
}
func (ph *photoHandler) UpdatePhoto(ctx *gin.Context) {
	if code, err := ph.updatePhoto(ctx); err != nil {
		errorResponse(ctx, code, err)
	}
}

func isAdmin(ctx *gin.Context, userRepo domain.UserRepository) bool {
	id := ctx.GetFloat64("userID")
	users, err := userRepo.GetMap(query{"id": uint(id)})
	if err != nil || len(users) == 0 {
		panic(err)
	}
	if users[0].Role == "admin" {
		return true
	}
	return false
}

func (ph *photoHandler) deletePhoto(ctx *gin.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return http.StatusUnprocessableEntity, err
	}
	photos, err := ph.repo.GetPhotoMap(query{"id": uint(id)})
	if err != nil || len(photos) == 0 {
		return http.StatusInternalServerError, errors.New("photo not found")
	}
	photo := photos[0]

	// check if user owns post
	if photo.UserID != uint(ctx.GetFloat64("userID")) {
		if !isAdmin(ctx, ph.userRepo) {
			return http.StatusNotFound, errors.New("photo not found")
		}
	}

	err = ph.repo.DeletePhoto(photo)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": "photo deleted",
	})
	return 0, nil
}
func (ph *photoHandler) DeletePhoto(ctx *gin.Context) {
	if code, err := ph.deletePhoto(ctx); err != nil {
		errorResponse(ctx, code, err)
	}
}
