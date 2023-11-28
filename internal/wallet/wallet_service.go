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
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyAES, err := encryption.Encrypt(privateKeyBytes, encryptionKey)
	if err != nil {
		return Wallet{}, fmt.Errorf("failed to encrypt wallet private key: %w", err)
	}

	hashedPrivateKey := hashPassword(privateKeyAES)
	err = s.r.SavePassword(hashedPrivateKey)
	if err != nil {
		return Wallet{}, err
	}

	return New(address), nil
}

func (s Service) UnlockWallet(userPasswd string) bool {
	passwords, err := s.r.Passwords()
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword := string(passwords)
	return hashedPassword == userPasswd
}
