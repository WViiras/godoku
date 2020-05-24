package sudoku

import (
	"bytes"
)

type GameArea struct {
	Fields [9][9]Field
}

func (gameArea *GameArea) String() string {
	var final bytes.Buffer

	for row, fields := range gameArea.Fields {
		for col, field := range fields {
			final.WriteString(field.String())
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
