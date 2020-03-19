package errors

// ErrorResponse 本文件中返回的错误结构体
type ErrorResponse struct {
	Error string `json:"error"`
}

// MakeErrorResponse 组装一个错误返回
func MakeErrorResponse(error error) ErrorResponse {

	return ErrorResponse{Error: FilterError(error.Error())}
}

// FilterError 错误过滤
func FilterError(error string) string {

	return error
}
