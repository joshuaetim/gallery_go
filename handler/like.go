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

type likeHandler struct {
	repo     domain.LikeRepository
	userRepo domain.UserRepository
}

func NewLikeHandler(repos *domain.Repositories) interfaces.LikeHandler {
	return &likeHandler{
		repo:     repos.Like,
		userRepo: repos.User,
	}
}

func cleanLikesArray(input []model.Like) []model.Like {
	var output []model.Like
	for _, like := range input {
		l := like
		l.Photo = cleanPhotoArray([]model.Photo{l.Photo})[0]
		output = append(output, l)
	}
	return output
}

func (lh *likeHandler) createLike(ctx *gin.Context) (int, error) {
	var like model.Like
	if err := ctx.ShouldBindJSON(&like); err != nil {
		return http.StatusUnprocessableEntity, err
	}
	like.UserID = uint(ctx.GetFloat64("userID"))
	like, err := lh.repo.CreateLike(like)
	if err != nil {
		return http.StatusInternalServerError, errors.New("request failed")
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": "liked"})
	return 0, nil
}

func (lh *likeHandler) getLikes(ctx *gin.Context) (int, error) {
	id := uint(ctx.GetFloat64("userID"))
	likes, err := lh.repo.GetLikesMap(query{"user_id": id})
	if err != nil || len(likes) == 0 {
		return http.StatusNotFound, errors.New("no liked photos yet")
	}
	likes = new(model.Like).PublicArray(likes)
	likes = cleanLikesArray(likes)

	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
	return 0, nil
}

func (lh *likeHandler) getPhotoLikes(ctx *gin.Context) (int, error) {
	id := ctx.Param("photo")
	likes, err := lh.repo.GetLikesMap(query{"photo_id": id})
	if err != nil || len(likes) == 0 {
		return http.StatusNotFound, errors.New("no likes yet")
	}
	likes = new(model.Like).PublicArray(likes)
	likes = cleanLikesArray(likes)

	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
	return 0, nil
}

func (lh *likeHandler) deleteLike(ctx *gin.Context) (int, error) {
	likeId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return http.StatusUnprocessableEntity, err
	}
	likes, err := lh.repo.GetLikesMap(query{"id": uint(likeId)})
	if err != nil || len(likes) == 0 {
		return http.StatusInternalServerError, errors.New("photo not found")
	}

	userID := uint(ctx.GetFloat64("userID"))
	if likes[0].UserID != userID && !isAdmin(ctx, lh.userRepo) {
		return http.StatusUnauthorized, errors.New("like not found")
	}
	err = lh.repo.DeleteLike(likes[0])
	if err != nil {
		return http.StatusInternalServerError, err
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "delete successful"})
	return 0, nil
}

func (lh *likeHandler) CreateLike(ctx *gin.Context) {
	if code, err := lh.createLike(ctx); err != nil {
		errorResponse(ctx, code, err)
	}
}

func (lh *likeHandler) GetLikes(ctx *gin.Context) {
	if code, err := lh.getLikes(ctx); err != nil {
		errorResponse(ctx, code, err)
	}
}

func (lh *likeHandler) GetPhotoLikes(ctx *gin.Context) {
	if code, err := lh.getPhotoLikes(ctx); err != nil {
		errorResponse(ctx, code, err)
	}
}

func (lh *likeHandler) DeleteLike(ctx *gin.Context) {
	if code, err := lh.deleteLike(ctx); err != nil {
		errorResponse(ctx, code, err)
	}
}
