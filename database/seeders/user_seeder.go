package seeders

import (
	"github.com/goravel/framework/facades"
	"karuhundeveloper.com/gogo/app/models"
)

type UserSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

// Run executes the seeder logic.
func (s *UserSeeder) Run() error {
	var superadmin models.User

	// Password hash
	password, _ := facades.Hash().Make("password")

	// Superamdin email
	email := "superadmin@superadmin.com"

	superadminData := models.User{
		Name: "Super Admin",
		Email: email,
		Password: password,
	}

	// FirstOrCreate user
	return facades.Orm().Query().Where("email", email).FirstOrCreate(&superadmin, superadminData)
}
