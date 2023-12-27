package types

type ApiResult struct {
	Code int32                                 `json:"code" yaml:"code"`
	Msg  string                                `json:"msg" yaml:"msg"`
	Data QueryBrc20TxEventsByBlockHashResponse `json:"data" yaml:"data"`
}

func NewOKApiResult(data QueryBrc20TxEventsByBlockHashResponse) ApiResult {
	return ApiResult{
		Code: 0,
		Msg:  "ok",
		Data: data,
	}
}

func NewErrApiResult(err error) ApiResult {
	return ApiResult{
		Code: 1,
		Msg:  err.Error(),
		Data: QueryBrc20TxEventsByBlockHashResponse{},
	}
}
