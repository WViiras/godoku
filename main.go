package main

import (
	"fmt"
	"godoku/sudoku"
	"io/ioutil"
	"strings"
)

func main() {
	file, error := ioutil.ReadFile("unsolved")

	if error != nil {
		fmt.Println(error)
	}

	unsolved := string(file)
	gameArea := new(sudoku.GameArea)

	populateGameArea(&unsolved, gameArea)

	fmt.Println(gameArea)
}

func populateGameArea(unsolved *string, gameArea *sudoku.GameArea) {
	unsolvedFields := strings.Fields(*unsolved)
	col := 0
	row := 0
	for _, v := range unsolvedFields {
		for _, c := range v {
			intVal := 0
			if c != '-' {
				intVal = int(c - '0')
			}
			gameArea.Fields[row][col] = sudoku.Field{
				Value:    intVal,
				Possible: make([]int, 0, 9),
			}
			col++
		}
		if col >= 9 {
			row++
			col = 0
		}
	}
}
