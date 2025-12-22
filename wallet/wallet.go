package wallet

import "fmt"

type Wallet struct {
	Balance float64
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
}

func (w *Wallet) SpendMoney(amount float64) {
	if amount > w.Balance {
		fmt.Println("Not enough balance")
		return
	}
	w.Balance -= amount
}

func (w *Wallet) GetBalance() float64 {
	return w.Balance
}
