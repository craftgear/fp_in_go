package main

import (
	"fmt"

	"craftgear.net/fp_in_go/fibonacci"
	"craftgear.net/fp_in_go/sum"
)

func main() {
	fib50 := fibonacci.MyFibMemoized(50)
	fmt.Printf("fib50 = %+v\n", fib50)
	sum10 := sum.Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Printf("sum10 = %+v\n", sum10)
	sum10WithAccum := sum.SumWithAccumulator([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0)
	fmt.Printf("sum10 = %+v\n", sum10WithAccum)
}
