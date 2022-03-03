package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper manages key-value store.
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryCodec
}

// NewKeeper creates a new keeper object.
func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}
