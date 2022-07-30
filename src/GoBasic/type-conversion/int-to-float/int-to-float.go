package intToFloat

import (
	"fmt"
	"reflect"
)

func IntToFloat() {
	/// Converting Int data type to Float in Go
	
	var f32 float32 = 10.6556
	fmt.Println(reflect.TypeOf(f32))

	i32 := int32(f32)
	fmt.Println(reflect.TypeOf(i32))
	fmt.Println(i32)

	f64 := float64(i32)
	fmt.Println(reflect.TypeOf(f64))
}