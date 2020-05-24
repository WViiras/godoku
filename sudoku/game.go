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
