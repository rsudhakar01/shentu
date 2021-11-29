package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// RegisterInvariants registers all governance invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper, bk govtypes.BankKeeper) {
	ir.RegisterRoute(govtypes.ModuleName, "module-account", ModuleAccountInvariant(keeper, bk))
}

// AllInvariants runs all invariants of the governance module
func AllInvariants(keeper Keeper, bk govtypes.BankKeeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return ModuleAccountInvariant(keeper, bk)(ctx)
	}
}

// ModuleAccountInvariant checks that the module account coins reflects the sum of
// deposit amounts held on store
func ModuleAccountInvariant(keeper Keeper, bk govtypes.BankKeeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var expectedDeposits sdk.Coins

		keeper.IterateAllDeposits(ctx, func(deposit govtypes.Deposit) bool {
			expectedDeposits = expectedDeposits.Add(deposit.Amount...)
			return false
		})

		macc := keeper.GetGovernanceAccount(ctx)
		balances := bk.GetAllBalances(ctx, macc.GetAddress())
		broken := !balances.IsEqual(expectedDeposits)

		return sdk.FormatInvariant(govtypes.ModuleName, "deposits",
			fmt.Sprintf("\tgov ModuleAccount coins: %s\n\tsum of deposit amounts:  %s\n",
				balances, expectedDeposits)), broken
	}
}
