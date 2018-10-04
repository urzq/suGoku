package sudoku

import (
	"testing"
)

func TestCreateGrid_2_2(t *testing.T) {
	src := `
		2 3 | 1 2 
		5 6 | 4 5 
		---------
		1 2 | . 2 
		. 8 | 7 8`

	grid := createGridFromString(src)

	if grid.Size != 4 {
		t.Errorf("Incorrect grid size. (%d instead of 4)", grid.Size)
	}

	if grid.RegionWidth != 2 {
		t.Errorf("Incorrect RegionWidth (%d instead of 2)", grid.RegionWidth)
	}

	if grid.RegionHeight != 2 {
		t.Errorf("Incorrect RegionHeight (%d instead of 2)", grid.RegionHeight)
	}
}

func TestCreateGrid_3_3(t *testing.T) {
	var src = `
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

	grid := createGridFromString(src)

	if grid.Size != 9 {
		t.Errorf("Incorrect grid size. (%d instead of 9)", grid.Size)
	}

	if grid.RegionWidth != 3 {
		t.Errorf("Incorrect RegionWidth (%d instead of 3)", grid.RegionWidth)
	}

	if grid.RegionHeight != 3 {
		t.Errorf("Incorrect RegionHeight (%d instead of 2)", grid.RegionHeight)
	}
}
