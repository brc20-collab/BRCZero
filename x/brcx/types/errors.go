package types

import (
	"fmt"

	sdkerrors "github.com/brc20-collab/brczero/libs/cosmos-sdk/types/errors"
)

const (
	manageContract = 10
	callContract   = 100
)

var (
	ErrChainConfigNotFound = sdkerrors.Register(ModuleName, 2, "chain configuration not found")
)

func ErrInternal(msg string) *sdkerrors.Error {
	return sdkerrors.New(ModuleName, 1, msg)
}

func ErrUnknownOperationOfManageContract(operation string) *sdkerrors.Error {
	return sdkerrors.New(ModuleName, manageContract+1, fmt.Sprintf("%s is unknown operation of manage contract", operation))
}

func ErrValidateBasic(msg string) *sdkerrors.Error {
	return sdkerrors.New(ModuleName, manageContract+2, fmt.Sprintf("ManageContract validateBasic error : %s", msg))
}

func ErrValidateInput(msg string) *sdkerrors.Error {
	return sdkerrors.New(ModuleName, manageContract+3, msg)
}

func ErrExecute(msg string) *sdkerrors.Error {
	return sdkerrors.New(ModuleName, manageContract+4, msg)
}

func ErrGetContractAddress(msg string) *sdkerrors.Error {
	return sdkerrors.New(ModuleName, callContract+1, msg)
}

func ErrPackInput(msg string) *sdkerrors.Error {
	return sdkerrors.New(ModuleName, callContract+2, msg)
}

func ErrCallMethod(msg string) *sdkerrors.Error {
	return sdkerrors.New(ModuleName, callContract+3, msg)
}

func ErrUnpackOutput(msg string) *sdkerrors.Error {
	return sdkerrors.New(ModuleName, callContract+4, msg)
}
