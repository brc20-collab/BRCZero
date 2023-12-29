package brcx

import (
	"github.com/brc20-collab/brczero/x/brcx/internal/keeper"
	"github.com/brc20-collab/brczero/x/brcx/types"
)

const (
	ModuleName                 = types.ModuleName
	StoreKey                   = types.StoreKey
	RouterKey                  = types.RouterKey
	QuerierRoute               = types.QuerierRoute
	ManageCreateContract       = types.ManageCreateContract
	ManageCallContract         = types.ManageCallContract
	ManageContractProtocolName = types.ManageContractProtocolName

	AttributeProtocolName            = types.AttributeProtocolName
	EventTypeBasicX                  = types.EventTypeBasicX
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
	NewMsgInscription                   = types.NewMsgInscription
	NewMsgBasicProtocolOp               = types.NewMsgBasicProtocolOp

	ErrInternal           = types.ErrInternal
	ErrValidateInput      = types.ErrValidateInput
	ErrExecute            = types.ErrExecute
	ErrGetContractAddress = types.ErrGetContractAddress
	ErrCallMethod         = types.ErrCallMethod
	ErrPackInput          = types.ErrPackInput
)

type (
	Keeper             = keeper.Keeper
	MsgInscription     = types.MsgInscription
	MsgBasicProtocolOp = types.MsgBasicProtocolOp
	ManageContract     = types.ManageContract
	ResultInfo         = types.ResultInfo
	InscriptionContext = types.InscriptionContext
)
