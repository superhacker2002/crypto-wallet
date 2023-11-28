package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	psswdrepo "wallet/internal/repository/password"
	"wallet/internal/wallet"
)

var key = []byte("supersecretkey123")

const passwordsStorageFile = "data/wallet.dat"

func main() {
	passwordRepo := psswdrepo.New(passwordsStorageFile)
	walletService := wallet.NewService(passwordRepo)

	if _, err := os.Stat(passwordsStorageFile); os.IsNotExist(err) {
		fmt.Println("No accounts found. Creating a new one...")
		w, err := walletService.CreateWallet()
		if err != nil {
			log.Fatalf("failed to create new wallet: %v", err)
		}
		fmt.Printf("New account has been successfully created with address: %v", w.Address)
	} else {
		fmt.Println("Enter password to unlock the wallet:")
		reader := bufio.NewReader(os.Stdin)
		password, _ := reader.ReadString('\n')
		if walletService.UnlockWallet(password, key) {
			fmt.Println("Wallet unlocked!")
		} else {
			fmt.Println("Incorrect password. Please, try again.")
		}
	}
}
