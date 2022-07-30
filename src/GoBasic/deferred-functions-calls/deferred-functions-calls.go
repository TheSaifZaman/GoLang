package deferredFucntionsCalls

import "fmt"

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func DeferredFucntionsCalls() {
	defer second()
	DeferredFuncs()
	first()
}

/*
This has various advantages:
It keeps our Close call near our Open call so it's easier to understand.
If our function had multiple return statements (perhaps one in an if and one in an else), Close will happen before both of them.
Deferred Functions are run even if a runtime panic occurs.
Deferred functions are executed in LIFO order, so the above code prints: 4 3 2 1 0.
You can put multiple functions on the "deferred list", like this example.
*/
func DeferredFuncs() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}