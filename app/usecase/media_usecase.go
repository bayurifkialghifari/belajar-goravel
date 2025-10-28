package usecase

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"karuhundeveloper.com/gogo/app/models/spatie"
)

type MediaUsecase struct {
	// Dependent services
}

func NewMediaUsecase() *MediaUsecase {
	return &MediaUsecase{}
}

func (u *MediaUsecase) UploadMedia(ctx http.Context, fileRequest string, collectionName string, modelType string, modelId uint) (media spatie.Media, err error) {
	// Check if the request has a file
	file, err := ctx.Request().File(fileRequest)

	// If file is not present then do nothing
	if err == nil {
		// Generate a hash name for the file
		filename := file.HashName()
		originalName := file.GetClientOriginalName()

		// Handle file upload
		path, err := file.StoreAs("media", filename)

		// If there was an error storing the file, return the error
		if err != nil {
			facades.Log().Error("Failed to store file: " + err.Error())
			return media, err
		}

		// Get file info
		mime, _ := facades.Storage().MimeType(path)
		size, _ := facades.Storage().Size(path)

		// Create media record
		mediaData := spatie.Media{
			ModelType: 		modelType,
			ModelID: 		modelId,
			CollectionName: collectionName,
			Name: 			originalName,
			FileName: 		filename,
			MimeType: 		mime,
			Disk: 			"local",
			ConversionsDisk: "local",
			Size: 			uint64(size),
			OrderColumn: 	0,
		}
		facades.Orm().Query().Model(&spatie.Media{}).Create(&mediaData)

		return mediaData, nil
	}
	return
}