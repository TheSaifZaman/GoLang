package clousers

import "fmt"
//Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func Closures() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt = intSeq()

	fmt.Println(nextInt())
}
