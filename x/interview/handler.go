package interview

import (
	"github.com/certikfoundation/shentu/v2/x/cert/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "cert" type messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	// msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Unrecognized cert Msg type: %v", msg)
		}
	}
}
