package sudoku

import "fmt"

type Game struct {
	gameArea *GameArea
}

func (game *Game) PopulateGameArea(unsolvedFields []string) {
	game.gameArea = new(GameArea)

	col := 0
	row := 0
	for _, v := range unsolvedFields {
		for _, c := range v {
			intVal := 0
			if c != '-' {
				intVal = int(c - '0')
			}
			game.gameArea.Fields[row][col] = Field{
				Value:    intVal,
				Possible: make([]int, 0, 9),
			}
			col++
			if col >= 9 {
				row++
				col = 0
			}
		}
	}

	fmt.Println(game.gameArea)
}
