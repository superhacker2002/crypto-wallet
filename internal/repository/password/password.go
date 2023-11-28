package password

import (
	"fmt"
	"os"
)

type Repository struct {
	fileName string
}

func New(name string) Repository {
	return Repository{fileName: name}
}

func (r Repository) SavePassword(password []byte) error {
	err := os.WriteFile(r.fileName, password, 0644)
	if err != nil {
		return fmt.Errorf("error saving password to a file: %w", err)
	}

	return nil
}

func (r Repository) Passwords() ([]byte, error) {
	encryptedData, err := os.ReadFile(r.fileName)
	if err != nil {
		return nil, fmt.Errorf("error getting passwords from a file: %w", err)
	}

	return encryptedData, nil
}
