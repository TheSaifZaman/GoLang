package floatToSpecificFloat

import (
	"fmt"
	"reflect"
)

func FloatToSpecificFloat() {
	// Convert Float32 to Float64 and Float64 to Float32
	
	var f32 float32 = 10.6556
	fmt.Println(reflect.TypeOf(f32))

	f64 := float64(f32)
	fmt.Println(reflect.TypeOf(f64))

	f64 = 1097.655698798798
	fmt.Println(f64)

	f32 = float32(f64)
	fmt.Println(f32)
}