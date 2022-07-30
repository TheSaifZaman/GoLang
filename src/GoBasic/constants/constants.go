package costants

import "fmt"

const PRODUCT string = "Canada"
const PRICE = 500

func ConstantDeclaration() {
	fmt.Println(PRODUCT)
	fmt.Println(PRICE)
}

const (
	QUANTITY = 50
	STOCK    = true
)

func ContantBlock() {
	fmt.Println(QUANTITY)
	fmt.Println(STOCK)
}
func Constants() {
	// Declaring (Creating) Constants
	/*
		The keyword const is used for declaring constants followed by the desired name
		and the type of value the constant will hold. You must assign a value at the
		time of the constant declaration, you can't assign a value later as with variables.
	*/
	ConstantDeclaration()

	// Multilple Constants Declaration Block
	// Constants declaration can to be grouped together into blocks for greater readability and code quality.
	ContantBlock()

	// Naming Conventions for Golang Constants
	/*
		Name of constants must follow the same rules as variable names,
			which means a valid constant name must starts with a letter or underscore, followed by any number of letters, numbers or underscores.
		By convention, constant names are usually written in uppercase letters.
			This is for their easy identification and differentiation from variables in the source code.
	*/
}
