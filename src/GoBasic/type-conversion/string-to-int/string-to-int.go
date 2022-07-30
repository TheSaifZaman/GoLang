package stringToInt

import (
	"fmt"
	"strconv"
	"reflect"
)

/*
Atoi() Function
You can use the strconv package's Atoi() function to convert the string into an integer value. Atoi stands for ASCII to integer. 
The Atoi() function returns two values: the result of the conversion, and the error (if any).

Syntax
func Atoi(s string) (int, error)
*/

func StringToInt() {
	strVar := "100"
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}

/*
ParseInt() Function
ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i. 
This function accepts a string parameter, convert it into a corresponding int type based on a base parameter. By default, it returns Int64 value.

Syntax
func ParseInt(s string, base int, bitSize int) (i int64, err error)
*/

func StringToIntByParseInt() {
	strVar := "100"

	intVar, err := strconv.ParseInt(strVar, 0, 8)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 16)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 32)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 64)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}

/*
Using fmt.Sscan
The fmt package provides sscan() function which scans string argument and store into variables. 
This function read the string with spaces and assign into consecutive Integer variables.
*/

func StringToIntByFmtSscan(){

		strVar := "100"
		intValue := 0
		_, err := fmt.Sscan(strVar, &intValue)
		fmt.Println(intValue, err, reflect.TypeOf(intValue))
	
}