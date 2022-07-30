package structs

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string, age int) * person {
	p := person{name: name, age: age}
	return &p
}

func structs() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name:"Bob", age:20})
	fmt.Println(person{name:"Bob"})
	fmt.Println(&person{name:"Bob", age:30})
	fmt.Println(newPerson("Bob",30))

	s := person{name:"Bob", age:210}
	fmt.Println(s.age)

	sp := &s
	sp.age = 12
	fmt.Println(sp.age, s.age)
}