package models

import (
	"github.com/goravel/framework/database/orm"
	"karuhundeveloper.com/gogo/app/models/spatie"
)

type User struct {
	orm.Model
	Name     string
	Email    string
	Password string
	File 	 *spatie.Media `gorm:"polymorphic:Model"`
}
