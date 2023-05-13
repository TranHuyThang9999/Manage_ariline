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

type ResponseResourcTooken struct {
	Code    enums.CodeResponse `json:"code"`
	Message string             `json:"message"`
	Token   interface{}        `json:"token"`
}

func NewResponseResourceToken(code enums.CodeResponse, message string, token interface{}) *ResponseResourcTooken {
	return &ResponseResourcTooken{
		Code:    code,
		Message: message,
		Token:   token,
	}
}
