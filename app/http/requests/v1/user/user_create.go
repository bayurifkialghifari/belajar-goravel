package user

import (
	"mime/multipart"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UserCreate struct {
	Name string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password"`
	Image *multipart.FileHeader `form:"image" json:"image"`
}

func (r *UserCreate) Authorize(ctx http.Context) error {
	return nil
}

func (r *UserCreate) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserCreate) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":             "required|string|max_len:255",
		"email":            "required|email|unique:users,email",
		"password":         "required|string|min_len:8|eq_field:confirm_password",
		"confirm_password": "required|string|min_len:8|eq_field:password",
		"image":            "required_with:image|image|max_file_size:2048",
	}
}

func (r *UserCreate) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserCreate) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserCreate) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
