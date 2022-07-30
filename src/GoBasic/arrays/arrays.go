package arrays

import (
	"fmt"
	"reflect"
)

/*
An array is a data structure that consists of a collection of elements of a single type or simply you can say a special variable,
which can hold more than one value at a time. The values an array holds are called its elements or items. An array holds a specific
number of elements, and it cannot grow or shrink. Different data types can be handled as elements in arrays such as Int, String, Boolean, and others.
The index of the first element of any dimension of an array is 0, the index of the second element of any array dimension is 1, and so on.
*/

func SimpleArrayDeclaration() {
	var intArray [5]int
	var strArray [5]string

	fmt.Println(reflect.ValueOf(intArray).Kind())
	fmt.Println(reflect.ValueOf(strArray).Kind())
}

func AssignAndAccessIntoArray() {
	var theArray [3]string
	theArray[0] = "India"  // Assign a value to the first element
	theArray[1] = "Canada" // Assign a value to the second element
	theArray[2] = "Japan"  // Assign a value to the third element

	fmt.Println(theArray[0]) // Access the first element value
	fmt.Println(theArray[1]) // Access the second element valu
	fmt.Println(theArray[2]) // Access the third element valu
}

func IntializingArrayWithLiteral() {
	x := [5]int{10, 20, 30, 40, 50}   // Intialized with values
	var y [5]int = [5]int{10, 20, 30} // Partial assignment

	fmt.Println(x)
	fmt.Println(y)
}

func IntializingArrayWithEllipses() {
	x := [...]int{10, 20, 30}

	fmt.Println(reflect.ValueOf(x).Kind())
	fmt.Println(len(x))
}

func IntializingArrayWithSpecificValues() {
	x := [5]int{1: 10, 3: 30}
	fmt.Println(x)
}

func LoopThroughIndexedArray() {
	intArray := [5]int{10, 20, 30, 40, 50}

	fmt.Println("---------------Example 1--------------------")
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	fmt.Println("---------------Example 2--------------------")
	for index, element := range intArray {
		fmt.Println(index, "=>", element)

	}

	fmt.Println("---------------Example 3--------------------")
	for _, value := range intArray {
		fmt.Println(value)
	}

	j := 0
	fmt.Println("---------------Example 4--------------------")
	for range intArray {
		fmt.Println(intArray[j])
		j++
	}
}

func CopyArray() {

	strArray1 := [3]string{"Japan", "Australia", "Germany"}
	strArray2 := strArray1  // data is passed by value
	strArray3 := &strArray1 // data is passed by refrence

	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)

	strArray1[0] = "Canada"

	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)
	fmt.Printf("*strArray3: %v\n", *strArray3)
}

// Check if Element Exists

func ItemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func CheckIfItemExists() {
	strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(ItemExists(strArray, "Canada"))
	fmt.Println(ItemExists(strArray, "Africa"))
}

func FilterArrayElements() {
	countries := [...]string{"India", "Canada", "Japan", "Germany", "Italy"}
	length := len(countries)
	fmt.Printf("Countries: %v\n", countries)

	fmt.Printf(":2 %v\n", countries[:2])

	fmt.Printf("1:3 %v\n", countries[1:3])

	fmt.Printf("2: %v\n", countries[2:])

	fmt.Printf("2:5 %v\n", countries[2:5])

	fmt.Printf("0:3 %v\n", countries[0:3])

	fmt.Printf("Last element: %v\n", countries[length-1])

	fmt.Printf("All elements: %v\n", countries[0:length])
	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:length])

	fmt.Printf("Last two elements: %v\n", countries[length-2:length])
}

func Arrays() {
	SimpleArrayDeclaration()
	AssignAndAccessIntoArray()
	IntializingArrayWithLiteral()
	IntializingArrayWithEllipses()
	IntializingArrayWithSpecificValues()
	LoopThroughIndexedArray()
	CopyArray()
	CheckIfItemExists()
	FilterArrayElements()
}
