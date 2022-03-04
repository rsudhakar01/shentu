package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Error Code Enums
const (
	errEmptySender uint32 = iota + 101
	errInvalidId
	errUserNotFound
	errUserAlreadyLocked
	errUserAlreadyUnlocked
)

var (
	ErrEmptySender         = sdkerrors.Register(ModuleName, errEmptySender, "missing sender address")
	ErrInvalidId           = sdkerrors.Register(ModuleName, errInvalidId, "invalid user id")
	ErrUserNotFound        = sdkerrors.Register(ModuleName, errUserNotFound, "user not found")
	ErrUserAlreadyLocked   = sdkerrors.Register(ModuleName, errUserAlreadyLocked, "user already locked")
	ErrUserAlreadyUnlocked = sdkerrors.Register(ModuleName, errUserAlreadyUnlocked, "user already unlocked")
)
