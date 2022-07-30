package intToString

import (
	"fmt"
	"reflect"
	"strconv"
)

func IntToStringuUSingFormatString() {
	/*
			FormatInt() Method
		You can use the strconv package's FormatInt() function to convert the int into an string value.
		FormatInt returns the string representation of i in the given base,
		for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
		for digit values >= 10.

		Syntax
		func FormatInt(i int64, base int) string
	*/
	var i int64 = 125
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i)

	var s string = strconv.FormatInt(i, 10)
	fmt.Println(reflect.TypeOf(s))

	fmt.Println("Base 10 value of s:", s)

	s = strconv.FormatInt(i, 8)
	fmt.Println("Base 8 value of s:", s)

	s = strconv.FormatInt(i, 16)
	fmt.Println("Base 16 value of s:", s)

	s = strconv.FormatInt(i, 32)
	fmt.Println("Base 32 value of s:", s)
}

func IntToStringUsingFmtSprintf() {
	/*
	   fmt.Sprintf() Method
	   Sprintf formats according to a format specifier and returns the resulting string.
	   Here, a is of Interface type hence you can use this method to convert any type to string.

	   Syntax
	   func Sprintf(format string, a ...interface{}) string
	*/
	b := 1225
	fmt.Println(reflect.TypeOf(b))

	s := fmt.Sprintf("%v", b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}
