package models

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"messsage"`
}

func SetResponse(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
	}
}
