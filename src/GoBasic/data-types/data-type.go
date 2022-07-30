package dataTypes

import "fmt"

//Global variable declaration
var (
	m int
	n int
)

func DataTypes() {
	/*
		Following is the list of predeclared identifiers:

		For Constants:
		--------------
		true, false, iota, nil


		For Types:
		----------
		int, int8, int16, int32, int64, uint,
		uint8, uint16, uint32, uint64, uintptr,
		float32, float64, complex128, complex64,
		bool, byte, rune, string, error


		For Functions:
		--------------
		make, len, cap, new, append, copy, close,
		delete, complex, real, imag, panic, recover


		There are total 25 keywords present in the Go language as follows:
		------------------------------------------------------------------
		break	case	chan	const	continue	default
		defer	else	fallthrough		for		func		go
		goto	if		import		interface	map		package
		range	return	select		struct		switch		type	var


		Data-Type	Description
		------------------------
		int8		8-bit signed integer
		int16		16-bit signed integer
		int32		32-bit signed integer
		int64		64-bit signed integer
		uint8		8-bit unsigned integer
		uint16		16-bit unsigned integer
		uint32		32-bit unsigned integer
		uint64		64-bit unsigned integer
		int			Both in and uint contain same size, either 32 or 64 bit.
		uint		Both in and uint contain same size, either 32 or 64 bit.
		rune		It is a synonym of int32 and also represent Unicode code points.
		byte		It is a synonym of int8.
		uintptr		It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.
		float32		32-bit IEEE 754 floating-point number
		float64		64-bit IEEE 754 floating-point number
		complex64	Complex numbers which contain float32 as a real and imaginary component.
		complex128	Complex numbers which contain float64 as a real and imaginary component.
		The boolean data type represents only one bit of information either true or false.
		The string data type represents a sequence of Unicode code points.

	*/

	var x int = 1 // Integer Data Type
	var y int     //  Integer Data Type
	fmt.Println(x)
	fmt.Println(y)

	var a, b, c = 5.25, 25.25, 14.15 // Multiple float32 variable declaration
	fmt.Println(a, b, c)

	city := "Berlin"     // String variable declaration
	Country := "Germany" // Variable names are case sensitive
	fmt.Println(city)
	fmt.Println(Country) // Variable names are case sensitive

	food, drink, price := "Pizza", "Pepsi", 125 // Multiple type of variable declaration in same line
	fmt.Println(food, drink, price)
	m, n = 1, 2
	fmt.Println(m, n)
}
