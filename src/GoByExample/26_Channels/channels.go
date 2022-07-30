package channels

import "fmt"

func Channels() {
	message := make(chan string)

	go func(){message <- "Ping"}()

	msg := <-message
	fmt.Println(msg)
}