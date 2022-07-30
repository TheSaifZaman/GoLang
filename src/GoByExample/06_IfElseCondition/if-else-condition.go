package ifElseCondition

import "fmt"

func IfElseCondition() {

    if 12%4 == 0 {
        fmt.Println("12 is divisible by 4")
    }

    if 11%2 == 0 {
        fmt.Println("11 is even")
    } else {
        fmt.Println("11 is odd")
    }

    if num := 18; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

	/*
	There is no ternary if in Go, so you’ll need to use a full
	if statement even for basic conditions.
	*/
}