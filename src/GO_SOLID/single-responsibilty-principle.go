package main

import (
	//"encoding/json"
	//"fmt"
	//"log"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) name() string {
	return "circle"
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius,2)
}

type square struct {
	length float64
}

func (s square) name() string {
	return "square"
}
func (s square) area() float64 {
	return math.Pow(s.length,2)
}

//type shape interface {
//	area() float64
//	name() string
//}
//type outputer struct {
//}
//
//func (o outputer) Text(s shape) string {
//	return fmt.Sprintf("Area of the %s: %f", s.name(), s.area())
//}
//func (o outputer) JSON(s shape) string {
//	res := struct {
//		Name string  `json:"shape"`
//		Area float64 `json:"area"`
//	}{
//		Name: s.name(),
//		Area: s.area(),
//	}
//	bs, err := json.Marshal(res)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return string(bs)
//}
//func main() {
//	c := circle{radius: 5}
//	s := square{length: 6}
//	out := outputter{}
//	fmt.Println(out.Text(c))
//	fmt.Println(out.JSON(c))
//	fmt.Println(out.Text(s))
//	fmt.Println(out.JSON(s))
//}
