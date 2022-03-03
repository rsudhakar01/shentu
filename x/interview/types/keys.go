package types

import "encoding/binary"

var (
	UserStoreKeyPrefix = []byte{0x0}
)

// GetUserKey gets the key for the user identified by user ID.
func GetUserKey(id uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, id)
	return append(UserStoreKeyPrefix, b...)
}
