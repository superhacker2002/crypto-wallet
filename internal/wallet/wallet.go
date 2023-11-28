package wallet

type Wallet struct {
	Address string
}

func New(addr string) Wallet {
	return Wallet{Address: addr}
}
