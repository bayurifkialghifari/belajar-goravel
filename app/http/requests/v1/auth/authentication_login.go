package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AuthenticationLogin struct {
	Email   string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *AuthenticationLogin) Authorize(ctx http.Context) error {
	return nil
}

func (r *AuthenticationLogin) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthenticationLogin) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"email":    "required|email",
		"password": "required|string|min_len:8",
	}
}

func (r *AuthenticationLogin) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthenticationLogin) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthenticationLogin) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
