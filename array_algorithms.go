package sudoku

import (
	"math/rand"
	"time"
)

var s = rand.NewSource(time.Now().UnixNano())
var r = rand.New(s)

// Contains tells whether arr contains elem.
func Contains(arr []int, elem int) bool {
	for _, n := range arr {
		if elem == n {
			return true
		}
	}
	return false
}

// RandomNumbers returns an array of size "count", containing  a random
// permutation of all numbers from start to (start + count - 1).
func RandomNumbers(start int, count int) []int {
	arr := make([]int, count)

	for i := 0; i < count; i++ {
		arr[i] = i + start
	}

	for i := 0; i < count; i++ {
		j := r.Intn(count)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

// FirstMissingNumber loops through arr to find the first missing number in
// the sequence [1..len(arr)-1].
func FirstMissingNumber(arr []int) int {
	// start presentNumbers with a 0, so we can easily test if a number n
	// is present in arr by looking at the value of presentNumber[x]
	presentNumbers := make([]bool, len(arr)+1)
	presentNumbers[0] = true

	for _, n := range arr {
		if n > 0 {
			presentNumbers[n] = true
		}
	}

	for i, n := range presentNumbers {
		if !n {
			return i
		}
	}

	return 0
}
