package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(message string) {
	for i := 0; ; i++ {
		fmt.Println(message, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	boring("boring !!!")
}
