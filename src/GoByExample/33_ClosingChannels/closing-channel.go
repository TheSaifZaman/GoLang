package closingChannels

import "fmt"


//Closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers.
func CloseChannels() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	//Here’s the worker goroutine. It repeatedly receives from jobs with j, more := <-jobs. In this special 2-value form of receive, the more value will be false if jobs has been closed and all values in the channel have already been received. We use this to notify on done when we’ve worked all our jobs.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received Job", j)
			} else {
				fmt.Println("Received All Jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Sent Job", j)
	}
	close(jobs)
	fmt.Println("Sent All Jobs.")
	<-done
}

/*
Result: 

sent job 1
received job 1
sent job 2
received job 2
sent job 3
received job 3
sent all jobs
received all jobs
*/