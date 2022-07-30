package functions

import "fmt"

// SimpleFunction prints a message
func SimpleFunction() {
	fmt.Println("Hello World")
}

// Function accepting arguments
func Add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

// Function with int as return type
func AddWithReturnType(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func Rectangle(l int, b int) (area int) {
	var parameter int
	parameter = 2 * (l + b)
	fmt.Println("Parameter: ", parameter)

	area = l * b
	return // Return statement without specify variable name
}

func RectangleWithMultipleReturn(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

// Golang Passing Address to a Function

func UpdateWithPassingAddress(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	return
}

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func Sum(x, y int) int {
	return x + y
}

func PartialSum(x int) func(int) int {
	return func(y int) int {
		return Sum(x, y)
	}
}

func HigerOrderFunction() {
	partial := PartialSum(3)
	fmt.Println(partial(7))
}


func SquareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}
func ReturningFunctionsFromOtherFunctions() {
	// 5*5 + 6*6 + 7*7
	fmt.Println(SquareSum(5)(6)(7))
}


// User Defined Function Types in Golang
type First func(int) int
type Second func(int) First

func SquareSumWithUserDefinedFunctionType(x int) Second {
	return func(y int) First {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func FunctionUserDefinedFunctionType() {
	// 5*5 + 6*6 + 7*7
	fmt.Println(SquareSumWithUserDefinedFunctionType(5)(6)(7))
}

func Functions() {

	SimpleFunction()

	// Passing arguments
	Add(20, 30)

	fmt.Println(AddWithReturnType(20, 30))

	fmt.Println("Area: ", Rectangle(20, 30))

	var a, p int
	a, p = RectangleWithMultipleReturn(20, 30)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)

	/*
		Naming Conventions for Golang Functions
			=> A name must begin with a letter, and can have any number of additional letters and numbers.
			=> A function name cannot start with a number.
			=> A function name cannot contain spaces.
			=> If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package.
			=> If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
			=> function names are case-sensitive (car, Car and CAR are three different variables).
	*/

	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)

	UpdateWithPassingAddress(&age, &text)

	fmt.Println("After :", text, age)

	// Anonymous Functions in Golang
	fmt.Println(area(20, 30))

	// Passing arguments to anonymous functions.
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)

	// Function defined to accept a parameter and return value.
	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)

	/*

		Closures Functions in Golang
		Closures are a special case of anonymous functions. Closures are anonymous functions which access the variables defined outside the body of the function.

		Example
		Anonymous function accessing the variable defined outside body.
	*/

	length := 20
	width := 30

	func() {
		var area int
		area = length * width
		fmt.Println(area)
	}()

	// Anonymous function accessing variable on each iteration of loop inside function body.
	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
	}

	/*
	Higher Order Functions in Golang
	A Higher-Order function is a function that receives a function as an argument or returns the function as output.

	Higher order functions are functions that operate on other functions, either by taking them as arguments or by returning them.

	Passing Functions as Arguments to other Functions
	*/
	HigerOrderFunction()

	ReturningFunctionsFromOtherFunctions()

	FunctionUserDefinedFunctionType()

}
