package sum

import "testing"

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 3, 5, 8, 9, 19, 51, 135})
	}
}
func BenchmarkSumRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumRecursion([]int{1, 3, 5, 8, 9, 19, 51, 135})
	}
}
func BenchmarkSumWithAccumulator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumWithAccumulator([]int{1, 3, 5, 8, 9, 19, 51, 135}, 0)
	}
}
