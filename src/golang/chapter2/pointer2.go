package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
}

func SetStudent(stu Student) {
	stu.firstName = "toan"
	stu.lastName = "tran"
}

func SetStudent2(stu *Student) {
	stu.firstName = "toan"
	stu.lastName = "tran"
}

func Print(stu Student) {
	fmt.Println("LastName:", stu.lastName)
	fmt.Println("FirstName:", stu.firstName)
}

func main() {
	a := 5
	fmt.Println("a:", a)
	fmt.Println("&a:", &a)

	var b *int = &a
	fmt.Println("b:", b)
	fmt.Println("*b:", *b)

	*b = 42
	fmt.Println("set value for *b")
	fmt.Println("a:", a)

	student := Student{"ten", "ho"}
	SetStudent(student)
	Print(student)

	fmt.Println("Pass pointer student")
	SetStudent2(&student)
	Print(student)
}
