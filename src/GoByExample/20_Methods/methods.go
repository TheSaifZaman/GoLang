package methods

import "fmt"

type rect struct {
	width, height float32
}

func (r *rect) area() float32 {
	return r.width*r.height
}

func (r rect) perim() float32 {
	return 2*(r.width + r.height)
}

func Methods() {
	r := rect{width:111.11, height:22.22}
	fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())
}