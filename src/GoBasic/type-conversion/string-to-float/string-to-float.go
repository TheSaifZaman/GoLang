package stringToFloat

/*
ParseFloat() Function
ParseFloat converts the string s to a floating-point number with the precision specified by 
bitSize: 32 for float32, or 64 for float64. When bitSize=32, 
the result still has type float64, but it will be convertible to float32 
without changing its value.

Syntax
func ParseFloat(s string, bitSize int) (float64, error)
*/

import (
	"fmt"
	"reflect"
	"strconv"
)

func StringToFloat() {
	s := "3.1415926535"
	f, err := strconv.ParseFloat(s, 8)
	fmt.Println(f, err, reflect.TypeOf(f))

	s1 := "-3.141"
	f1, err := strconv.ParseFloat(s1, 8)
	fmt.Println(f1, err, reflect.TypeOf(f1))

	s2 := "A-3141X"
	f2, err := strconv.ParseFloat(s2, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f2, err, reflect.TypeOf(f2))
	}
}