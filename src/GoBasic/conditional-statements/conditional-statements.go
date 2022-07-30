package conditionalStatements

import (
	"fmt"
	"time"
)

/*
Golang Conditional Statements
Like most programming languages, Golang borrows several of its control flow syntax from the C-family of languages. In Golang we have the following conditional statements:


The if statement - executes some code if one condition is true
The if...else statement - executes some code if a condition is true and another code if that condition is false
The if...else if....else statement - executes different codes for more than two conditions
The switch...case statement - selects one of many blocks of code to be executed
We will explore each of these statements in the coming sections.
*/

func IfStatement() {
	var s = "Japan"
	x := true
	if x {
		fmt.Println(s)
	}
}

func IfElseStatement() {
	x := 100
 
	if x == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}
}

func IfElseIfElseStatement() {
	x := 100
 
	if x == 50 {
		fmt.Println("Germany")
	} else if x == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}
}

func IfStatementInitialization() {
	if x := 100; x == 100 {
		fmt.Println("Germany")
	}
}

func SwitchStatement() {
	today := time.Now()
 
	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

func SwitchMultipleCasesStatement() {
	today := time.Now()
	var t int = today.Day()
 
	switch t {
	case 5, 10, 15:
		fmt.Println("Clean your house.")
	case 25, 26, 27:
		fmt.Println("Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

func SwitchFallthroughCaseStatement() {
	today := time.Now()
 
	switch today.Day() {
	case 5:
		fmt.Println("Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 25:
		fmt.Println("Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

func SwitchConditionalCasesStatement() {
	today := time.Now()
 
	switch {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}

func SwitchInitializerStatement() {
	switch today := time.Now(); {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}