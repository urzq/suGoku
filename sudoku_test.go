package main

import (
	"strings"
	"testing"
)

var grid_str_1_1 = `
. | . 
-----
. | .`

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
		{grid_str_1_1, 2, 1, 1},
		{grid_str_2_2, 4, 2, 2},
		{grid_str_2_3, 6, 2, 3},
		{grid_str_3_2, 6, 3, 2},
		{grid_str_3_3, 9, 3, 3},
		{grid_str_4_4, 16, 4, 4},
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

func TestToString(t *testing.T) {
	gridsStr := []string{
		grid_str_2_2,
		grid_str_2_3,
		grid_str_3_2,
		grid_str_3_3,
		grid_str_4_4,
	}

	for _, gridTestStr := range gridsStr {
		grid := createGridFromString(gridTestStr)
		gridStr := grid.String()

		gridStrTrim := strings.TrimSpace(gridStr)
		gridTestStrTrim := strings.TrimSpace(gridTestStr)

		if gridStrTrim != gridTestStrTrim {
			t.Errorf("Failed to convert grid to string.\n%s\n instead of\n %s",
				strings.TrimSpace(gridStr),
				strings.TrimSpace(gridTestStr))
		}
	}
}
