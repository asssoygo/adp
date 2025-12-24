package main

import (
	"adp/Employee"
	"adp/Gym"
	"adp/Hotel"
	"adp/Wallet"
	"fmt"
)

func main() {
	for {
		fmt.Println("\n===== MAIN MENU =====")
		fmt.Println("1. Hotel System")
		fmt.Println("2. Employee Salaries")
		fmt.Println("3. Gym Membership")
		fmt.Println("4. Wallet System")
		fmt.Println("0. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			Hotel.RunHotelMenu()
		case 2:
			Employee.RunEmployeeDemo()
		case 3:
			Gym.RunGymDemo()
		case 4:
			Wallet.RunWalletMenu()
		case 0:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
