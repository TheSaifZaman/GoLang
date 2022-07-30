package operators

import "fmt"

/*
Operators
An operator is a symbol that tells the compiler to perform certain actions. 
The following lists describe the different operators used in Golang.

Arithmetic Operators
Assignment Operators
Comparison Operators
Logical Operators
Bitwise Operators
*/


/*
Arithmetic Operators in Go Programming Language
The arithmetic operators are used to perform common arithmetical operations, such as addition, subtraction, multiplication etc.

Here's a complete list of Golang's arithmetic operators:

Operator	Description		Example		Result
+			Addition		x + y		Sum of x and y
-			Subtraction		x - y		Subtracts one value from another
*			Multiplication	x * y		Multiplies two values
/			Division		x / y		Quotient of x and y
%			Modulus			x % y		Remainder of x divided by y
++			Increment		x++			Increases the value of a variable by 1
--			Decrement		x--			Decreases the value of a variable by 1

The following example will show you these arithmetic operators in action:
*/


func ArithmeticOperators() {
	var x, y = 35, 7

	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("x / y = %d\n", x/y)
	fmt.Printf("x mod y = %d\n", x%y)

	x++
	fmt.Printf("x++ = %d\n", x)

	y--
	fmt.Printf("y-- = %d\n", y)
}

/*
Assignment Operators in Go Programming Language
The assignment operators are used to assign values to variables

Assignment	Description					Example
x = y		Assign						x = y
x += y		Add and assign				x = x + y
x -= y		Subtract and assign			x = x - y
x *= y		Multiply and assign			x = x * y
x /= y		Divide and assign quotient	x = x / y
x %= y		Divide and assign modulus	x = x % y

The following example will show you these assignment operators in action:
*/


func AssignmentOperators() {
	var x, y = 15, 25
	x = y
	fmt.Println("= ", x)

	x = 15
	x += y
	fmt.Println("+=", x)

	x = 50
	x -= y
	fmt.Println("-=", x)

	x = 2
	x *= y
	fmt.Println("*=", x)

	x = 100
	x /= y
	fmt.Println("/=", x)

	x = 40
	x %= y
	fmt.Println("%=", x)
}

/*
Comparison Operators in Go Programming Language
Comparison operators are used to compare two values.

Operator	Name						Example		Result
==			Equal						x == y		True if x is equal to y
!=			Not equal					x != y		True if x is not equal to y
<			Less than					x < y		True if x is less than y
<=			Less than or equal to		x <= y		True if x is less than or equal to y
>			Greater than				x > y		True if x is greater than y
>=			Greater than or equal to	x >= y		True if x is greater than or equal to y

The following example will show you these comparison operators in action:
*/

func ComparisonOperators() {
	var x, y = 15, 25

	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x > y)
	fmt.Println(x >= y)
}

/*
Logical Operators in Go Programming Language
Logical operators are used to determine the logic between variables or values.

Operator	Name		Description													Example
&&			Logical And	Returns true if both statements are true					x < y && x > z
||			Logical Or	Returns true if one of the statements is true				x < y || x > z
!			Logical Not	Reverse the result, returns false if the result is true		!(x == y && x > z)

The following example will show you these logical operators in action:
*/

func LogicalOperators() {
	var x, y, z = 10, 20, 30

	fmt.Println(x < y && x > z)
	fmt.Println(x < y || x > z)
	fmt.Println(!(x == y && x > z))
}

/*
Bitwise Operators in Go Programming Language
Bitwise operators are used to compare (binary) numbers.

Operator	Name					Description
&			AND						Sets each bit to 1 if both bits are 1
|			OR						Sets each bit to 1 if one of two bits is 1
^			XOR						Sets each bit to 1 if only one of two bits is 1
<<			Zero fill left shift	Shift left by pushing zeros in from the right and let the leftmost bits fall off
>>			Signed right shift		hift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off

The following example will show you these bitwise operators in action:
*/


func BitwiseOperators() {
	var x uint = 9  //0000 1001
	var y uint = 65 //0100 0001
	var z uint

	z = x & y
	fmt.Println("x & y  =", z)

	z = x | y
	fmt.Println("x | y  =", z)

	z = x ^ y
	fmt.Println("x ^ y  =", z)

	z = x << 1
	fmt.Println("x << 1 =", z)

	z = x >> 1
	fmt.Println("x >> 1 =", z)
}