package main

import (
	"fmt"
	"testing"
)

var unsolvedGridStr = `
1 . . | . . . | 7 . .
. . . | . . . | . . .
. . . | . . . | . . .
---------------------
. . . | . . . | . . .
. . 1 | . . . | . . .
. . . | . . . | . . .
---------------------
. 1 . | . 2 . | 3 . .
. . . | . . . | . . .
. . . | . . . | . . .`

// var unsolvedGridStr = `
// 1 . | 4 .
// . . | . 3
// ---------
// . . | . .
// . 1 | . .`

func TestSolveGrid(t *testing.T) {
	unsolvedGrid := createGridFromString(unsolvedGridStr)
	success, solvedGrid := SolveGrid(unsolvedGrid)

	for row := 0; row < solvedGrid.Size; row++ {
		for col := 0; col < solvedGrid.Size; col++ {
			if !IsCellUniqueInCol(&solvedGrid, col, row) {
				t.Errorf("Cell at [%d][%d] is not unique within its column", col, row)
			}
			if !IsCellUniqueInRow(&solvedGrid, col, row) {
				t.Errorf("Cell at [%d][%d] is not unique within its row", col, row)
			}
			if !IsCellUniqueInRegion(&solvedGrid, col, row) {
				t.Errorf("Cell at [%d][%d] is not unique within its region,", col, row)
			}
		}
	}

	if !success {
		t.Errorf("Success was expected, but SolveGrid failed")
	} else {
		fmt.Printf(solvedGrid.String())
	}
}
