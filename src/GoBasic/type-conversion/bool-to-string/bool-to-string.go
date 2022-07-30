package boolToString

import (
	"fmt"
	"reflect"
	"strconv"
)

func BoolToStringUsingFormatBool() {
	/*
	FormatBool function
	You can use the strconv package's FormatBool() function to convert the boolean into an string value.
	FormatBool returns "true" or "false" according to the value of b.

	Syntax
	func FormatBool(b bool) string
	*/

	var b bool = true
	fmt.Println(reflect.TypeOf(b))

	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
}

func BoolToStringUsingFmtSprintf() {
	/*
	fmt.Sprintf() method
	Sprintf formats according to a format specifier and returns the resulting string.
	Here, a is of Interface type hence you can use this method to convert any type to string.

	Syntax
	func Sprintf(format string, a ...interface{}) string

	*/
	b := true
	s := fmt.Sprintf("%v", b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}
