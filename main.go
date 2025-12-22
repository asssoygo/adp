package main

import (
	"fmt"

	"github.com/Asylkhan/Assignment1/employee"
	"github.com/Asylkhan/Assignment1/gym"
	"github.com/Asylkhan/Assignment1/hotel"
	"github.com/Asylkhan/Assignment1/wallet"
)

func main() {
	// ---- HOTEL ----
	h := hotel.Hotel{Rooms: make(map[string]hotel.Room)}
	h.AddRoom(hotel.Room{"101", "Single", 100, false})
	h.AddRoom(hotel.Room{"102", "Double", 150, false})
	h.CheckIn("101")
	h.ListVacantRooms()

	// ---- EMPLOYEE ----
	employees := []employee.Employee{
		employee.FullTime{MonthlySalary: 2000, BonusRate: 0.1},
		employee.PartTime{HourlyRate: 10, HoursWorked: 80},
	}

	for _, e := range employees {
		fmt.Println("Salary:", e.CalculateSalary())
	}

	// ---- GYM ----
	g := gym.Gym{Members: make(map[uint64]gym.Member)}
	g.AddMember(1, gym.BasicMember{Name: "Ali"})
	g.ListMembers()

	// ---- WALLET ----
	w := wallet.Wallet{}
	w.AddMoney(500)
	w.SpendMoney(120)
	fmt.Println("Wallet balance:", w.GetBalance())
}
