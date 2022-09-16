package main

import "fmt"

func main() {
	pessoaUm := NewPerson("Fulano", 20)
	pessoaDois := NewPerson("Ciclano", 42)
	fmt.Println(pessoaUm.Name)
	fmt.Println(pessoaDois.Name)
}

type Person struct {
	Name  string
	Age   int
	Money int
}

func NewPerson(name string, age int) *Person {
	return &Person{name, age, 1000}
}

// Ao inves de manualmente sempre adicionar 1000 eu crio uma factory function
// que já retorna isso default
// Poderia ter validações também, lógica adiconal etc
