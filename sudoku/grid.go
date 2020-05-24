package sudoku

import (
	"bytes"
)

// Grid holds all the cells
type Grid struct {
	Cells [9][9]Cell
}

func (grid *Grid) String() string {
	var final bytes.Buffer

	for row, cells := range grid.Cells {
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
