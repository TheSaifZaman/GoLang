package forLoop

import "fmt"

func ForLoop() {

	// Single condition
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

	// Classical For
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

	// Without a condition
    for {
        fmt.Println("loop")
        break
    }
	
	// With Continue
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}