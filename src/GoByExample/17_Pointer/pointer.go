package pointer

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func Pointer() {
	i := 1
	fmt.Println("initial:", i)
	fmt.Println("initial-pointer:", &i)
	
	zeroval(i)
	fmt.Println("zeroval:", i)
	fmt.Println("zeroval-pointer:", &i)

	zeroptr(&i)
	fmt.Println("zeroptr-value:", i)
	fmt.Println("zeroptr-pointer:", &i)
}