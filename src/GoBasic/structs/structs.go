package structs

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func DeclarationOfStructType() {
	type rectangle struct {
		length  float64
		breadth float64
		color   string
	}

	fmt.Println(rectangle{10.5, 25.10, "red"})
}

func CreatingInstancesOfStructTypes() {
	type rectangle struct {
		length  int
		breadth int
		color   string

		geometry struct {
			area      int
			perimeter int
		}
	}

	var rect rectangle // dot notation
	rect.length = 10
	rect.breadth = 20
	rect.color = "Green"

	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.length + rect.breadth)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area)
	fmt.Println("Perimeter:", rect.geometry.perimeter)
}

func CreatingStructInstanceUsingStructLiteral() {
	type rectangle struct {
		length  int
		breadth int
		color   string
	}

	var rect1 = rectangle{10, 20, "Green"}
	fmt.Println(rect1)

	var rect2 = rectangle{length: 10, color: "Green"} // breadth value skipped
	fmt.Println(rect2)

	rect3 := rectangle{10, 20, "Green"}
	fmt.Println(rect3)

	rect4 := rectangle{length: 10, breadth: 20, color: "Green"}
	fmt.Println(rect4)

	rect5 := rectangle{breadth: 20, color: "Green"} // length value skipped
	fmt.Println(rect5)
}

func StructInstantiationUsingNewKeyword() {

	type rectangle struct {
		length  int
		breadth int
		color   string
	}

	rect1 := new(rectangle) // rect1 is a pointer to an instance of rectangle
	rect1.length = 10
	rect1.breadth = 20
	rect1.color = "Green"
	fmt.Println(rect1)

	var rect2 = new(rectangle) // rect2 is an instance of rectangle
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2)
}

func StructInstantiationUsingPointerAddressOperator() {

	type rectangle struct {
		length  int
		breadth int
		color   string
	}
	var rect1 = &rectangle{10, 20, "Green"} // Can't skip any value
	fmt.Println(rect1)

	var rect2 = &rectangle{}
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2) // breadth skipped

	var rect3 = &rectangle{}
	(*rect3).breadth = 10
	(*rect3).color = "Blue"
	fmt.Println(rect3) // length skipped
}

func NestedStructType() {
	type Salary struct {
		Basic, HRA, TA float64
	}

	type Employee struct {
		FirstName, LastName, Email string
		Age                        int
		MonthlySalary              []Salary
	}

	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	fmt.Println(e.MonthlySalary[0])
	fmt.Println(e.MonthlySalary[1])
	fmt.Println(e.MonthlySalary[2])
}

func UseFieldTagsInTheDefinitionOfStructType() {

	type Employee struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		City      string `json:"city"`
	}
	json_string := `
    {
        "firstname": "Rocky",
        "lastname": "Sting",
        "city": "London"
    }`

	emp1 := new(Employee)
	json.Unmarshal([]byte(json_string), emp1)
	fmt.Println(emp1)

	emp2 := new(Employee)
	emp2.FirstName = "Ramesh"
	emp2.LastName = "Soni"
	emp2.City = "Mumbai"
	jsonStr, _ := json.Marshal(emp2)
	fmt.Printf("%s\n", jsonStr)
}

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func (e Employee) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}
	return "----------------------"
}

func AddMethodToStructType() {

	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	fmt.Println(e.EmpInfo())
}

type Employee_1 struct {
	Name string
	Age  int
}

func (obj *Employee_1) Info() {
	if obj.Name == "" {
		obj.Name = "John Doe"
	}
	if obj.Age == 0 {
		obj.Age = 25
	}
}

func AssignDefaultValueForStructField() {
	emp1 := Employee_1{Name: "Mr. Fred"}
	emp1.Info()
	fmt.Println(emp1)

	emp2 := Employee_1{Age: 26}
	emp2.Info()
	fmt.Println(emp2)
}

func FindTypeOfStruct() {

	type rectangle struct {
		length  float64
		breadth float64
		color   string
	}

	var rect1 = rectangle{10, 20, "Green"}
	fmt.Println(reflect.TypeOf(rect1))         // main.rectangle
	fmt.Println(reflect.ValueOf(rect1).Kind()) // struct

	rect2 := rectangle{length: 10, breadth: 20, color: "Green"}
	fmt.Println(reflect.TypeOf(rect2))         // main.rectangle
	fmt.Println(reflect.ValueOf(rect2).Kind()) // struct

	rect3 := new(rectangle)
	fmt.Println(reflect.TypeOf(rect3))         // *main.rectangle
	fmt.Println(reflect.ValueOf(rect3).Kind()) // ptr

	var rect4 = &rectangle{}
	fmt.Println(reflect.TypeOf(rect4))         // *main.rectangle
	fmt.Println(reflect.ValueOf(rect4).Kind()) // ptr
}

// Comparing Structs with the Different Values Assigned to Data Fields

func ComparingSameTypeStructs() {

	type rectangle struct {
		length  float64
		breadth float64
		color   string
	}

	var rect1 = rectangle{10, 20, "Green"}
	rect2 := rectangle{length: 20, breadth: 10, color: "Red"}

	if rect1 == rect2 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	rect3 := new(rectangle)
	var rect4 = &rectangle{}

	if rect3 == rect4 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

// Copy Struct Type Using Value and Pointer Reference

func CopyStruct() {

	type rectangle struct {
		length  float64
		breadth float64
		color   string
	}

	r1 := rectangle{10, 20, "Green"}
	fmt.Println(r1)

	r2 := r1
	r2.color = "Pink"
	fmt.Println(r2)

	r3 := &r1
	r3.color = "Red"
	fmt.Println(r3)

	fmt.Println(r1)
}
