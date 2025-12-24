package Employee

import "fmt"

type Employee interface {
	CalculateSalary() float64
	CalculateBonus() float64
}

type FullTime struct {
	MonthlySalary float64
	BonusRate     float64
}

func (f FullTime) CalculateSalary() float64 {
	return f.MonthlySalary
}

func (f FullTime) CalculateBonus() float64 {
	return f.MonthlySalary * f.BonusRate
}

type PartTime struct {
	HourlyRate  float64
	HoursWorked int
}

func (p PartTime) CalculateSalary() float64 {
	return p.HourlyRate * float64(p.HoursWorked)
}

func (p PartTime) CalculateBonus() float64 {
	return 0
}

type Contractor struct {
	ProjectRate       float64
	ProjectsCompleted int
}

func (c Contractor) CalculateSalary() float64 {
	return c.ProjectRate * float64(c.ProjectsCompleted)
}

func (c Contractor) CalculateBonus() float64 {
	return c.ProjectRate * 0.1
}

type Intern struct {
	DailyRate  float64
	DaysWorked int
}

func (i Intern) CalculateSalary() float64 {
	return i.DailyRate * float64(i.DaysWorked)
}

func (i Intern) CalculateBonus() float64 {
	return 0
}

func RunEmployeeDemo() {
	employees := []Employee{
		FullTime{500000, 0.2},
		PartTime{3000, 80},
		Contractor{200000, 2},
		Intern{5000, 20},
	}

	for i, e := range employees {
		fmt.Printf("Employee %d Salary: %.2f Bonus: %.2f\n",
			i+1, e.CalculateSalary(), e.CalculateBonus())
	}
}
