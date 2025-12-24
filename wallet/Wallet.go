package Wallet

import "fmt"

type Wallet struct {
	Balance float64
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
}

func (w *Wallet) SpendMoney(amount float64) {
	if amount > w.Balance {
		fmt.Println("Insufficient funds")
		return
	}
	w.Balance -= amount
}

func (w *Wallet) GetBalance() float64 {
	return w.Balance
}

func RunWalletMenu() {
	wallet := Wallet{}
	transactions := []float64{}

	for {
		fmt.Println("\n--- WALLET MENU ---")
		fmt.Println("1. Add Money")
		fmt.Println("2. Spend Money")
		fmt.Println("3. Balance")
		fmt.Println("0. Back")

		var choice int
		fmt.Scanln(&choice)

		var amount float64
		switch choice {
		case 1:
			fmt.Print("Amount: ")
			fmt.Scanln(&amount)
			wallet.AddMoney(amount)
			transactions = append(transactions, amount)
		case 2:
			fmt.Print("Amount: ")
			fmt.Scanln(&amount)
			wallet.SpendMoney(amount)
			transactions = append(transactions, -amount)
		case 3:
			fmt.Println("Balance:", wallet.GetBalance())
		case 0:
			return
		}
	}
}
