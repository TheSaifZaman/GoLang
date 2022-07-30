package main

type triangle struct {
	height float64
	base   float64
}

func (t triangle) area() float64 {
	return t.base * t.height / 2
}

type shape interface {
	area() float64
}

type calculator struct {
}

func (a calculator) areaSum(shapes ...shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.area()
	}
	return sum
}
//func main() {
//	c := circle{radius: 5}
//	s := square{length: 7}
//	t := triangle{height: 7, base: 8}
//	calc := calculator{}
//	fmt.Println("Area sum:", calc.areaSum(c, s, t))
//}
