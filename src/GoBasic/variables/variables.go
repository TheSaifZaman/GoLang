package variables

import (
	"fmt"
	"reflect"
)

func VariablesWithInit() {
	var integer int = 10
	var str string = "Europe"

	fmt.Println(integer)
	fmt.Println(str)
}

func VariablesWithOutInit() {
	var i int
	var s string

	i = 10
	s = "Canada"

	fmt.Println(i)
	fmt.Println(s)
}

func VariableWithTypeInteface() {
	var i = 10
	var s = "Canada"

	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(s))
}

func MultipleVariableDeclaration() {
	var fname, lname string = "John", "Doe"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 2000

	fmt.Println(fname + lname)
	fmt.Println(m + n + o)
	fmt.Println(item, "-", price)
}

func ShortVariableDeclaration() {
	name := "John Doe"
	fmt.Println(reflect.TypeOf(name))
}

var s = "Japan"

func ScoopVaribale() {
	fmt.Println(s)
	x := true

	if x {
		y := 1
		if x != false {
			fmt.Println(s)
			fmt.Println(x)
			fmt.Println(y)
		}
	}
	fmt.Println(x)
}

func ZeroValue() {
	var quantity float32
	fmt.Println(quantity)

	var price int16
	fmt.Println(price)

	var product string
	fmt.Println(product)

	var inStock bool
	fmt.Println(inStock)
}

var (
	product  = "Mobile"
	quantity = 50
	price    = 50.50
	inStock  = true
)

func VariableDeclarationBlock() {

	fmt.Println(quantity)
	fmt.Println(price)
	fmt.Println(product)
	fmt.Println(inStock)

}
func Variables() {
	/*
	A variable is a piece of storage containing data temporarily to work with it.
	The most general form to declare a variable in Golang uses the var keyword,
	an explicit type, and an assignment.
		// var name type = expression
	*/
	// Declaration with initialization
	// If one know beforehand what a variable's value will be, one can declare
	// variables and assign them values on the same line.
	VariablesWithInit()

	// Declaration without initialization
	/*
		The keyword var is used for declaring variables followed by the desired name
		 and the type of value the variable will hold.
	*/
	VariablesWithOutInit()

	// Declaration with type inference
	/*
		one can omit the variable type from the declaration, when one are assigning
		a value to a variable at the time of declaration. The type of the value
		assigned to the variable will be used as the type of that variable.
	*/
	VariableWithTypeInteface()

	// Declaration of multiple variables
	// Golang allows one to assign values to multiple variables in one line.
	MultipleVariableDeclaration()

	// Short Variable Declaration in Golang
	/*
		The := short variable assignment operator indicates that short variable declaration is being used.
		There is no need to use the var keyword or declare the variable type.
	*/
	ShortVariableDeclaration()

	// Scope of Golang Variables Defined by Brace Brackets
	/*
		Golang uses lexical scoping based on code blocks to determine the scope of
		variables. Inner block can access its outer block defined variables,
		but outer block cannot access inner block defined variables.
	*/
	ScoopVaribale()

	/*

		Naming Conventions for Golang Variables
		These are the following rules for naming a Golang variable:

		A name must begin with a letter, and can have any number of additional letters and numbers.
		A variable name cannot start with a number.
		A variable name cannot contain spaces.
		If the name of a variable begins with a lower-case letter, it can only be accessed within the current package this is considered as unexported variables.
		If the name of a variable begins with a capital letter, it can be accessed from packages outside the current package one this is considered as exported variables.
		If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
		Variable names are case-sensitive (car, Car and CAR are three different variables).

	*/

	// Zero Values
	// If one declare a variable without assigning it a value, Golang will automatically bind a default value (or a zero-value) to the variable.
	ZeroValue()

	// Golang Variable Declaration Block
	// Variables declaration can be grouped together into blocks for greater readability and code quality.
	VariableDeclarationBlock()

}
