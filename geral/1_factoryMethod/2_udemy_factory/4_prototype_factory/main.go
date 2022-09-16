package main

import "fmt"

func main() {
	n := NewEmployee(Manager)
	n.Name = "Mateus"
	fmt.Println(n)
}

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

// functional
func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "Developer", 80000}
	case Manager:
		return &Employee{"", "Manager", 100000}
	default:
		return &Employee{"", "Client", 0}
	}
}
