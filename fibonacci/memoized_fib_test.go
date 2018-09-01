package fibonacci

import "testing"

var testcases = []struct {
	input    int
	expected int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{21, 10946},
	{43, 433494437},
}

func TestMyFibMemoized(t *testing.T) {
	for _, testcase := range testcases {
		if v := MyFibMemoized(testcase.input); v != testcase.expected {
			t.Errorf("input value %d should yield %d, but got %d", testcase.input, testcase.expected, v)
		}
	}
}

func BenchmarkMyFibMemoized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MyFibMemoized(8)
	}
}
