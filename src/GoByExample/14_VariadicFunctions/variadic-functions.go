package variadicFunctions

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0

    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
  
func VariadicFunc(messages ...interface{}) {
    for _, i := range messages {
      fmt.Println(i)
    }
  }

func VariadicFunctions() {

    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...)

    VariadicFunc("hello", "Change", "the ", "World using Golang", 1, 2, 3, 4)
  }
