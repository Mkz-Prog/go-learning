package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	name string
	owner string
}

type Cat struct {
	name string
	address string
}

func (Dog) Speak() { fmt.Println("Woof!") }
func (Cat) Speak() { fmt.Println("Meow!") }

func main() {
	animals := []Speaker{
		Dog{name: "Бобик", owner: "Вася"},
		Cat{name: "Мурзик", address: "ул. Пушкина"},
	}

	for _, animal := range animals {
		animal.Speak()
	}
}