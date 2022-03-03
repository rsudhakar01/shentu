package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Error Code Enums
const (
	errEmptySender uint32 = iota + 101
	errInvalidId
)

var (
	ErrEmptySender = sdkerrors.Register(ModuleName, errEmptySender, "missing sender address")
	ErrInvalidId   = sdkerrors.Register(ModuleName, errInvalidId, "invalid user id")
)
