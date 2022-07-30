package loops

import "fmt"

func TraditionalForStatement() {
 
	k := 1
	for ; k <= 10; k++ {
		fmt.Println(k)
	}
 
	k = 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}
 
	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}
}

func ForRangeStatement() {
 
	// Example 1
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}
 
	// Example 2
	for key := range strDict {
		fmt.Println(key)
	}
 
	// Example 3
	for _, value := range strDict {
		fmt.Println(value)
	}
}

func RangeLoopOverString() {
	for range "Hello" {
		fmt.Println("Hello")
	}
}

func InfiniteLoop() {
	// i := 5
	// for {
	// 	fmt.Println("Hello")
	// 	if i == 10 {
	// 		break
	// 	}
	// 	i++
	// }
}


func SwitchCaseWithBreakInForLoop() {

    testLoop:for val := 1; val < 7; val++ {
        fmt.Printf("%d", val)
        switch {
        case val == 1:
            fmt.Println("->Start")
        case val == 5:
            fmt.Println("->Break")
            break testLoop
        case val > 2:
            fmt.Println("->Running")
            break 
        default:
            fmt.Println("->Progress")
        }
    }
}