package sudoku

// IsCellUniqueInRow returns true if the value of the cell at [col, row] is
// unique within its row.
func IsCellUniqueInRow(grid *Grid, col int, row int) bool {
	index := row*grid.Size + col
	cell := grid.Cells[index]

	start := row * grid.Size
	end := start + grid.Size

	for i := start; i < end; i++ {
		if i != index && grid.Cells[i] == cell {
			return false
		}
	}

	return true
}

// IsCellUniqueInCol returns true if the value of the cell at [col, row] is
// unique within its col.
func IsCellUniqueInCol(grid *Grid, col int, row int) bool {
	index := row*grid.Size + col
	cell := grid.Cells[index]

	start := col
	end := col + grid.Size*(grid.Size-1)

	for i := start; i <= end; i += grid.Size {
		if i != index && grid.Cells[i] == cell {
			return false
		}
	}

	return true
}

// IsCellUniqueInRegion returns true if the value of the cell at [col, row] is
// unique within its region.
func IsCellUniqueInRegion(grid *Grid, col int, row int) bool {
	index := row*grid.Size + col
	cell := grid.Cells[index]

	startColumn := col - (col % grid.RegionWidth)
	startRow := row - (row % grid.RegionHeight)

	for r := startRow; r < startRow+grid.RegionHeight; r++ {
		for c := startColumn; c < startColumn+grid.RegionWidth; c++ {
			i := r*grid.Size + c

			if i != index && grid.Cells[i] == cell {
				return false
			}
		}
	}

	return true
}
