package types

const (
	QueryKV = "kv"
)

type QueryValueParams struct {
	Key string
}

func NewQueryValueParams(key string) QueryValueParams {
	return QueryValueParams{key}
}
