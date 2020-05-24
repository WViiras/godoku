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
		panic(error)
	}

	unsolved := string(file)
	unsolvedGrid := strings.Fields(unsolved)

	game := new(sudoku.Game)
	game.PopulateGrid(unsolvedGrid)
	game.Solve()
}
