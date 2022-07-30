package rangeOverChannels

import "fmt"


//In a previous example we saw how for and range provide iteration over basic data structures. We can also use this syntax to iterate over values received from a channel.
func RangeOverChannels() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	//This range iterates over each element as it’s received from queue. Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	for element := range queue {
		fmt.Println(element)
	}
}

/*
Result:
one
two
*/