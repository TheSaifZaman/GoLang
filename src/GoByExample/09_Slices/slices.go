package slices

import "fmt"

/*
In Go language slice is more powerful, flexible, convenient than an array, 
and is a lightweight data structure. Slice is a variable-length sequence which stores elements of a similar type, 
you are not allowed to store different type of elements in the same slice. It is just like an array having an index value and length, 
but the size of the slice is resized they are not in fixed-size just like an array. 
Internally, slice and an array are connected with each other, a slice is a reference to an underlying array. It is allowed to store duplicate elements in the slice. 
The first index position in a slice is always 0 and the last one will be (length of slice – 1).
*/

func SimpleSlice() {
    
    s := make([]string, 3)
    fmt.Println("emp:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)
}

func TwoDimensionalSlice() {

    twoDimension := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoDimension[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoDimension[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoDimension)
}

func Slices() {

    SimpleSlice()
    TwoDimensionalSlice()
}