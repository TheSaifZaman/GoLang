package jsonFile
 
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)
 
 
/*
Reading JSON file
The json package includes Unmarshal() function which supports decoding data from a byte slice into values. The decoded values are generally assigned to struct fields, the field names must be exported and should be in capitalize format.

The JSON file test.json is read with the ioutil.ReadFile() function, which returns a byte slice that is decoded into the struct instance using the json.Unmarshal() function. At last, the struct instance member values are printed using for loop to demonstrate that the JSON file was decoded.
*/
func ReadJson() {
	type Catlog struct {
		Product_id string `json: "product_id"`
		Quantity   int    `json: "quantity"`
	}
	
	type CatlogNodes struct {
		CatlogNodes []Catlog `json:"catlog_nodes"`
	}

	file, _ := ioutil.ReadFile("test.json")
 
	data := CatlogNodes{}
 
	_ = json.Unmarshal([]byte(file), &data)
 
	for i := 0; i < len(data.CatlogNodes); i++ {
		fmt.Println("Product Id: ", data.CatlogNodes[i].Product_id)
		fmt.Println("Quantity: ", data.CatlogNodes[i].Quantity)
	}
 
}

/*
Writing JSON file
The json package has a MarshalIndent() function which is used to serialized values from a struct and write them to a file in JSON format.

The Salary struct is defined with json fields. The struct values are initialized and then serialize with the json.MarshalIndent() function. The serialized JSON formatted byte slice is received which then written to a file using the ioutil.WriteFile() function.
*/
func WriteJson() {

	type Salary struct {
		Basic, HRA, TA float64
	}
	
	type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
	}

	
data := Employee{
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

file, _ := json.MarshalIndent(data, "", " ")

_ = ioutil.WriteFile("test.json", file, 0644)
}