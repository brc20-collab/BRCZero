package types

type ApiResponse[T interface{}] struct {
	Code int32  `json:"code" yaml:"code"`
	Msg  string `json:"msg" yaml:"msg"`
	Data T      `json:"data" yaml:"data"`
}

func NewOKApiResponse[T interface{}](data T) ApiResponse[T] {
	return ApiResponse[T]{
		Code: 0,
		Msg:  "ok",
		Data: data,
	}
}

type ApiError struct {
	Code int    `json:"code" yaml:"code"`
	Msg  string `json:"msg" yaml:"msg"`
}

func NewApiError(code int, msg string) ApiError {
	return ApiError{
		Code: code,
		Msg:  msg,
	}
}
