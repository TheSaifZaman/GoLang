package goSelect

import (
	"fmt"
	"time"
)
//Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.
func Select() {
	channel_1 := make(chan string)
	channel_2 := make(chan string)

	//Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.
	go func() {
		time.Sleep(time.Second * 1)
		channel_1 <- "One"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		channel_2 <- "Two"
	}()

	//We’ll use select to await both of these values simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case message_1 := <-channel_1:
			fmt.Println("Received: ",message_1)
		case message_2 := <-channel_2:
			fmt.Println("Received: ",message_2)
		}
	}
	//Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.
}