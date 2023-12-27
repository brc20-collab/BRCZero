package types

type ApiResult struct {
	Code int32                            `json:"code" yaml:"code"`
	Msg  string                           `json:"msg" yaml:"msg"`
	Data QueryTxEventsByBlockHashResponse `json:"data" yaml:"data"`
}

func NewOKApiResult(data QueryTxEventsByBlockHashResponse) ApiResult {
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
		Data: QueryTxEventsByBlockHashResponse{},
	}
}
