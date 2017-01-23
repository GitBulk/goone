package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doWork(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go doWork(i)
	}

	fmt.Println("7777")
	var input string
	fmt.Scanln(&input)
}
