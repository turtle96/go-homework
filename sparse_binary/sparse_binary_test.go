package sparse_binary

import (
	"testing"
)

func TestGetSparseNumber(t *testing.T) {
	// todo add more test cases
	expected := map[int]int{
		21: 21,
		22: 32,
	}

	for input, output := range expected {
		answer := GetSparseNumber(input)

		if answer != output {
			t.Errorf("Input: %d, Expected Output: %d\n, Actual Output: %d", input, output, answer)
		}
	}
}
