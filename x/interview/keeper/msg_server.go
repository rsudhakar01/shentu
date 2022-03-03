package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/v2/x/interview/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) LockUser(goCtx context.Context, msg *types.MsgLockUser) (*types.MsgLockUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, found := k.Keeper.GetUser(ctx, msg.Id)
	if !found {
		return nil, types.ErrUserNotFound
	}
	if user.IsLocked {
		return nil, types.ErrUserAlreadyLocked
	}
	user.IsLocked = true
	k.Keeper.SetUser(ctx, user)

	return &types.MsgLockUserResponse{}, nil
}
