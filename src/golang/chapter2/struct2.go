package main

func main() {
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
