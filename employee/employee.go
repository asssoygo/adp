package employee

type Employee interface {
	CalculateSalary() float64
	GetWorkHours() int
}

type FullTime struct {
	MonthlySalary float64
	BonusRate     float64
}

func (f FullTime) CalculateSalary() float64 {
	return f.MonthlySalary + f.MonthlySalary*f.BonusRate
}

func (f FullTime) GetWorkHours() int {
	return 160
}

type PartTime struct {
	HourlyRate  float64
	HoursWorked int
}

func (p PartTime) CalculateSalary() float64 {
	return float64(p.HoursWorked) * p.HourlyRate
}

func (p PartTime) GetWorkHours() int {
	return p.HoursWorked
}
