package main

import "fmt"

type SalaryCalculator interface {
	CalculatorSalary() int
	ShowID() int
	ShowName() string
	ShowBasicPay() int
}

type Permanent struct {
	ID       int
	Name     string
	BasicPay int
	PF       int
}

type Contract struct {
	ID       int
	Name     string
	BasicPay int
}

func (p Permanent) CalculatorSalary() int {
	return p.BasicPay + p.PF
}

func (p Permanent) ShowID() int {
	return p.ID
}

func (p Permanent) ShowBasicPay() int {
	return p.BasicPay
}

func (p Permanent) ShowName() string {
	return p.Name
}

func (c Contract) CalculatorSalary() int {
	return c.BasicPay
}

func (c Contract) ShowID() int {
	return c.ID
}

func (c Contract) ShowBasicPay() int {
	return c.BasicPay
}

func (c Contract) ShowName() string {
	return c.Name
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculatorSalary()
	}

	fmt.Printf("Total expense per month $%d \n", expense)
}

func showInfo(s []SalaryCalculator) {
	for _, v := range s {
		fmt.Println("-----------------------")
		fmt.Println("ID:", v.ShowID())
		fmt.Println("Name:", v.ShowName())
		fmt.Println("Basic Pay:", v.ShowBasicPay())
		fmt.Println("-----------------------")
	}
}

func main() {
	pem1 := Permanent{
		ID:       1,
		Name:     "Hello",
		BasicPay: 5000,
		PF:       20,
	}

	pem2 := Permanent{
		ID:       2,
		Name:     "World",
		BasicPay: 6000,
		PF:       30,
	}

	cemp1 := Contract{
		ID:       3,
		Name:     "Ready",
		BasicPay: 3000,
	}

	employees := []SalaryCalculator{pem1, pem2, cemp1}

	totalExpense(employees)
	showInfo(employees)
}
