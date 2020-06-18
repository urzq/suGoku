package main

import "fmt"

// SolveGrid takes as an input an unsolved sudoku grid, and it tries to solve it.
func SolveGrid(grid Grid) (success bool, result Grid) {

	if grid.Cells[0] == 0 {
		candidates := RandomNumbers(1, grid.RegionWidth*grid.RegionHeight)
		for _, candidate := range candidates {
			success, result = solveGridHelper(grid, 0, candidate)
			if success {
				return
			}
		}
	} else {
		success, result = solveGridHelper(grid, 0, grid.Cells[0])
		if success {
			return
		}
	}

	success = false
	return
}

// TODO: expliquer que cell veut dire "je remplis grid[index] = cell"
func solveGridHelper(grid Grid, index int, cell int) (success bool, result Grid) {

	grid.Cells[index] = cell

	col := index % grid.Size
	row := index / grid.Size

	fmt.Printf("i:%d col:%d row:%d cells:\n%s\n\n", index, col, row, grid.String())

	// if the current cell is not valid, terminate this recursion.
	if !IsCellUniqueInCol(&grid, col, row) ||
		!IsCellUniqueInRow(&grid, col, row) ||
		!IsCellUniqueInRegion(&grid, col, row) {
		success = false
		return
	}

	if index == len(grid.Cells)-1 {
		// This is the last cell of the grid, it all the previous cells (plus the current)
		// respect the unicity constraints. We succeed :)
		success = true
		result = grid
		return
	} else if grid.Cells[index+1] == 0 {
		// We need to find the next cell value. Just test all the possibilities.
		candidates := RandomNumbers(1, grid.RegionWidth*grid.RegionHeight)
		for _, candidate := range candidates {

			nextGrid := grid
			nextGrid.Cells = make([]int, len(grid.Cells))
			copy(nextGrid.Cells, grid.Cells)

			success, result = solveGridHelper(nextGrid, index+1, candidate)
			if success {
				return
			}
		}
	} else {
		// The next cell value is known. Advance.
		success, result = solveGridHelper(grid, index+1, grid.Cells[index+1])
		if success {
			return
		}
	}

	return
}
