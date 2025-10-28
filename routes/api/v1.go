package api

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	systemMiddleware "github.com/goravel/framework/http/middleware"

	v1 "karuhundeveloper.com/gogo/app/http/controllers/v1"
	"karuhundeveloper.com/gogo/app/http/middleware"
	"karuhundeveloper.com/gogo/app/usecase"
)

func V1() {
	facades.Route().Middleware(
		systemMiddleware.Throttle("api"),
		middleware.AuthJwt(),
	).Prefix("api/v1").Group(func (router route.Router) {
		userUseCase := usecase.NewMediaUsecase()
		userController := v1.NewUserController(userUseCase)

		// User routes
		router.Get("/users/{id}", userController.Show)
		router.Post("/users/create", userController.Create)
	})
}
