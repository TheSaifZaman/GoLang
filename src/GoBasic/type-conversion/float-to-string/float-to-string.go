package floatToString

import (
	"fmt"
	"reflect"
	"strconv"
)

func FloatToStringUsingFormatFloat() {
	/*
	FormatFloat method
You can use the strconv package's FormatFloat() function to convert the float into an string value. 
FormatFloat converts the floating-point number f to a string, according to the format fmt and precision prec. It rounds the result assuming that the original was obtained from a floating-point value of bitSize bits (32 for float32, 64 for float64).

Syntax
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	*/
	var f float64 = 3.1415926535
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(f)

	var s string = strconv.FormatFloat(f, 'E', -1, 32)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s)
}

func FloatToStringUsingFmtSprintf() {
	/*
fmt.Sprintf() method
Sprintf formats according to a format specifier and returns the resulting string. 
Here, a is of Interface type hence you can use this method to convert any type to string.

Syntax
func Sprintf(format string, a ...interface{}) string
	*/
	b := 12.454
	fmt.Println(reflect.TypeOf(b))

	s := fmt.Sprintf("%v", b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}