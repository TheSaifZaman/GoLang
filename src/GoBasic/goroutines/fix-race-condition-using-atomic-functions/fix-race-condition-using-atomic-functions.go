/*
Race conditions occur due to unsynchronized access to shared resource and attempt to read and write to that resource at the same time.

Atomic functions provide low-level locking mechanisms for synchronizing access to integers and pointers. 
Atomic functions generally used to fix the race condition.

The functions in the atomic under sync packages provides support to synchronize goroutines by locking access to shared resources.
*/

package fixRaceConditionUsingAtomicFunctions

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int32          // counter is a variable incremented by all goroutines.
	wg      sync.WaitGroup // wg is used to wait for the program to finish.
)


func increment(name string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for range name {
		atomic.AddInt32(&counter, 1)
		runtime.Gosched() // Yield the thread and be placed back in queue.
	}
}

func FixRaceConditionUsingAtomicFunctions() {
	wg.Add(3) // Add a count of two, one for each goroutine.

	go increment("Python")
	go increment("Java")
	go increment("Golang")

	wg.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter)

}

/*
Output
The AddInt32 function from the atomic package synchronizes the adding of integer values by enforcing that only one goroutine can perform and complete this add operation at a time. When goroutines attempt to call any atomic function, they're automatically synchronized against the variable that's referenced.

Note if you replace the code line atomic.AddInt32(&counter, 1) with counter++, then you will see the below output-

C:\Golang\goroutines>go run -race main.go
==================
WARNING: DATA RACE
Read at 0x0000006072b0 by goroutine 7:
  main.increment()
      C:/Golang/goroutines/main.go:31 +0x76

Previous write at 0x0000006072b0 by goroutine 8:
  main.increment()
      C:/Golang/goroutines/main.go:31 +0x90

Goroutine 7 (running) created at:
  main.main()
      C:/Golang/goroutines/main.go:18 +0x7e

Goroutine 8 (running) created at:
  main.main()
      C:/Golang/goroutines/main.go:19 +0x96
==================
Counter: 15
Found 1 data race(s)
exit status 66

C:\Golang\goroutines>

Output

C:\Golang\goroutines>go run -race main.go
Counter: 15

*/