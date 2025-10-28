package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251028051129CreateMedia struct{}

// Signature The unique signature for the migration.
func (r *M20251028051129CreateMedia) Signature() string {
	return "20251028051129_create_media"
}

// Up Run the migrations.
func (r *M20251028051129CreateMedia) Up() error {
	if !facades.Schema().HasTable("media") {
		return facades.Schema().Create("media", func(table schema.Blueprint) {
			table.ID()
			table.String("model_type", 255)
			table.UnsignedBigInteger("model_id")
			table.String("collection_name", 255)
			table.String("name", 255)
			table.String("file_name", 255)
			table.String("mime_type", 100).Nullable()
			table.String("disk", 50)
			table.String("conversions_disk", 50).Nullable()
			table.UnsignedBigInteger("size")
			table.UnsignedInteger("order_column").Nullable()
			table.TimestampsTz()
			table.Index("model_type", "model_id", "collection_name", "order_column")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20251028051129CreateMedia) Down() error {
 	return facades.Schema().DropIfExists("media")
}
