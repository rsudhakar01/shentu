package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgLockUser = "MsgLockUser"
)

// NewMsgLockUser creates a new MsgLockUser instance.
func NewMsgLockUser(from string, id uint64) *MsgLockUser {
	return &MsgLockUser{
		From: from,
		Id:   id,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgLockUser) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgLockUser) Type() string { return TypeMsgLockUser }

// GetSigners implements the sdk.Msg interface.
func (msg MsgLockUser) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgLockUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgLockUser) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return err
	}
	if from.Empty() {
		return ErrEmptySender
	}
	if msg.Id < 0 {
		return ErrInvalidId
	}
	return nil
}
