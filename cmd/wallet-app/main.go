package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	psswdrepo "wallet/internal/repository/password"
	"wallet/internal/wallet"
)

var key = []byte("M73aVlB63JArFuA4p4anyw==")

const passwordsStorageFile = "data/wallet.txt"

func main() {
	passwordRepo := psswdrepo.New(passwordsStorageFile)
	walletService := wallet.NewService(passwordRepo)

	if fileInfo, err := os.Stat(passwordsStorageFile); os.IsNotExist(err) || fileInfo.Size() == 0 {
		fmt.Println("No accounts found. Creating a new one...")
		w, err := walletService.CreateWallet(key)
		if err != nil {
			log.Fatalf("failed to create new wallet: %v", err)
		}
		fmt.Printf("New account has been created with address: %v\n", w.Address)
		return
	}

	fmt.Println("Enter password to unlock the wallet:")
	reader := bufio.NewReader(os.Stdin)
	password, _ := reader.ReadString('\n')
	if walletService.UnlockWallet(password) {
		fmt.Println("Wallet unlocked!")
	} else {
		fmt.Println("Incorrect password. Please, try again.")
	}
}
