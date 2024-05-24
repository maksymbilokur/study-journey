/*
Use function is a goroutine to run code asynchronously.
Use channels to exchange data between goroutines
Reads and writes data from/to the channel
Reads data from the channel using for loop
Multiplexes using select
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan bool)
	go someWorker(c)

	select {
	case <-c:
		fmt.Println("get signal from channel")
	case <-time.After(time.Second):
		fmt.Println("got timeout")
	}
}

func someWorker(c chan<- bool) {
	time.Sleep(2 * time.Second)
	c <- true
}
