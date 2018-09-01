package main

import (
	"fmt"

	"craftgear.net/fp_in_go/fibonacci"
)

func main() {
	fib50 := fibonacci.MyFibMemoized(50)
	fmt.Printf("fib50 = %+v\n", fib50)
}
