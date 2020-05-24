package sudoku

import (
	"strconv"
)

// Cell is a single square on the grid
// For convinience of getting the string value
type Cell struct {
	Value int
}

func (cell *Cell) String() string {
	if cell.Value != 0 {
		return strconv.Itoa(cell.Value)
	}
	return "-"
}
