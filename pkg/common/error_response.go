package common

type errorFormat struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func NewError(message string, detail string) *errorFormat {
	return &errorFormat{
		Message: message,
		Detail:  detail,
	}
}

func NewSimpleError(message string) *errorFormat {
	return &errorFormat{
		Message: message,
		Detail:  "",
	}
}