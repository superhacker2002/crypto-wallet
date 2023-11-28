package wallet

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"wallet/internal/encryption"
)

type Repository interface {
	SavePassword(password []byte) error
	Passwords() ([]byte, error)
}

type Service struct {
	r Repository
}

func NewService(r Repository) Service {
	return Service{r: r}
}

func (s Service) CreateWallet(encryptionKey []byte) (Wallet, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return Wallet{}, fmt.Errorf("error generating private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return Wallet{}, errors.New("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	privateKeyAES, err := encryption.Encrypt(crypto.FromECDSA(privateKey), encryptionKey)
	if err != nil {
		return Wallet{}, fmt.Errorf("failed to encrypt wallet private key: %w", err)
	}

	err = s.r.SavePassword(privateKeyAES)
	if err != nil {
		return Wallet{}, err
	}

	return New(address), nil
}

func (s Service) UnlockWallet(password string, decryptionKey []byte) bool {
	hashedPassword := hashPassword(password)

	encryptedData, err := s.r.Passwords()
	if err != nil {
		log.Fatal(err)
	}

	decryptedData, err := encryption.Decrypt(encryptedData, decryptionKey)
	if err != nil {
		log.Fatal(err)
	}

	hashedPasswordFromData := string(decryptedData)
	return hashedPassword == hashedPasswordFromData
}
