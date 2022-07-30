package maps

import "fmt"

// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).

func Maps() {

    m := make(map[string]int)

    m["key_1"] = 7
    m["key_2"] = 13

    fmt.Println("map:", m)

    value_1 := m["key_1"]
    fmt.Println("value_1: ", value_1)

    fmt.Println("len:", len(m))

    delete(m, "key_2")
    fmt.Println("map:", m)

	/*
	The optional second return value when getting a value from a map indicates if the key was present in the map. 
	This can be used to disambiguate between missing keys and keys with zero values like 0 or "". 
	Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	*/
    _, prs := m["key_2"]
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}