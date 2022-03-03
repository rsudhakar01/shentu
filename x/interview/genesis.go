package interview

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/v2/x/interview/keeper"
	"github.com/certikfoundation/shentu/v2/x/interview/types"
)

func InitDefaultGenesis(ctx sdk.Context, k keeper.Keeper) {
	InitGenesis(ctx, k, *types.DefaultGenesisState())
}

// InitGenesis initialize default parameters and the keeper's address to pubkey map.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {

}

// ExportGenesis writes the current store values to a genesis file, which can be imported again with InitGenesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return nil
}
