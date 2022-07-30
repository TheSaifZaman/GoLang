package intToSpecificInt

import (
	"fmt"
	"reflect"
)

func IntToSpecificInt() {
	// Convert Int data type to Int16 Int32 Int64

	var i int = 10
	fmt.Println(reflect.TypeOf(i))

	i16 := int16(i)
	fmt.Println(reflect.TypeOf(i16))

	i32 := int32(i)
	fmt.Println(reflect.TypeOf(i32))

	i64 := int64(i)
	fmt.Println(reflect.TypeOf(i64))
}