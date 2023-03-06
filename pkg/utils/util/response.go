package util

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var WrapErrorResponse = func(message string, data interface{}) Response {
	return Response{
		Status:  false,
		Message: message,
		Data:    data,
	}
}

var WrapSuccessResponse = func(message string, data interface{}) Response {
	return Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}
