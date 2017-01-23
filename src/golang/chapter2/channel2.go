package main

import (
	"fmt"
	"strconv"
	"time"
)

var i int

// http://golangtutorials.blogspot.com/2011/06/channels-in-go.html

func makeCakeAndSend(cs chan string) {
	i = i + 1
	cakeName := "Strawberry Cake " + strconv.Itoa(i)
	fmt.Println("Making a cake and sending ...", cakeName)
	cs <- cakeName // put cake to channel
}

func receiveCakeAndPack(cs chan string) {
	cake := <-cs // get whatever cake is on the channel
	fmt.Println("Packing received cake:", cake)
}

func main() {
	channel := make(chan string)
	for i := 0; i < 5; i++ {
		go makeCakeAndSend(channel)
		go receiveCakeAndPack(channel)

		//sleep for a while so that the program doesn’t exit immediately and output is clear for illustration
		time.Sleep(1 * 1e9)
	}
	// var input string
	// fmt.Scanf("%s", &input)
}
