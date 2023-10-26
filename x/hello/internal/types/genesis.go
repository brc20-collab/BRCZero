package types

type GenesisState struct {
	EnableHello bool `json:"enable_hello" yaml:"enable_hello"`
}

func NewGenesisState() GenesisState {
	return GenesisState{EnableHello: true}
}

func DefaultGenesisState() GenesisState { return NewGenesisState() }

func ValidateGenesis(data GenesisState) error { return nil }
