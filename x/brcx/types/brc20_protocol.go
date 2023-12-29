package types

import (
	_ "embed"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	Brc20GetTickInfoMethod            = "getTickInformation"
	Brc20GetAllTickInfoMethod         = "getAllTickInformation"
	Brc20GetBalanceMethod             = "getBalance"
	Brc20GetAllBalanceMethod          = "getAllBalance"
	Brc20GetTotalTickHoldersMethod    = "getTotalTickHolders"
	Brc20GetTransferableTickMethod    = "getUserTransferableTickInformation"
	Brc20GetAllTransferableTickMethod = "getUserAllTransferableTickInformation"
)

var (
	brc20EntryPointABI abi.ABI
	//go:embed abi/BRC20EntryPoint.json
	brc20EntryPointABIJson []byte

	brc20ABI abi.ABI
	//go:embed abi/BRC20.json
	brc20ABIJson []byte
)

func init() {
	brc20EntryPointABI = GetEVMABIConfig(brc20EntryPointABIJson)
	brc20ABI = GetEVMABIConfig(brc20ABIJson)
}

func GetBrc20EntryPointInput(context InscriptionContext, inscription string) ([]byte, error) {
	data, err := brc20EntryPointABI.Pack(EntryPointMethodName, context, inscription)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetBrc20TickInfoInput(tickName string) ([]byte, error) {
	data, err := brc20EntryPointABI.Pack(Brc20GetTickInfoMethod, tickName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnpackGetBrc20TickInfoOutput(ret []byte) (WrappedBRC20Information, error) {
	var output WrappedBRC20Information
	err := brc20EntryPointABI.UnpackIntoInterface(&output, Brc20GetTickInfoMethod, ret)
	if err != nil {
		return WrappedBRC20Information{}, err
	}

	return output, err
}

func GetBrc20AllTickInfoInput() ([]byte, error) {
	data, err := brc20EntryPointABI.Pack(Brc20GetAllTickInfoMethod)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnpackGetBrc20AllTickInfoOutput(ret []byte) ([]BRC20Information, error) {
	var output []BRC20Information
	err := brc20EntryPointABI.UnpackIntoInterface(&output, Brc20GetAllTickInfoMethod, ret)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func GetBrc20BalanceInput(addr string, tickName string) ([]byte, error) {
	data, err := brc20EntryPointABI.Pack(Brc20GetBalanceMethod, addr, tickName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnpackGetBrc20BalanceOutput(ret []byte) (*big.Int, *big.Int, *big.Int, error) {
	res, err := brc20EntryPointABI.Methods[Brc20GetBalanceMethod].Outputs.Unpack(ret)
	if len(res) != 3 || err != nil {
		return nil, nil, nil, err
	}

	totalBalances, ok := res[0].(*big.Int)
	if !ok {
		return nil, nil, nil, errors.New("decode totalBalances failed")
	}

	availableBalance, ok := res[1].(*big.Int)
	if !ok {
		return nil, nil, nil, errors.New("decode availableBalance failed")
	}

	transferableBalance, ok := res[2].(*big.Int)
	if !ok {
		return nil, nil, nil, errors.New("decode transferableBalance failed")
	}

	return totalBalances, availableBalance, transferableBalance, nil
}

func GetBrc20AllBalanceInput(addr string) ([]byte, error) {
	data, err := brc20EntryPointABI.Pack(Brc20GetAllBalanceMethod, addr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnpackGetBrc20AllBalanceOutput(ret []byte) ([]BRC20Balance, error) {
	var output []BRC20Balance
	err := brc20EntryPointABI.UnpackIntoInterface(&output, Brc20GetAllBalanceMethod, ret)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func GetBrc20TotalTickHoldersInput() ([]byte, error) {
	data, err := brc20EntryPointABI.Pack(Brc20GetTotalTickHoldersMethod)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnpackGetBrc20TotalTickHoldersOutput(ret []byte) (*big.Int, error) {
	res, err := brc20EntryPointABI.Methods[Brc20GetTotalTickHoldersMethod].Outputs.Unpack(ret)
	if len(res) != 1 || err != nil {
		return nil, err
	}

	holders, ok := res[0].(*big.Int)
	if !ok {
		return nil, errors.New("decode totalTickHolders failed")
	}

	return holders, nil
}

func UnpackBrc20EventContext(ret []byte) (Brc20WrappedEvent, error) {
	var ec Brc20WrappedEvent
	err := brc20ABI.UnpackIntoInterface(&ec, EventsName, ret)
	if err != nil {
		return Brc20WrappedEvent{}, err
	}
	return ec, nil
}

func GetBrc20TransferableTickInput(addr string, tickName string) ([]byte, error) {
	data, err := brc20EntryPointABI.Pack(Brc20GetTransferableTickMethod, addr, tickName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnpackGetBrc20TransferableTickOutput(ret []byte) ([]Brc20TransferableInscription, error) {
	var tc []Brc20TransferableInscription
	err := brc20EntryPointABI.UnpackIntoInterface(&tc, Brc20GetTransferableTickMethod, ret)
	if err != nil {
		return nil, err
	}
	return tc, nil
}

func GetBrc20AllTransferableTickInput(addr string) ([]byte, error) {
	data, err := brc20EntryPointABI.Pack(Brc20GetAllTransferableTickMethod, addr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnpackGetBrc20AllTransferableTickOutput(ret []byte) ([]Brc20TransferableInscription, error) {
	var tc []Brc20TransferableInscription
	err := brc20EntryPointABI.UnpackIntoInterface(&tc, Brc20GetAllTransferableTickMethod, ret)
	if err != nil {
		return nil, err
	}
	return tc, nil
}
