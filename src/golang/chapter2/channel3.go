package main

import (
	"fmt"
	"time"
)

func Runner(baton chan int) {
	var newRunner int

	// wait to receive the baton
	runner := <-baton

	fmt.Printf("Runner %d Running with baton\n", runner)

	// new runner to the line
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line\n", newRunner)
		go Runner(baton)
	}

	// running around the track
	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		return
	}

	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)
	baton <- newRunner
}

func main() {
	// Create an unbuffered channel
	baton := make(chan int)

	// First runner to his mark
	go Runner(baton)

	// start the Race
	baton <- 1

	//  give the runners time to race
	time.Sleep(500 * time.Millisecond)
}
