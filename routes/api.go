package routes

import (
	"github.com/goravel/framework/facades"

	"karuhundeveloper.com/gogo/app/http/controllers"
	"karuhundeveloper.com/gogo/app/usecase"
)

func Api() {
	userUseCase := usecase.NewMediaUsecase()
	userController := controllers.NewUserController(userUseCase)
	facades.Route().Get("/users/{id}", userController.Show)
	facades.Route().Post("/users/create", userController.Create)
}
