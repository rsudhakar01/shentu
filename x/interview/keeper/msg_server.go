package keeper

import (
	"context"

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
	return &types.MsgLockUserResponse{}, nil
}
