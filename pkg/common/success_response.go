package common

type successResponse struct {
	Message   string        `json:"message"`
	Data      interface{} `json:"data"`
}

func NewSuccessResponse(message string, data interface{}) *successResponse {
	return &successResponse{
		Message:   message,
		Data:       data,
	}
}
