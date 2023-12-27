package types

type ApiResult struct {
	Code int32       `json:"code" yaml:"code"`
	Msg  string      `json:"msg" yaml:"msg"`
	Data interface{} `json:"data" yaml:"data"`
}
