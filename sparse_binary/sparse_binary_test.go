package sparse_binary

import (
	"testing"
)

// todo add more test cases
var expected = map[int]int{
	21: 21,
	22: 32,
}

func TestGetSparseNumber(t *testing.T) {
	for input, output := range expected {
		answer := GetSparseNumber(input)

		if answer != output {
			t.Errorf("Input: %d, Expected Output: %d\n, Actual Output: %d", input, output, answer)
		}
	}
}

func BenchmarkGetSparseNumber(b *testing.B) {
	// todo write benchmark

	for i := 0; i < b.N; i++ {
		for input := range expected {
			GetSparseNumber(input)
		}
	}
}
