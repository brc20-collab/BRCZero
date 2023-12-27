package types

type ApiResult[T interface{}] struct {
	Code int32  `json:"code" yaml:"code"`
	Msg  string `json:"msg" yaml:"msg"`
	Data T      `json:"data" yaml:"data"`
}

func NewOKApiResult[T interface{}](data T) ApiResult[T] {
	return ApiResult[T]{
		Code: 0,
		Msg:  "ok",
		Data: data,
	}
}
