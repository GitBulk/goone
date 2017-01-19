package main

import "fmt"

func main() {
	p := Person{
		name: "toan",
		address: Address{
			city:     "SG",
			district: "1",
			ward:     "Da Kao",
		},
	}

	p.Talk()
	p.Location()

	c := Citizen{
		country: "VN",
		person:  p,
	}
	c.Nationality()

	fmt.Println("-- mutable --")
	m := Mutable{
		a: 3,
		b: 1,
	}
	m.mutate()
	fmt.Println(m)
}

type Mutable struct {
	a int
	b int
}

func (m *Mutable) mutate() {
	m.a = 10
	m.b = 9
}

type Citizen struct {
	country string
	person  Person
}

func (c Citizen) Nationality() {
	fmt.Println(c.person.name, "is a citizen of", c.country)
}

type Address struct {
	city     string
	district string
	ward     string
}

type Person struct {
	name    string
	address Address
}

func (p Person) Talk() {
	fmt.Println("Hi, my name is", p.name)
}

func (p Person) Location() {
	fmt.Println("I’m at", p.address.city, p.address.district, p.address.ward)
}
