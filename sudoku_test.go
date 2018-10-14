package sudoku

import (
	"testing"
)

var grid_str_2_2 = `
	2 3 | 1 2 
	5 6 | 4 5 
	---------
	1 2 | . 2 
	. 8 | 7 8`

var grid_str_2_3 = `
	1 2 | 2 3 | 2 3
	4 5 | 5 6 | 5 6
	7 8 | 8 9 | . 9
	---------------
	1 2 | 1 2 | 2 3
	4 5 | 4 5 | 5 6 
	7 8 | 7 8 | 7 9`

var grid_str_3_2 = `
	1 2 3 | 1 2 3
	4 5 6 | 4 5 6
	-------------
	1 2 3 | 1 2 3
	7 8 9 | 7 8 9
	-------------
	1 2 3 | 1 2 3 
	7 8 9 | 7 8 9`

var grid_str_3_3 = `
	1 2 3 | 1 2 3 | 1 2 3
	4 5 6 | 4 5 6 | . 5 6
	7 8 9 | 7 8 9 | 7 . 9
	---------------------
	1 2 3 | 1 2 3 | 1 2 3
	4 5 6 | 4 5 6 | 4 5 6
	7 8 9 | 7 8 9 | 7 8 9
	---------------------
	1 2 3 | 1 2 3 | 1 2 3
	4 5 6 | 4 5 6 | 4 5 6
	7 8 9 | 7 8 9 | 7 8 9`

var grid_str_4_4 = `
	 1  2  3 10 |  1  2  3 10 |  1  2  3 10 |  1  2  3 10
	 4  5  6 11 |  4  5  6 11 |  4  5  6 11 |  4  5  6 11
	 7  8  9 12 |  7  8  9 12 |  7  8  9 12 |  7  8  9 12
	13 14 15 16 | 13 14 15 16 | 13 14 15 16 | 13 14 15 16
	-----------------------------------------------------
	 1  2  3 10 |  1  2  3 10 |  1  2  3 10 |  1  2  3 10
	 4  5  6 11 |  4  5  6 11 |  4  5  6 11 |  4  5  6 11
	 7  8  9 12 |  7  8  9 12 |  7  8  9 12 |  7  8  9 12
	13 14 15 16 | 13 14 15 16 | 13 14 15 16 | 13 14 15 16
	-----------------------------------------------------
	 1  2  3 10 |  1  2  3 10 |  1  2  3 10 |  1  2  3 10
	 4  5  6 11 |  4  5  6 11 |  4  5  6 11 |  4  5  6 11
	 7  8  9 12 |  7  8  9 12 |  7  8  9 12 |  7  8  9 12
	13 14 15 16 | 13 14 15 16 | 13 14 15 16 | 13 14 15 16
	-----------------------------------------------------
	 1  2  3 10 |  1  2  3 10 |  1  2  3 10 |  1  2  3 10
	 4  5  6 11 |  4  5  6 11 |  4  5  6 11 |  4  5  6 11
	 7  8  9 12 |  7  8  9 12 |  7  8  9 12 |  7  8  9 12
	13 14 15 16 | 13 14 15 16 | 13 14 15 16 | 13 14 15 16`

func TestCreateGrid(t *testing.T) {
	tests := []struct {
		gridStr              string
		expectedSize         int
		expectedRegionWidth  int
		expectedRegionHeight int
	}{
		{gridStr: grid_str_2_2, expectedSize: 4, expectedRegionWidth: 2, expectedRegionHeight: 2},
		{gridStr: grid_str_2_3, expectedSize: 6, expectedRegionWidth: 2, expectedRegionHeight: 3},
		{gridStr: grid_str_3_2, expectedSize: 6, expectedRegionWidth: 3, expectedRegionHeight: 2},
		{gridStr: grid_str_3_3, expectedSize: 9, expectedRegionWidth: 3, expectedRegionHeight: 3},
		{gridStr: grid_str_4_4, expectedSize: 16, expectedRegionWidth: 4, expectedRegionHeight: 4},
	}

	for _, test := range tests {
		grid := createGridFromString(test.gridStr)

		if grid.Size != test.expectedSize {
			t.Errorf("Incorrect grid size. (%d instead of %d)", grid.Size, test.expectedSize)
		}

		if grid.RegionWidth != test.expectedRegionWidth {
			t.Errorf("Incorrect RegionWidth (%d instead of %d)", grid.RegionWidth, test.expectedRegionWidth)
		}

		if grid.RegionHeight != test.expectedRegionHeight {
			t.Errorf("Incorrect RegionHeight (%d instead of %d)", grid.RegionHeight, test.expectedRegionHeight)
		}
	}
}
