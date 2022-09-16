package main

import "fmt"

func main() {
	developerFactory := NewEmployeeFactory("Developer", 100000)
	managerFactory := NewEmployeeFactory("Manager", 150000)
	developer := developerFactory("Fulano")
	fmt.Println(developer)
	manager := managerFactory("Ciclano")
	fmt.Println(manager)
	bossFactory := NewEmployeeFactory2("CTO", 300000)
	boss := bossFactory.Create("Mateus")
	fmt.Println(boss)
	//Uma vez gerado o Factory seja de developer, manager, eu não consigo alterar
	// Os valores padrões
	// Mas com o de struct sim
	bossFactory.AnualIncome = 500000
	bossPromotion := bossFactory.Create("Mateus")
	fmt.Println(bossPromotion)
}

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// Observação: functional approach parece ser melhor mesmo, o structural é mais
// se eu tiver a necessidade de modificar algum valor depois e seria bom
// fazer uso de interfaces quando for utilizar ele, pra esse método que modifica algo.
// what if we want factories for specific roles?
// functional approach
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// structural approach
type EmployeeFactory struct {
	Position    string
	AnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}
