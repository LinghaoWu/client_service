package controller

import "client_service/consts"

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	TraceID string      `json:"trace_id,omitempty"`
}

func RespSuccess() *Response {
	return &Response{
		Status:  consts.Success,
		Data:    nil,
		Message: consts.GetMsg(consts.Success),
	}
}

func RespSuccessWithData(data interface{}, code ...int) *Response {
	return &Response{
		Status:  consts.Success,
		Data:    data,
		Message: consts.GetMsg(consts.Success),
	}
}

func RespError(code int, errMsg error, data string) *Response {
	return &Response{
		Status:  code,
		Data:    data,
		Message: consts.GetMsg(code),
		Error:   errMsg.Error(),
	}
}
