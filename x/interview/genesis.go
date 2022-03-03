package interview

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/v2/x/interview/keeper"
	"github.com/certikfoundation/shentu/v2/x/interview/types"
)

// InitGenesis initialize default parameters and the keeper's address to pubkey map.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	users := data.Users

	for _, user := range users {
		k.SetUser(ctx, user)
	}
}

// ExportGenesis writes the current store values to a genesis file, which can be imported again with InitGenesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	users := k.GetAllUsers(ctx)
	return types.NewGenesisState(users)
}
