package timers

import (
	"fmt"
	"time"
)

func Timers() {
	//Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides a channel that will be notified at that time. This timer will wait 2 seconds.
	timer_1 := time.NewTimer(time.Second * 2)

	//The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired.
	<-timer_1.C
	fmt.Println("Timer 1 Fired")

	//If you just wanted to wait, you could have used time.Sleep. One reason a timer may be useful is that you can cancel the timer before it fires. Here’s an example of that.
	timer_2 := time.NewTimer(time.Second)

	go func() {
		<-timer_2.C
		fmt.Println("Timer 2 Fired")
	}()
	stop_2 := timer_2.Stop()
	if stop_2 {
		fmt.Println("Timer 2 Stopped")
	}
	//Give the timer2 enough time to fire, if it ever was going to, to show it is in fact stopped.
	time.Sleep(2 * time.Second)
}

/*
Result:
Timer 1 fired
Timer 2 stopped
*/