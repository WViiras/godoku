package sudoku

import "strconv"

type Field struct {
	Value    int
	Possible []int
}

func (field *Field) String() string {
	if field.Value != 0 {
		return strconv.Itoa(field.Value)
	}
	return "-"
}
