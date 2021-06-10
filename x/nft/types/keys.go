package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName is the name of this module
	ModuleName = "nft"

	// RouterKey is used to route messages.
	RouterKey = ModuleName

	// StoreKey is the prefix under which we store this module's data.
	StoreKey = ModuleName

	// QuerierRoute is used to handle abci_query requests.
	QuerierRoute = ModuleName
)

var (
	AdminKeyPrefix = []byte{0x10}
)

func AdminKey(addr sdk.AccAddress) []byte {
	return append(AdminKeyPrefix, addr...)
}
