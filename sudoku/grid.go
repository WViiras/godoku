package sudoku

import (
	"bytes"
)

// Grid holds all the cells
type Grid struct {
	cells [9][9]Cell
}

func (grid *Grid) String() string {
	var final bytes.Buffer

	for row, cells := range grid.cells {
		for col, cell := range cells {
			final.WriteString(cell.String())
			if (col+1)%3 == 0 {
				final.WriteString(" ")
			}
		}
		final.WriteString("\n")
		if (row+1)%3 == 0 {
			final.WriteString("\n")
		}
	}

	return final.String()
}

// CanBeInBox checks if value can be placed in the surrounding 3x3 box
func (grid *Grid) CanBeInBox(col, row, value int) bool {
	boxRowS := row - (row % 3)
	boxRowE := boxRowS + 3
	boxColS := col - (col % 3)
	boxColE := boxColS + 3

	for rowI := boxRowS; rowI < boxRowE; rowI++ {
		for colI := boxColS; colI < boxColE; colI++ {
			if grid.GetCellValue(rowI, colI) == value {
				return false
			}
		}
	}
	return true
}

// CanBeInColumn checks if value can be placed on given column in the grid
func (grid Grid) CanBeInColumn(col, row, value int) bool {
	for rowI := 0; rowI < 9; rowI++ {
		if grid.GetCellValue(rowI, col) == value {
			return false
		}
	}
	return true
}

// CanBeInRow checks if value can be placed on given row in the grid
func (grid Grid) CanBeInRow(col, row, value int) bool {
	for colI := 0; colI < 9; colI++ {
		if grid.GetCellValue(row, colI) == value {
			return false
		}
	}
	return true
}

// GetCellValue is a wrapper for prettier code
func (grid Grid) GetCellValue(row, col int) int {
	return grid.cells[row][col].Value
}

// SetCellValue is a wrapper for prettier code
func (grid *Grid) SetCellValue(row, col, value int) {
	grid.cells[row][col].Value = value
}
