package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
}

func (u User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type Namer interface {
	Name() string
}

type Customer struct {
	Id       int
	FullName string
}

func (c *Customer) Name() string {
	return c.FullName
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

func main() {
	u := User{"Matt", "Aimonetti"}
	fmt.Println(Greet(u))

	c := &Customer{1, "Tony Stark"}
	fmt.Println(Greet(c))
}
