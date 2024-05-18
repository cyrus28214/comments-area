package server

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	codeSuccess        = 0
	codeInvalidRequest = 10000
	codeServerError    = 20000
)

func Success(data interface{}) Response {
	return Response{
		Code: codeSuccess,
		Msg:  "success",
		Data: data,
	}
}

func InvalidRequest(msg string) Response {
	return Response{
		Code: codeInvalidRequest,
		Msg:  msg,
		Data: nil,
	}
}

func InternalServerError(msg string) Response {
	return Response{
		Code: codeServerError,
		Msg:  msg,
		Data: nil,
	}
}
