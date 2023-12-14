package types

const (
	ProtocolBRC20 = "brc-20"

	QueryTick                = "tick"
	QueryAllTick             = "allTick"
	QueryBalance             = "balance"
	QueryAllBalance          = "allBalance"
	QueryTotalTickHolders    = "totalTickHolders"
	QueryTransferableTick    = "transferableTick"
	QueryAllTransferableTick = "allTransferableTick"
)

type QueryTickParams struct {
	Name string
}

func NewQueryTickParams(name string) QueryTickParams {
	return QueryTickParams{
		Name: name,
	}
}

type QueryDataParams struct {
	Addr string
	Name string
}

func NewQueryDataParams(addr string, name string) QueryDataParams {
	return QueryDataParams{
		Addr: addr,
		Name: name,
	}
}

type QueryAllDataParams struct {
	Addr string
}

func NewQueryAllDataParams(addr string) QueryAllDataParams {
	return QueryAllDataParams{
		Addr: addr,
	}
}
