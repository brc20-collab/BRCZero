package localhost

import "github.com/brc20-collab/brczero/libs/ibc-go/modules/light-clients/09-localhost/types"

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}
