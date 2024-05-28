package helpers

func FailedResponse(response string) map[string]interface{} {
	return map[string]interface{}{
		"status":   "Failed",
		"messages": response,
	}
}

func SuccessResponse(response string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":   "success",
		"messages": response,
		"data":     data,
	}
}

func SuccessResponseHelper(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
	}
}