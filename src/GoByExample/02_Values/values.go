package values

import "fmt"

func Values() {

    fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println("And Operations:")
    fmt.Println("(true && true) = ",true && true)
    fmt.Println("(true && false) = ",true && false)
    fmt.Println("(false && true) = ",false && true)
    fmt.Println("(false && false) = ",false && false)

	fmt.Println("Or Operations:")
    fmt.Println("(true || true) = ",true || true)
    fmt.Println("(true || false) = ",true || false)
    fmt.Println("(false || true) = ",false || true)
    fmt.Println("(false || false) = ",false || false)

	fmt.Println("Not Operations:")
    fmt.Println("!true = ",!true)
    fmt.Println("!false = ",!false)
}