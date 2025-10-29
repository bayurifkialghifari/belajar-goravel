package api

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	systemMiddleware "github.com/goravel/framework/http/middleware"

	v1 "karuhundeveloper.com/gogo/app/http/controllers/v1"
	"karuhundeveloper.com/gogo/app/http/controllers/v1/auth"
	"karuhundeveloper.com/gogo/app/http/middleware"
	"karuhundeveloper.com/gogo/app/usecase"
)

func V1() {
	// Authentication Controller
	authenticationController := auth.NewAuthenticationController()

	// User Controller
	userUseCase := usecase.NewMediaUsecase()
	userController := v1.NewUserController(userUseCase)

	// v1 Without Auth Middleware
	facades.Route().Middleware(
		systemMiddleware.Throttle("api"),
	).Prefix("api/v1").Group(func (router route.Router) {

		// Login route
		router.Post("/auth/login", authenticationController.Login)
	})

	// V1 With JWT Auth Middleware
	facades.Route().Middleware(
		systemMiddleware.Throttle("api"),
		middleware.AuthJwt(),
	).Prefix("api/v1").Group(func (router route.Router) {
		// Logout route
		router.Post("/auth/logout", authenticationController.Logout)

		// User routes
		router.Get("/users/{id}", userController.Show)
		router.Post("/users/create", userController.Create)
	})
}
