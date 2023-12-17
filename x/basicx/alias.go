package basicx

import (
	"github.com/brc20-collab/brczero/x/basicx/internal/keeper"
	"github.com/brc20-collab/brczero/x/basicx/types"
)

const (
	ModuleName                 = types.ModuleName
	StoreKey                   = types.StoreKey
	RouterKey                  = types.RouterKey
	QuerierRoute               = types.QuerierRoute
	ManageCreateContract       = types.ManageCreateContract
	ManageCallContract         = types.ManageCallContract
	ManageContractProtocolName = types.ManageContractProtocolName

	AttributeProtocol                = types.AttributeProtocol
	EventTypeBasicx                  = types.EventTypeBasicx
	EventTypeBasicxProtocol          = types.EventTypeBasicxProtocol
	EventTypeManageContract          = types.EventTypeManageContract
	EventTypeEntryPoint              = types.EventTypeEntryPoint
	AttributeManageContractOperation = types.AttributeManageContractOperation

	AttributeManageContractAddress = types.AttributeManageContractAddress
	AttributeEvmOutput             = types.AttributeEvmOutput
	AttributeManageLog             = types.AttributeManageLog
	AttributeBTCTXID               = types.AttributeBTCTXID
)

var (
	RegisterCodec = types.RegisterCodec
	ModuleCdc     = types.ModuleCdc

	NewKeeper                           = keeper.NewKeeper
	NewQuerier                          = keeper.NewQuerier
	ErrUnknownOperationOfManageContract = types.ErrUnknownOperationOfManageContract
	ConvertBTCAddress                   = types.ConvertBTCAddress
	NewMsgProtocolOp                    = types.NewMsgProtocolOp

	ErrInternal           = types.ErrInternal
	ErrValidateInput      = types.ErrValidateInput
	ErrExecute            = types.ErrExecute
	ErrGetContractAddress = types.ErrGetContractAddress
	ErrCallEntryPoint     = types.ErrCallEntryPoint
	ErrPackInput          = types.ErrPackInput
)

type (
	Keeper             = keeper.Keeper
	MsgProtocolOp      = types.MsgProtocolOp
	ManageContract     = types.ManageContract
	ResultInfo         = types.ResultInfo
	InscriptionContext = types.InscriptionContext
)
