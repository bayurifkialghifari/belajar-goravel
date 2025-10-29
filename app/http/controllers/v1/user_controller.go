package v1

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/errors"
	"github.com/goravel/framework/facades"
	"karuhundeveloper.com/gogo/app/http/requests/v1/user"
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
	var user models.User

	// Get user id
	id := ctx.Request().RouteInt64("id")

	// Find user by id
	err := facades.Orm().Query().With("File").Where("id", id).Omit("password").FirstOrFail(&user)

	if errors.Is(err, errors.OrmRecordNotFound) {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "User not found",
		})
	}


	return ctx.Response().Success().Json(http.Json{
		"message": "User found",
		"user":    user,
	})
}

func (r *UserController) Create(ctx http.Context) http.Response {
	var userCreateRequest user.UserCreate

	errors, err := ctx.Request().ValidateRequest(&userCreateRequest)

	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "Validation error",
			"error":   err.Error(),
		})
	}

	if errors != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "Validation failed",
			"errors":  errors.All(),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Validation passed",
		"name":	userCreateRequest.Name,
		"email":	userCreateRequest.Email,
		"password":	userCreateRequest.Password,
		"confirm_password":	userCreateRequest.ConfirmPassword,
	})

	// Create new user facades.Hash().Make("secretpassword")
	// user := &models.User{
	// 	Name:     "John Doe",
	// 	Email:    "jhondoe@gmail.com",
	// 	Password: password,
	// }
	// facades.Orm().Query().Model(&models.User{}).Create(user)

	// // Get insert id
	// var userModel models.User
	// facades.Orm().Query().OrderByDesc("id").Limit(1).Get(&userModel)
	// // Upload profile picture
	// media, _ := r.mediaUseCase.UploadMedia(ctx, "file", "pp", "user", userModel.ID)

	// // if err != nil {
	// // 	return ctx.Response().Json(http.StatusBadRequest, http.Json{
	// // 		"message": "Failed to upload profile picture",
	// // 		"error":   err.Error(),
	// // 	})
	// // }

	// return ctx.Response().Json(http.StatusOK, http.Json{
	// 	"message": "Profile picture uploaded successfully",
	// 	"user":	 userModel,
	// 	"media":   media,
	// })
	// password, _ := facades.Hash().Make("secretpassword")
	// user := &models.User{
	// 	Name:     "John Doe",
	// 	Email:    "jhondoe@gmail.com",
	// 	Password: password,
	// }
	// facades.Orm().Query().Model(&models.User{}).Create(user)

	// // Get insert id
	// var userModel models.User
	// facades.Orm().Query().OrderByDesc("id").Limit(1).Get(&userModel)
	// // Upload profile picture
	// media, _ := r.mediaUseCase.UploadMedia(ctx, "file", "pp", "user", userModel.ID)

	// // if err != nil {
	// // 	return ctx.Response().Json(http.StatusBadRequest, http.Json{
	// // 		"message": "Failed to upload profile picture",
	// // 		"error":   err.Error(),
	// // 	})
	// // }

	// return ctx.Response().Json(http.StatusOK, http.Json{
	// 	"message": "Profile picture uploaded successfully",
	// 	"user":	 userModel,
	// 	"media":   media,
	// })
}
