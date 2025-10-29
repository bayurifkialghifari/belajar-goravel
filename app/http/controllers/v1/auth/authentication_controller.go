package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/errors"
	"github.com/goravel/framework/facades"
	"karuhundeveloper.com/gogo/app/http/requests/v1/auth"
	"karuhundeveloper.com/gogo/app/http/responses"
	"karuhundeveloper.com/gogo/app/models"
)

type AuthenticationController struct {
	// Dependent services
}

func NewAuthenticationController() *AuthenticationController {
	return &AuthenticationController{
		// Inject services
	}
}

func (r *AuthenticationController) Login(ctx http.Context) http.Response {
	var user models.User
	var loginRequest auth.AuthenticationLogin

	// Validate request
	validationErrors, err := ctx.Request().ValidateRequest(&loginRequest)

	// Handle validation errors
	if err != nil {
		return ctx.Response().Json(http.StatusUnprocessableEntity, responses.ErrorResponse("Validation Error", err.Error()))
	}

	if validationErrors != nil {
		return ctx.Response().Json(http.StatusUnprocessableEntity, responses.ErrorValidationResponse("Validation Error", validationErrors.All()))
	}

	// Authentication logic
	err = facades.Orm().Query().Where("email", loginRequest.Email).FirstOrFail(&user)

	// Check if user exists
	if errors.Is(err, errors.OrmRecordNotFound) {
		return ctx.Response().Json(http.StatusUnauthorized, responses.ErrorResponse("Invalid credentials", ""))
	}

	// Verify password
	if !facades.Hash().Check(loginRequest.Password, user.Password) {
		return ctx.Response().Json(http.StatusUnauthorized, responses.ErrorResponse("Invalid credentials", ""))
	}

	// Generate token
	token, err := facades.Auth(ctx).Login(&user)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, responses.ErrorResponse("Could not generate token", err.Error()))
	}

	// Return success response with token
	return ctx.Response().Json(http.StatusCreated, responses.SuccessResponse("Login successful", http.Json{
		"token": token,
		"user":  user,
	}))
}

func (r *AuthenticationController) Logout(ctx http.Context) http.Response {
	err := facades.Auth(ctx).Logout()
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, responses.ErrorResponse("Could not logout", err.Error()))
	}

	return ctx.Response().Json(http.StatusOK, responses.SuccessResponse("Logout successful", nil))
}