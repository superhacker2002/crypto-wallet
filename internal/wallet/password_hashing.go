package wallet

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

func hashPassword(password string) string {
	hasher := sha3.New256()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
