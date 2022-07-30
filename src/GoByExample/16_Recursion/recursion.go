package recursion

import "fmt"

func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func Recursion() {
	fmt.Println(fact(7))
}