package rules

import (
	"mime/multipart"

	"github.com/goravel/framework/contracts/validation"
)

type MaxFileSize struct {
}

// Signature The name of the rule.
func (receiver *MaxFileSize) Signature() string {
	return "max_file_size"
}

// Passes Determine if the validation rule passes.
func (receiver *MaxFileSize) Passes(data validation.Data, val any, options ...any) bool {
	// val is *multipart.FileHeader
	fileHeader, ok := val.(*multipart.FileHeader)

	// options[0] is max size in KB
	println(fileHeader)

	if !ok {
		return false
	}

	return false
}

// Message Get the validation error message.
func (receiver *MaxFileSize) Message() string {
	return "The :attribute must not be greater than :max kilobytes."
}
