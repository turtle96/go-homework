package sparse_number

import (
	"testing"
)

// todo add more test cases
var expected = map[int]int{
	21:      21,
	22:      32,
	15847:   16384,
	11:      16,
	363:     512,
	1923:    2048,
	3938361: 4194304,
}

func TestGetSparseNumber_BruteForce(t *testing.T) {
	for input, output := range expected {
		bruteForce := bruteForce(input)

		if output != bruteForce {
			t.Errorf("Input: %d, Expected Output: %d\n, Brute Force Output: %d", input, output, bruteForce)
		}
	}
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
	for input := range expected {
		for i := 0; i < b.N; i++ {
			GetSparseNumber(input)
		}
	}
}

func BenchmarkBruteForce(b *testing.B) {
	for input := range expected {
		for i := 0; i < b.N; i++ {
			bruteForce(input)
		}
	}
}
