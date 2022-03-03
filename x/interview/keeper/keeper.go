package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/v2/x/interview/types"
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

// SetUser sets data of a user in kv-store.
func (k Keeper) SetUser(ctx sdk.Context, user types.User) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalLengthPrefixed(&user)
	store.Set(types.GetUserKey(user.Id), bz)
}

// GetUser gets data of a user given user ID.
func (k Keeper) GetUser(ctx sdk.Context, id uint64) (types.User, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetUserKey(id))
	if bz == nil {
		return types.User{}, false
	}
	var user types.User
	k.cdc.MustUnmarshalLengthPrefixed(bz, &user)
	return user, true
}

// IterateAllUsers iterates over the all the stored users and performs a callback function.
func (k Keeper) IterateAllUsers(ctx sdk.Context, callback func(user types.User) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UserStoreKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var user types.User
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &user)

		if callback(user) {
			break
		}
	}
}

// GetAllUsers retrieves all users in the store.
func (k Keeper) GetAllUsers(ctx sdk.Context) (users []types.User) {
	k.IterateAllUsers(ctx, func(user types.User) bool {
		users = append(users, user)
		return false
	})
	return users
}

// SetNextUserID sets the latest pool ID to store.
func (k Keeper) SetNextUserID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, id)
	store.Set(types.GetNextUserIDKey(), bz)
}

// GetNextUserID gets the latest pool ID from store.
func (k Keeper) GetNextUserID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	opBz := store.Get(types.GetNextUserIDKey())
	return binary.LittleEndian.Uint64(opBz)
}
