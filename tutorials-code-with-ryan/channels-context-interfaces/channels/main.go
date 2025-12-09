package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Version 1
// func main() {
// 	//Version 1 approach that led to deadlock without buffer
// 	// dataChan := make(chan int) // any type could be sent into th channel. THis is unbuffered channel. Think of it as a portal

// 	dataChan := make(chan int, 3) // last value is size of channel, buffered, note comma before the size of the buffer

// 	go func() { // using go func, concurrency, for one thread to add items to the channel
// 		dataChan <- 69 // send a value into the channel, of the correct type.
// 	}()

// 	n := <-dataChan // send value from channel, back, or set value of 'n' as the value from the channel
// 	// for this channel, without buffering, no space in the channel, and needs to be read immediately as it is being added
// 	//
// 	fmt.Printf("n = %d\n", n)
// 	//fatal error: all goroutines are asleep - deadlock!
// 	// goroutine 1 [chan send]:
// 	// main.main()
// 	// 		C:/Users/jdc/Documents/GithubPersonal/jdc.code.learn.golang/tutorials-code-with-ryan/channels-context-interfaces/channels/main.go:9 +0x36
// 	// exit status 2

// }

func DoWork() int {
	fmt.Print("Doing Work")
	return rand.Intn(100)
}

// Version 2:
func main() {
	dataChan := make(chan int)

	go func() {
		wg := sync.WaitGroup{}

		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() { // starting on a new thread
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		} // Running in background thread
		wg.Wait() // waiting for all wait groups to complete and close
		close(dataChan)
	}() // remember the closing parenthesis after the go func.

	for n := range dataChan {
		fmt.Printf("n = %d\n", n)
	}

}

// General notes: Channels, as communication between threads, meant to be used in a concurrent model
