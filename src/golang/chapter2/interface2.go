package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Gau gau"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meo meo"
}

func main() {
	// d := Dog{}
	// fmt.Println(d.Speak())
	animals := []Animal{Dog{}, Cat{}}
	for _, item := range animals {
		fmt.Println(item.Speak())
	}
}
