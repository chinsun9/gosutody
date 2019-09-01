package pointers_and_errors

type Wallet struct {
	balance int
}

func (wallet *Wallet) Deposit(amount int) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() int {
	return wallet.balance
}