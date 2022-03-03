package types

// NewUser creates a new user object
func NewUser(id uint64, name string, isLocked bool) User {
	return User{
		Id:       id,
		Name:     name,
		IsLocked: isLocked,
	}
}
