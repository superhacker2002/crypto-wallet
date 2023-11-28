# Crypto Wallet

A small CLI application that allows to create and get access to the wallet.

## Description

Through the go Ethereum library, you can create wallets, a password for them is generated there.
In the future, this password can be used to access the wallet. The application, when launched, if there are no accounts, 
asks you to create it, and if there is, asks for a password. That is, the account in this case will be a created wallet.
The password that is transmitted is kept in a file in the form of a hash. The data in the file 
with the hash is encrypted using AES, the encryption key is used with some kind of hardcoded thread. 
If the password matches the hash from the encrypted file, then it is used to unlock the wallet.