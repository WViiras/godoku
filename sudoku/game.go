package sudoku

import (
	"fmt"
)

// Game will hold the grid and present a solution
type Game struct {
	grid *Grid
}

// PopulateGrid fills Grid with known values
func (game *Game) PopulateGrid(unsolvedGrid []string) {
	game.grid = new(Grid)

	col := 0
	row := 0
	for _, v := range unsolvedGrid {
		for _, c := range v {
			intVal := 0
			if c != '-' {
				intVal = int(c - '0')
			}
			game.grid.SetCellValue(row, col, intVal)
			col++
			if col >= 9 {
				row++
				col = 0
			}
		}
	}
}

// Solve returns a solution
func (game *Game) Solve() Grid {
	fmt.Println("Found this sudoku")
	fmt.Println(game.grid)

	solvedGrid, solved := recursiveSolve(0, 0, *game.grid)
	if solved {
		fmt.Println("Found a solution")
		fmt.Println(solvedGrid)
	} else {
		fmt.Println("Unable to find a solution")
	}
	return *solvedGrid
}

func recursiveSolve(row, col int, grid Grid) (*Grid, bool) {
	// // if this cell already has a value, try next cell
	if grid.GetCellValue(row, col) != 0 {
		col++
		if col >= 9 {
			// go to next row
			col = 0
			row++
			if row >= 9 {
				// last column and row
				// if the last cell of the grid is set, it can be there
				// set solved and return
				return &grid, true
			}
		}
		return recursiveSolve(row, col, grid)
	}

	// recursively try all values
	for i := 1; i <= 9; i++ {
		if !grid.CanBeInBox(col, row, i) {
			continue
		}
		if !grid.CanBeInRow(col, row, i) {
			continue
		}
		if !grid.CanBeInColumn(col, row, i) {
			continue
		}
		// place value
		grid.SetCellValue(row, col, i)

		// try next cell
		if solvedGrid, solved := recursiveSolve(row, col, grid); solved {
			return solvedGrid, solved
		}
	}

	return &grid, false
}
