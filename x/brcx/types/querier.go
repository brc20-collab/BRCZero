package types

const (
	//todo: fix p name
	ProtocolBRC20 = "leotest2"

	QueryTick             = "tick"
	QueryAllTick          = "allTick"
	QueryBalance          = "balance"
	QueryAllBalance       = "allBalance"
	QueryTotalTickHolders = "totalTickHolders"
)

type QueryTickParams struct {
	Name string
}

func NewQueryTickParams(name string) QueryTickParams {
	return QueryTickParams{
		Name: name,
	}
}

type QueryBalanceParams struct {
	Addr string
	Name string
}

func NewQueryBalanceParams(addr string, name string) QueryBalanceParams {
	return QueryBalanceParams{
		Addr: addr,
		Name: name,
	}
}

type QueryAllBalanceParams struct {
	Addr string
}

func NewQueryAllBalanceParams(addr string) QueryAllBalanceParams {
	return QueryAllBalanceParams{
		Addr: addr,
	}
}
