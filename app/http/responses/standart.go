package responses

func ErrorResponse(message string, err string) map[string]any {
	return map[string]any{
		"message": message,
		"error":   err,
	}
}

func ErrorValidationResponse(message string, errors map[string]map[string]string) map[string]any {
	return map[string]any{
		"message": message,
		"error":   errors,
	}
}

func SuccessResponse(message string, data any) map[string]any {
	return map[string]any{
		"message": message,
		"data":    data,
	}
}