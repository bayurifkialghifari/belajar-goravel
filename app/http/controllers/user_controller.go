package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"karuhundeveloper.com/gogo/app/models"
	"karuhundeveloper.com/gogo/app/usecase"
)

type UserController struct {
	mediaUseCase *usecase.MediaUsecase
}

func NewUserController(mediaUseCase *usecase.MediaUsecase) *UserController {
	return &UserController{
		mediaUseCase: mediaUseCase,
	}
}

func (r *UserController) Show(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}

func (r *UserController) Create(ctx http.Context) http.Response {
	// Create new user
	password, _ := facades.Hash().Make("secretpassword")
	user := &models.User{
		Name:     "John Doe",
		Email:    "jhondoe@gmail.com",
		Password: password,
	}
	facades.Orm().Query().Model(&models.User{}).Create(user)

	// Get insert id
	var userModel models.User
	facades.Orm().Query().OrderByDesc("id").Limit(1).Get(&userModel)
	// Upload profile picture
	media, _ := r.mediaUseCase.UploadMedia(ctx, "file", "pp", "user", userModel.ID)

	// if err != nil {
	// 	return ctx.Response().Json(http.StatusBadRequest, http.Json{
	// 		"message": "Failed to upload profile picture",
	// 		"error":   err.Error(),
	// 	})
	// }

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Profile picture uploaded successfully",
		"user":	 userModel,
		"media":   media,
	})
}
