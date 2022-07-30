package constants

import (
    "fmt"
    "math"
)

const const_str string = "constant"

func Constants() {
    fmt.Println(const_str)

    const n = 100000000

    const d = 2e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
}