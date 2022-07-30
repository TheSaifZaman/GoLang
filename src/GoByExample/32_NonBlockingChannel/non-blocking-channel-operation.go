package nonBlockingChannelOperations

import "fmt"

//Closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers.

func NonBlockingChannelOperations() {
	messages := make(chan string)
	signals := make(chan bool)

	//Here’s a non-blocking receive. If a value is available on messages then select will take the <-messages case with that value. If not it will immediately take the default case.
	select {
	case msg := <-messages:
		fmt.Println("Received Message:", msg)
	default:
		fmt.Println("No Message Received")
	}

	//A non-blocking send works similarly. Here msg cannot be sent to the messages channel, because the channel has no buffer and there is no receiver. Therefore the default case is selected.
	msg := "Hello World"
	select {
	case messages <- msg:
		fmt.Println("Sent Message:", msg)
	default:
		fmt.Println("No Message Sent")
	}

	//We can use multiple cases above the default clause to implement a multi-way non-blocking select. Here we attempt non-blocking receives on both messages and signals.
	select {
	case msg := <-messages:
		fmt.Println("Received Message", msg)
	case sig := <-signals:
		fmt.Println("Received Signal", sig)
	default:
		fmt.Println("No Activity")
	}
}