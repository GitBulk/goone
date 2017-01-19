package main

import (
	"fmt"
	"math"
)

func main() {
	var tom Person
	tom.name = "Tom Cat"
	tom.age = 5
	jerry := Person{age: 3, name: "Jerry Mouse"}
	olderPerson, diff := Older(tom, jerry)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, jerry.name, olderPerson.name, diff)

	mark := Student{specialty: "Computer Science", human: Human{"Mark", 25, 120}}
	fmt.Println("His name is", mark.human.name)
	fmt.Println("His specialty is", mark.specialty)
	fmt.Println("His weight is", mark.human.weight)
	mark.specialty = "AI"
	fmt.Println("Mark changed his specialty")
	fmt.Println("His specialty is", mark.specialty)

	fmt.Println("-- Embeded field --")
	toan := Employee{personInfo: PersonInfo{name: "toan tran", age: 10, phone: "123"}, specially: "IT", phone: "456"}
	fmt.Println("Personal's phone", toan.personInfo.phone)
	fmt.Println("Employee's phone", toan.phone)
	r1 := Rectangle{width: 3, height: 4}
	//	fmt.Println("Area of rectangle", area(r1))
	fmt.Println("Area of rectangle is", r1.area())

	c1 := Circle{radius: 4}
	fmt.Println("Area of circle is", c1.area())
}

func area(rec Rectangle) float64 {
	return rec.width * rec.height
}

func (rec Rectangle) area() float64 {
	return rec.width * rec.height
}

type Circle struct {
	radius float64
}

func (cir Circle) area() float64 {
	return cir.radius * cir.radius * math.Pi
}

type Rectangle struct {
	width, height float64
}

type PersonInfo struct {
	name  string
	age   int
	phone string
}

type Employee struct {
	personInfo PersonInfo
	specially  string
	phone      string
}

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	human     Human // embedded field
	specialty string
}

type Person struct {
	name string
	age  int
}

func Older(p1, p2 Person) (Person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}
