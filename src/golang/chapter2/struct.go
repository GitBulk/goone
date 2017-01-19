package main

import "fmt"

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
