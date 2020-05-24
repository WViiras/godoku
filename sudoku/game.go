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
			game.grid.Cells[row][col] = Cell{
				Value: intVal,
			}
			col++
			if col >= 9 {
				row++
				col = 0
			}
		}
	}

	fmt.Println(game.grid)
}

// Solve returns a solution
func (game *Game) Solve() Grid {
	solvedGrid, _ := recursiveSolve(0, 0, *game.grid)
	return *solvedGrid
}

func recursiveSolve(row, col int, grid Grid) (*Grid, bool) {
	// fmt.Println(grid.String())
	// // if this cell already has a value, try next cell
	if grid.Cells[row][col].Value != 0 {
		col++
		if col >= 9 {
			col = 0
			row++
			if row >= 9 {
				// fmt.Println("last cell, returning")
				// fmt.Println(grid.String())
				return &grid, true
			}
		}
		return recursiveSolve(row, col, grid)
	}

	// recursively try all values
	for i := 1; i <= 9; i++ {
		if !canBeInBox(col, row, i, &grid) {
			continue
		}
		if !canBeInRow(col, row, i, &grid) {
			continue
		}
		if !canBeInColumn(col, row, i, &grid) {
			continue
		}
		// place value
		grid.Cells[row][col].Value = i

		// try next cell
		if solvedGrid, solved := recursiveSolve(row, col, grid); solved {
			return solvedGrid, true
		}
	}

	return &grid, false
}

// check if can be placed in same box
func canBeInBox(col, row, value int, grid *Grid) bool {
	boxRowS := row - (row % 3)
	boxRowE := boxRowS + 3
	boxColS := col - (col % 3)
	boxColE := boxColS + 3

	for rowI := boxRowS; rowI < boxRowE; rowI++ {
		for colI := boxColS; colI < boxColE; colI++ {
			if grid.Cells[rowI][colI].Value == value {
				return false
			}
		}
	}
	return true
}

// check if can be placed on same column
func canBeInColumn(col, row, value int, grid *Grid) bool {
	for rowI := 0; rowI < 9; rowI++ {
		if grid.Cells[rowI][col].Value == value {
			return false
		}
	}
	return true
}

// check if can be placed on same row
func canBeInRow(col, row, value int, grid *Grid) bool {
	for colI := 0; colI < 9; colI++ {
		if grid.Cells[row][colI].Value == value {
			return false
		}
	}
	return true
}
