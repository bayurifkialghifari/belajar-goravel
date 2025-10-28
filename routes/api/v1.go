package api

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/http/middleware"

	"karuhundeveloper.com/gogo/app/http/controllers"
	"karuhundeveloper.com/gogo/app/usecase"
)

func V1() {
	facades.Route().Middleware(
		middleware.Throttle("api"),
	).Prefix("api/v1").Group(func (router route.Router) {
		userUseCase := usecase.NewMediaUsecase()
		userController := controllers.NewUserController(userUseCase)
		router.Get("/users/{id}", userController.Show).Name("api.v1.users.show")
		router.Post("/users/create", userController.Create).Name("api.v1.users.create")
	})
}
