package interfaces

/*
Declaring Interface Types
An Interface is an abstract type.

Interface describes all the methods of a method set and provides the signatures for each method.

To create interface use interface keyword, followed by curly braces containing a list of method names, along with any parameters or return values the methods are expected to have.

Example
// Declare an Interface Type and methods does not have a body
type Employee interface {
	PrintName() string                // Method with string return type
	PrintAddress(id int)              // Method with int parameter
	PrintSalary(b int, t int) float64 // Method with parameters and return type
}
*/

