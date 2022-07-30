package maps

import ("fmt"
"sort"
)
	
/*
A map is a data structure that provides you with an unordered collection of key/value pairs (maps are also sometimes called associative arrays in Php, hash tables in Java, or dictionaries in Python). 
Maps are used to look up a value by its associated key. 
You store values into the map based on a key.

The strength of a map is its ability to retrieve data quickly based on the key. 
A key works like an index, pointing to the value you associate with that key.

A map is implemented using a hash table, which is providing faster lookups on 
the data element and you can easily retrieve a value by providing the key. 
Maps are unordered collections, and there's no way to predict the order in which the key/value pairs will be returned. 
Every iteration over a map could return a different order.
*/

 
func MapInitialization() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println(employee)
}

func EmptyMapDeclaration() {
	var employee = map[string]int{}
	fmt.Println(employee)        // map[]
	fmt.Printf("%T\n", employee) // map[string]int
}

func MapDeclarationUsingMake() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	fmt.Println(employee)
 
	employeeList := make(map[string]int)
	employeeList["Mark"] = 10
	employeeList["Sandy"] = 20
	fmt.Println(employeeList)
}

func MapLength() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
 
	// Empty Map
	employeeList := make(map[string]int)
 
	fmt.Println(len(employee))     // 2
	fmt.Println(len(employeeList)) // 0
}

func AccessingItems() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
 
	fmt.Println(employee["Mark"])
}

func AddingItems() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println(employee) // Initial Map
 
	employee["Rocky"] = 30 // Add element
	employee["Josef"] = 40
 
	fmt.Println(employee)
}

func UpdateValues() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println(employee) // Initial Map
 
	employee["Mark"] = 50 // Edit item
	fmt.Println(employee)
}

func DeleteItems() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	employee["Rocky"] = 30
	employee["Josef"] = 40
 
	fmt.Println(employee)
 
	delete(employee, "Mark")
	fmt.Println(employee)
}

func IterateOverMap() {
    var employee = map[string]int{"Mark": 10, "Sandy": 20,
        "Rocky": 30, "Rajiv": 40, "Kate": 50}
    for key, element := range employee {
        fmt.Println("Key:", key, "=>", "Element:", element)
    }
}

func TruncateMap() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20,
		"Rocky": 30, "Rajiv": 40, "Kate": 50}
 
	// Method - I
	for k := range employee {
		delete(employee, k)
	}
 
	// Method - II
	employee = make(map[string]int)
}

func SortMapKeys() {
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}
 
	keys := make([]string, 0, len(unSortedMap))
 
	for k := range unSortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
 
	for _, k := range keys {
		fmt.Println(k, unSortedMap[k])
	}
}

func SortMapValues() {
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}
 
 // Int slice to store values of map.
	values := make([]int, 0, len(unSortedMap))
 
	for _, v := range unSortedMap {
		values = append(values, v)
	}
 
 // Sort slice values.
	sort.Ints(values)
 
 // Print values of sorted Slice.
	for _, v := range values {
		fmt.Println(v)
	}
}

func MergeMaps() {
	first := map[string]int{"a": 1, "b": 2, "c": 3}
	second := map[string]int{"a": 1, "e": 5, "c": 3, "d": 4}
 
	for k, v := range second {
		first[k] = v
	}
 
	fmt.Println(first)
}

func Slices() {
	MapInitialization()
	EmptyMapDeclaration()
	MapDeclarationUsingMake()
	MapLength()
	AccessingItems()
	AddingItems()
	UpdateValues()
	DeleteItems()
	IterateOverMap()
	TruncateMap()
	SortMapKeys()
	SortMapValues()
	MergeMaps()
}