package main

import "testing"

var gridStr33 = `
1 2 3 | 4 5 6 | 7 8 9
4 5 6 | 4 5 6 | . 5 6
7 8 9 | 7 8 9 | 7 . 9
---------------------
4 5 6 | 1 2 3 | 7 8 9
4 2 6 | 1 2 3 | 7 8 9
8 1 9 | 7 8 9 | 7 1 9
---------------------
3 4 5 | 6 2 7 | 8 9 2
6 . 6 | . 5 . | . . 6
5 8 9 | 7 8 9 | 7 8 9`

func TestIsCellUniqueInRow(t *testing.T) {
	tests := []struct {
		col    int
		row    int
		unique bool
	}{
		{0, 0, true},
		{4, 3, true},
		{4, 4, false},
		{4, 5, false},
		{4, 6, false},
		{4, 6, false},
		{4, 7, true},
	}

	grid := createGridFromString(gridStr33)

	for _, test := range tests {
		unique := IsCellUniqueInRow(&grid, test.col, test.row)
		if unique != test.unique {
			t.Errorf("Cell at [%d][%d] should be unique (=%t), but is computed as unique(=%t)",
				test.col, test.row, test.unique, unique)
		}
	}
}

var gridStr23 = `
1 1 | 2 5 | 2 3
4 3 | 5 6 | . .
6 4 | 8 9 | . 9
---------------
5 2 | 1 2 | 2 3
3 5 | 8 5 | 5 6
2 6 | 7 8 | 5 9`

func TestIsCellUniqueInCol(t *testing.T) {
	tests := []struct {
		col    int
		row    int
		unique bool
	}{
		{0, 1, true},
		{1, 3, true},
		{2, 2, false},
		{3, 4, false},
		{4, 4, false},
		{5, 4, true},
	}

	grid := createGridFromString(gridStr23)

	for _, test := range tests {
		unique := IsCellUniqueInCol(&grid, test.col, test.row)

		if unique != test.unique {
			t.Errorf("Cell at [%d][%d] should be unique (=%t), but is computed as unique(=%t)",
				test.col, test.row, test.unique, unique)
		}
	}
}

func TestIsCellUniqueInRegion(t *testing.T) {
	tests := []struct {
		col    int
		row    int
		unique bool
	}{
		{2, 0, true},
		{2, 1, false},
		{0, 3, false},
		{2, 3, true},
		{5, 5, true},
	}
	grid := createGridFromString(gridStr23)

	for _, test := range tests {
		unique := IsCellUniqueInRegion(&grid, test.col, test.row)

		if unique != test.unique {
			t.Errorf("Cell at [%d][%d] should be unique (=%t), but is computed as unique(=%t)",
				test.col, test.row, test.unique, unique)
		}
	}
}
