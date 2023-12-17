package types

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ManageContractProtocolName = "brczero"
	ManageCreateContract       = "create"
	ManageCallContract         = "call"
)

type ManageContract struct {
	Protocol  string `json:"p" yaml:"p"`
	Operation string `json:"op" yaml:"op"`
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Contract  string `json:"contract,omitempty" yaml:"contract,omitempty"`
	WaitInit  uint8  `json:"wait_init" yaml:"waitInit"`
	CallData  string `json:"d" yaml:"d"`
}

func (mc ManageContract) ValidateBasic() error {
	switch mc.Operation {
	case ManageCreateContract:
		if len(mc.Name) == 0 {
			return ErrValidateBasic("name can not be empty when create contract")
		}
	case ManageCallContract:
		if len(mc.Contract) == 0 {
			return ErrValidateBasic("contract can not be empty when call contract")
		}
		if !common.IsHexAddress(mc.Contract) {
			return ErrValidateBasic("contract must be hex address when call contract")
		}
	default:
		return ErrUnknownOperationOfManageContract(mc.Operation)
	}
	if _, err := hex.DecodeString(mc.CallData); err != nil {
		return ErrValidateBasic(fmt.Sprintf("decode callData error: %s", err.Error()))
	}

	return nil
}

func (mc ManageContract) GetCallData() ([]byte, error) {
	return hex.DecodeString(mc.CallData)
}
