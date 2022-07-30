package variadicFunctions

import (
	"fmt"
	"reflect"
)
/*
A variadic function is a function that accepts a variable number of arguments. 
In Golang, it is possible to pass a varying number of arguments of the same type as referenced in the function signature. 
To declare a variadic function, the type of the final parameter is preceded by an ellipsis, "...", 
which shows that the function may be called with any number of arguments of this type. 
This type of function is useful when you don't know the number of arguments you are passing to the function, 
the best example is built-in Println function of the fmt package which is a variadic function.
*/

func VariadicExample(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])
}

func VariadicExamplePassingMultipleStringArgs(s ...string) {
	fmt.Println(s)
}

func NormalFuncParameterWithVariadicFuncParameter(str string, y ...int) int {

	area := 1

	for _, val := range y {
		if str == "Rectangle" {
			area *= val
		} else if str == "Square" {
			area = val * val
		}
	}
	return area
}

func PassDiffTypesArgumentsWithVariadicFuncParameter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

func VariadicFunction() {
	VariadicExample("red", "blue", "green", "yellow")
	VariadicExamplePassingMultipleStringArgs()
	VariadicExamplePassingMultipleStringArgs("red", "blue")
	VariadicExamplePassingMultipleStringArgs("red", "blue", "green")
	VariadicExamplePassingMultipleStringArgs("red", "blue", "green", "yellow")

	fmt.Println(NormalFuncParameterWithVariadicFuncParameter("Rectangle", 20, 30))
	fmt.Println(NormalFuncParameterWithVariadicFuncParameter("Square", 20))

	PassDiffTypesArgumentsWithVariadicFuncParameter(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
	map[string]int{"apple": 23, "tomato": 13})
}
