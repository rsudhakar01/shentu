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
	return &GenesisState{
		Users: []User{
			{
				Id:       1,
				Name:     "Alice",
				IsLocked: false,
			},
			{
				Id:       2,
				Name:     "Bob",
				IsLocked: false,
			},
			{
				Id:       3,
				Name:     "Mary",
				IsLocked: false,
			},
		},
		NextUserId: 4,
	}
}

// ValidateGenesis - validate crisis genesis data
func ValidateGenesis(_ json.RawMessage) error {
	return nil
}
