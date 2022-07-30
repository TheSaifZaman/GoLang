package interfaceEmbeddimg

// Interfaces may embed other interfaces, this behavior is an aspect of interface polymorphism which is known as ad hoc polymorphism.

import "fmt"

type Geometry interface {
	Edges() int
}

type Polygons interface {
	Geometry // Interface embedding another interface
}

type Pentagon int
type Hexagon int
type Octagon int
type Decagon int

func (p Pentagon) Edges() int { return 5 }
func (h Hexagon) Edges() int  { return 6 }
func (o Octagon) Edges() int  { return 8 }
func (d Decagon) Edges() int  { return 10 }

func InterfaceEmbeddimg() {
	p := new(Pentagon)
	h := new(Hexagon)
	o := new(Octagon)
	d := new(Decagon)

	polygons := [...]Polygons{p, h, o, d}
	for i := range polygons {
		fmt.Println(polygons[i].Edges())
	}
}