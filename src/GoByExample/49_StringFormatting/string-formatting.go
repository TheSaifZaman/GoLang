package stringFormatting

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {

    p := point{1, 2}

	var pc = fmt.Printf

    pc("struct1: %v\n", p)

    pc("struct2: %+v\n", p)

    pc("struct3: %#v\n", p)

    pc("type: %T\n", p)

    pc("bool: %t\n", true)

    pc("int: %d\n", 123)

    pc("bin: %b\n", 14)

    pc("char: %c\n", 33)

    pc("hex: %x\n", 456)

    pc("float1: %f\n", 78.9)

    pc("float2: %e\n", 123400000.0)

    pc("float3: %E\n", 123400000.0)

    pc("str1: %s\n", "\"string\"")

    pc("str2: %q\n", "\"string\"")

    pc("str3: %x\n", "hex this")

    pc("pointer: %p\n", &p)

    pc("width1: |%6d|%6d|\n", 12, 345)

    pc("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

    pc("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

    pc("width4: |%6s|%6s|\n", "foo", "b")

    pc("width5: |%-6s|%-6s|\n", "foo", "b")

    s := fmt.Sprintf("sprintf: a %s", "string")
    fmt.Println(s)

    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}

/*
$ go run string-formatting.go
struct1: {1 2}
struct2: {x:1 y:2}
struct3: main.point{x:1, y:2}
type: main.point
bool: true
int: 123
bin: 1110
char: !
hex: 1c8
float1: 78.900000
float2: 1.234000e+08
float3: 1.234000E+08
str1: "string"
str2: "\"string\""
str3: 6865782074686973
pointer: 0xc0000ba000
width1: |    12|   345|
width2: |  1.20|  3.45|
width3: |1.20  |3.45  |
width4: |   foo|     b|
width5: |foo   |b     |
sprintf: a string
io: an error
*/