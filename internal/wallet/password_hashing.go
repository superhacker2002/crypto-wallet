package wallet

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

func hashPassword(password []byte) []byte {
	hasher := sha3.New256()
	hasher.Write(password)
	return []byte(hex.EncodeToString(hasher.Sum(nil)))
}
