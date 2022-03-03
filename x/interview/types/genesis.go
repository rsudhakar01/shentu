package types

import (
	"encoding/json"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(users []User, nextUserId uint64) GenesisState {
	return GenesisState{
		Users:      users,
		NextUserId: nextUserId,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState {
	return &GenesisState{NextUserId: 1}
}

// ValidateGenesis - validate crisis genesis data
func ValidateGenesis(_ json.RawMessage) error {
	return nil
}
