package exported

import (
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"

	tmbytes "github.com/brc20-collab/brczero/libs/tendermint/libs/bytes"
)

// Evidence defines the contract which concrete evidence types of misbehavior
// must implement.
type Evidence interface {
	Route() string
	Type() string
	String() string
	Hash() tmbytes.HexBytes
	ValidateBasic() error

	// The consensus address of the malicious validator at time of infraction
	GetConsensusAddress() sdk.ConsAddress

	// Height at which the infraction occurred
	GetHeight() int64

	// The total power of the malicious validator at time of infraction
	GetValidatorPower() int64

	// The total validator set power at time of infraction
	GetTotalPower() int64
}
