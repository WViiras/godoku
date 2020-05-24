package main

import (
	"fmt"
	"godoku/sudoku"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing file parameter, please provide a path to file")
		return
	}

	sudokuFilePath, filepathErr := filepath.Abs(os.Args[1])
	if filepathErr != nil {
		fmt.Println(filepathErr)
		return
	}

	fmt.Println("Searching for file '", sudokuFilePath, "'")
	fmt.Println()

	file, readfileErr := ioutil.ReadFile(sudokuFilePath)
	if readfileErr != nil {
		fmt.Println(readfileErr)
		return
	}

	unsolved := string(file)
	unsolvedGrid := strings.Fields(unsolved)

	game := new(sudoku.Game)
	game.PopulateGrid(unsolvedGrid)
	game.Solve()
}
