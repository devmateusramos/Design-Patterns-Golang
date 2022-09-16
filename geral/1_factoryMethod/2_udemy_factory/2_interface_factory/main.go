package main

import "fmt"

func main() {
	pessoaNova := NewPerson("Ana", 22)
	pessoaVelha := NewPerson("Fulano", 100)
	pessoaVelha.SayHello()
	pessoaNova.SayHello()

}

type person struct {
	name string
	age  int
}
type oldPerson struct {
	name string
	age  int
}

func (p *oldPerson) SayHello() {
	fmt.Printf("Hello, my name is %s, but i will not tell my age\n", p.name)
}

type Person interface {
	SayHello()
}

func (p *person) SayHello() {
	fmt.Printf("Hello, my name is %s, I am %d years old\n", p.name, p.age)
}

func NewPerson(name string, age int) Person {
	if age > 25 {
		return &oldPerson{name, age}
	}
	return &person{name, age}
}
