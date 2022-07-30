package interfaces

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width*r.height
}

func (r rect) perim() float64 {
	return 2*(r.width + r.height)
}

func (c circle) area() float64 {
	return 2*math.Pi*c.radius*c.radius
}

func (c circle) perim() float64 {
	return 2*math.Pi*c.radius
}

//If a variable has an interface type, then we can call methods that are in the named interface. Here’s a generic measure function taking advantage of this to work on any geometry.
func mesaure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Interfaces() {
	r := rect{width: 5, height: 3}
	c := circle{radius: 5}

	mesaure(r)
	mesaure(c)
}

