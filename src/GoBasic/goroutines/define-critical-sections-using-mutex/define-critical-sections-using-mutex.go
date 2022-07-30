/*
A mutex is used to create a critical section around code
that ensures only one goroutine at a time can execute that code section.
*/

package defineCriticalSectionsUsingMutex

import (
	"fmt"
	"sync"
)

var (
	counter int32          // counter is a variable incremented by all goroutines.
	wg      sync.WaitGroup // wg is used to wait for the program to finish.
	mutex   sync.Mutex     // mutex is used to define a critical section of code.
)

func increment(lang string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for i := 0; i < 3; i++ {
		mutex.Lock()
		{
			fmt.Println(lang)
			counter++
		}
		mutex.Unlock()
	}
}

func DefineCriticalSectionsUsingMutex() {
	wg.Add(3) // Add a count of two, one for each goroutine.

	go increment("Python")
	go increment("Go Programming Language")
	go increment("Java")

	wg.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter)

}

/*
Output
C:\Golang\goroutines>go run -race main.go
PHP stands for Hypertext Preprocessor.
PHP stands for Hypertext Preprocessor.
The Go Programming Language, also commonly referred to as Golang
The Go Programming Language, also commonly referred to as Golang
Counter: 4

C:\Golang\goroutines>
A critical section defined by the calls to Lock() and Unlock() protects the actions against 
the counter variable and reading the text of name variable.
*/