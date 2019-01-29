package sudoku

import (
	"testing"
)

func TestContains(t *testing.T) {
	arr := []int{2, 3, 5, 7, 11, 13}

	if !Contains(arr, 7) {
		t.Errorf("arr should contain 7")
	}

	if Contains(arr, 1) {
		t.Errorf("arr should NOT contain 1")
	}
}
func TestRandomNumbers(t *testing.T) {
	start, count := 3, 7
	randomNumbers := RandomNumbers(start, count)

	for i := start; i < start+count-1; i++ {
		if !Contains(randomNumbers, i) {
			t.Errorf("arr should contain %d", i)
		}
	}

}
func TestFirstMissingNumber(t *testing.T) {

	tests := []struct {
		numbers     []int
		expectedRes int
	}{
		{[]int{1, 2, 0, 4, 5}, 3},
		{[]int{5, 0, 3, 0, 1}, 2},
		{[]int{9, 7, 5, 4, 3, 2, 1, 0, 0}, 6},
	}

	for _, test := range tests {
		n := FirstMissingNumber(test.numbers)
		if n != test.expectedRes {
			t.Errorf("expected %d, but found %d", test.expectedRes, n)
		}
	}
}
