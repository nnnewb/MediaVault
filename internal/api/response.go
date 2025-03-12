package api

import "fmt"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ListData struct {
	Total int64 `json:"total"`
	Data  any   `json:"data"`
}

func (r *Response) String() string {
	return fmt.Sprintf("Response(Code=%d, Message=%s, Data=%T)", r.Code, r.Message, r.Data)
}

func NewResponse(code int, message string, data any) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func OK(data any) *Response {
	return NewResponse(0, "OK", data)
}

func OKList(data any, total int64) *Response {
	return NewResponse(0, "OK", ListData{Total: total, Data: data})
}

func BadRequest(err error) *Response {
	return NewResponse(2, "Bad Request: "+err.Error(), nil)
}

func ServerError(err error) *Response {
	return NewResponse(1, "Server Error: "+err.Error(), nil)
}
