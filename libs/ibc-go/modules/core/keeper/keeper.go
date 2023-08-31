package keeper

import (
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	types2 "github.com/brc20-collab/brczero/libs/cosmos-sdk/codec/types"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	capabilitykeeper "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/capability/keeper"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/params"
	paramtypes "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/params"
	clientkeeper "github.com/brc20-collab/brczero/libs/ibc-go/modules/core/02-client/keeper"
	clienttypes "github.com/brc20-collab/brczero/libs/ibc-go/modules/core/02-client/types"
	connectionkeeper "github.com/brc20-collab/brczero/libs/ibc-go/modules/core/03-connection/keeper"
	connectiontypes "github.com/brc20-collab/brczero/libs/ibc-go/modules/core/03-connection/types"
	channelkeeper "github.com/brc20-collab/brczero/libs/ibc-go/modules/core/04-channel/keeper"
	portkeeper "github.com/brc20-collab/brczero/libs/ibc-go/modules/core/05-port/keeper"
	porttypes "github.com/brc20-collab/brczero/libs/ibc-go/modules/core/05-port/types"
	"github.com/brc20-collab/brczero/libs/ibc-go/modules/core/types"
)

var _ types.QueryServer = (*Keeper)(nil)
var _ IBCServerKeeper = (*Keeper)(nil)

// Keeper defines each ICS keeper for IBC
type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryService

	cdc        *codec.CodecProxy
	paramSpace params.Subspace

	ClientKeeper     clientkeeper.Keeper
	ConnectionKeeper connectionkeeper.Keeper
	ChannelKeeper    channelkeeper.Keeper
	PortKeeper       portkeeper.Keeper
	Router           *porttypes.Router
}

// NewKeeper creates a new ibc Keeper
func NewKeeper(
	proxy *codec.CodecProxy,
	key sdk.StoreKey, paramSpace paramtypes.Subspace,
	stakingKeeper clienttypes.StakingKeeper, upgradeKeeper clienttypes.UpgradeKeeper,
	scopedKeeper *capabilitykeeper.ScopedKeeper,
	registry types2.InterfaceRegistry,
) *Keeper {
	//mm := codec.NewProtoCodec(registry)
	//proxy:=codec.NewMarshalProxy(mm,cdcc)
	if !paramSpace.HasKeyTable() {
		keyTable := types.ParamKeyTable()
		keyTable.RegisterParamSet(&clienttypes.Params{})
		keyTable.RegisterParamSet(&connectiontypes.Params{})
		paramSpace = paramSpace.WithKeyTable(keyTable)
	}
	clientKeeper := clientkeeper.NewKeeper(proxy, key, paramSpace, stakingKeeper, upgradeKeeper)
	connectionKeeper := connectionkeeper.NewKeeper(proxy, key, paramSpace, clientKeeper)
	portKeeper := portkeeper.NewKeeper(scopedKeeper)
	channelKeeper := channelkeeper.NewKeeper(proxy, key, clientKeeper, connectionKeeper, portKeeper, scopedKeeper)

	return &Keeper{
		cdc:              proxy,
		ClientKeeper:     clientKeeper,
		ConnectionKeeper: connectionKeeper,
		ChannelKeeper:    channelKeeper,
		PortKeeper:       portKeeper,
		paramSpace:       paramSpace,
	}
}

// Codec returns the IBC module codec.
func (k Keeper) Codec() *codec.CodecProxy {
	return k.cdc
}

// SetRouter sets the Router in IBC Keeper and seals it. The method panics if
// there is an existing router that's already sealed.
func (k *Keeper) SetRouter(rtr *porttypes.Router) {
	if k.Router != nil && k.Router.Sealed() {
		panic("cannot reset a sealed router")
	}

	k.PortKeeper.Router = rtr
	k.Router = rtr
	k.Router.Seal()
}

///
func (k Keeper) GetPacketReceipt(ctx sdk.Context, portID, channelID string, sequence uint64) (string, bool) {
	return k.ChannelKeeper.GetPacketReceipt(ctx, portID, channelID, sequence)
}

func (k Keeper) GetPacketCommitment(ctx sdk.Context, portID, channelID string, sequence uint64) []byte {
	return k.ChannelKeeper.GetPacketCommitment(ctx, portID, channelID, sequence)
}
