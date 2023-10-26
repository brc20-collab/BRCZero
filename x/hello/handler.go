package hello

import (
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	sdkerrors "github.com/brc20-collab/brczero/libs/cosmos-sdk/types/errors"
	"github.com/brc20-collab/brczero/x/hello/internal/keeper"
	"github.com/brc20-collab/brczero/x/hello/internal/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx.SetEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgKV:
			return handleMsgKV(ctx, k, msg)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized bank message type: %T", msg)
		}
	}
}

func handleMsgKV(ctx sdk.Context, k keeper.Keeper, msg types.MsgKV) (*sdk.Result, error) {
	k.SetValue(ctx, msg.Key, msg.Value)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, "kv store"),
			sdk.NewAttribute(msg.Key, msg.Value),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
