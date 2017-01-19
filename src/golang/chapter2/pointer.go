package main

import "fmt"

func main() {
	x := 5
	one(&x)
	fmt.Println(x)
	a := 9
	b := 4
	swap(&a, &b)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

func zero(input int) {
	input = 0
}

func one(input *int) {
	*input = 1
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
