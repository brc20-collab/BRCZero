package types

import (
	_ "embed"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	Src20GetTickInfoMethod    = "getSRC20TokenInfo"
	Src20GetBalanceMethod     = "getSRC20Balance"
	Src20GetAllBalanceMethod  = "getAllSRC20Balance"
	Src20GetAllTickInfoMethod = "getSRC20AllTokenInfo"
)

var (
	src20EntryPointABI abi.ABI
	//go:embed abi/Src20Entrypoint.json
	src20EntryPointABIJson []byte

	src20ABI abi.ABI
	//go:embed abi/Src20.json
	src20ABIJson []byte
)

func init() {
	src20EntryPointABI = GetEVMABIConfig(src20EntryPointABIJson)
	src20ABI = GetEVMABIConfig(src20ABIJson)
}

func UnpackSrc20EventContext(ret []byte) (Src20WrappedEvent, error) {
	var ec Src20WrappedEvent
	err := src20ABI.UnpackIntoInterface(&ec, EventsName, ret)
	if err != nil {
		return Src20WrappedEvent{}, err
	}
	return ec, nil
}

func GetSrc20TickInfoInput(tickName string) ([]byte, error) {
	data, err := src20EntryPointABI.Pack(Src20GetTickInfoMethod, tickName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnpackGetSrc20TickInfoOutput(ret []byte) (WrappedSrc20TokenInfo, error) {
	var output WrappedSrc20TokenInfo
	err := src20EntryPointABI.UnpackIntoInterface(&output, Src20GetTickInfoMethod, ret)
	if err != nil {
		return WrappedSrc20TokenInfo{}, err
	}

	return output, err
}

func GetSrc20BalanceInput(address string, tickName string) ([]byte, error) {
	data, err := src20EntryPointABI.Pack(Src20GetBalanceMethod, address, tickName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnpackGetSrc20BalanceOutput(ret []byte) (WrappedSrc20Balance, error) {
	var output WrappedSrc20Balance
	err := src20EntryPointABI.UnpackIntoInterface(&output, Src20GetBalanceMethod, ret)
	if err != nil {
		return WrappedSrc20Balance{}, err
	}

	return output, err
}

func GetSrc20AllBalanceInput(address string) ([]byte, error) {
	data, err := src20EntryPointABI.Pack(Src20GetAllBalanceMethod, address)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnpackGetSrc20AllBalanceOutput(ret []byte) ([]Src20Balance, error) {
	var output []Src20Balance
	err := src20EntryPointABI.UnpackIntoInterface(&output, Src20GetAllBalanceMethod, ret)
	if err != nil {
		return nil, err
	}

	return output, err
}

func GetSrc20AllTickInfoInput(page, pageSize *big.Int) ([]byte, error) {
	data, err := src20EntryPointABI.Pack(Src20GetAllTickInfoMethod, page, pageSize)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnpackGetSrc20AllTickInfoOutput(ret []byte) (WrappedSrc20AllTokenInfo, error) {
	var output WrappedSrc20AllTokenInfo
	err := src20EntryPointABI.UnpackIntoInterface(&output, Src20GetAllTickInfoMethod, ret)
	if err != nil {
		return WrappedSrc20AllTokenInfo{}, err
	}

	return output, err
}
