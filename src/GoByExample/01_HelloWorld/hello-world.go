package helloWorld

import (
	"fmt"
	"runtime"
)

func HelloWorld() {
	fmt.Println("Hello World, This is my first go program")
	fmt.Printf("Hello, version of t his go is %s\n", runtime.Version())
}