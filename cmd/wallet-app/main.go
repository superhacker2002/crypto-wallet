package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"wallet/internal/repository/password"
	"wallet/internal/wallet"
)

const (
	passwordsStorageFile = "data/wallet.txt"
	encryptionKey        = "M73aVlB63JArFuA4p4anyw=="
)

func main() {
	passwordRepo := password.New(passwordsStorageFile)
	walletService := wallet.NewService(passwordRepo)

	if fileInfo, err := os.Stat(passwordsStorageFile); os.IsNotExist(err) || fileInfo.Size() == 0 {
		fmt.Println("No accounts found. Creating a new one...")

		w, err := walletService.CreateWallet([]byte(encryptionKey))
		if err != nil {
			log.Fatalf("failed to create new wallet: %v", err)
		}

		fmt.Printf("New account has been created with address: %v\n", w.Address)
		return
	}

	// Пытаемся разблокировать кошелек
	fmt.Println("Enter password to unlock the wallet:")
	reader := bufio.NewReader(os.Stdin)
	psswd, _ := reader.ReadString('\n')
	if walletService.UnlockWallet(psswd) {
		fmt.Println("Wallet unlocked!")
	} else {
		fmt.Println("Incorrect password. Please, try again.")
	}
}
