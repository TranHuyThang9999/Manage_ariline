package resources

import "btl/core/enums"

type ResponseResource struct {
	Code    enums.CodeResponse `json:"code"`
	Message string             `json:"message"`
	Body    interface{}        `json:"body"`
}

func NewResponseResource(code enums.CodeResponse, message string, body interface{}) *ResponseResource {
	return &ResponseResource{
		Code:    code,
		Message: message,
		Body:    body,
	}
}
