package stringToBool

import (
	"fmt"
	"strconv"
)

/*
String to Boolean Conversion
Package strconv is imported to perform conversions to and from string.
ParseBool returns the boolean value represented by the string.
It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
Any other value returns an error.
*/

func StringToBool() {
	s1 := "true"
	b1, _ := strconv.ParseBool(s1)
	fmt.Printf("%T, %v\n", b1, b1)

	s2 := "t"
	b2, _ := strconv.ParseBool(s2)
	fmt.Printf("%T, %v\n", b2, b2)

	s3 := "0"
	b3, _ := strconv.ParseBool(s3)
	fmt.Printf("%T, %v\n", b3, b3)

	s4 := "F"
	b4, _ := strconv.ParseBool(s4)
	fmt.Printf("%T, %v\n", b4, b4)
}
