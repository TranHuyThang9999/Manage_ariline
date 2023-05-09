package enums

type CodeResponse int

const (
	SuccessCode           CodeResponse = 0
	ErrorCodeNotFound     CodeResponse = 2004
	ErrorCodeSystemError  CodeResponse = 1
	ErrorCodeUnauthorized CodeResponse = 3
)
