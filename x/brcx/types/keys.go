package types

const (
	// ModuleName is the name of the module
	ModuleName = "brcx"

	// StoreKey is the store key string for slashing
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute is the querier route for slashing
	QuerierRoute = ModuleName

	MsgInscriptionType = "inscription"
	MsgBasicXType      = "basic_x"
)

var (
	ContractNameKey = []byte{0x01}
)

func GetContractAddressByName(name string) []byte {
	return append(ContractNameKey, []byte(name)...)
}
