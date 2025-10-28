package spatie

import (
	"github.com/goravel/framework/database/orm"
)
type Media struct {
	orm.Model
	ModelType      string
	ModelID        uint
	CollectionName string
	Name           string
	FileName       string
	MimeType       string
	Disk           string
	ConversionsDisk string
	Size           uint64
	OrderColumn    uint32
}

func (r *Media) TableName() string {
	return "media"
}
