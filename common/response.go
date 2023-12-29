package common

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(success bool, message string, data interface{}) *Response {
	return &Response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func NewSuccessResponse(message string, data interface{}) *Response {
	return NewResponse(true, message, data)
}

func NewErrorResponse(message string, data interface{}) *Response {
	return NewResponse(false, message, data)
}
