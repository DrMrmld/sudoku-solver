package main

import "fmt"

type Sudoku [9][9]int

func EmptySudoku() Sudoku {
	return [9][9]int{}
}

func NewSudoku(rows []string) Sudoku {
	if rows == nil || len(rows) != 9 {
		return EmptySudoku()
	}

	for _, row := range rows {
		if len(row) != 9 {
			return EmptySudoku()
		}
	}

	sudoku := EmptySudoku()

	for i, row := range rows {
		for j, val := range row {
			if val < '1' || val > '9' {
				continue
			}

			sudoku[i][j] = int(val - '0')
		}
	}

	return sudoku
}

func (s *Sudoku) Print() {
	topSep := "."
	midSep := "|"
	botSep := "'"
	hor := "---"

	fmt.Print(topSep)
	for range s {
		fmt.Print(hor)
		fmt.Print(topSep)
	}
	fmt.Print("\n")
	for i := range s {
		fmt.Print(midSep)
		for j := range s[i] {
			fmt.Print(" ")
			fmt.Print(s[i][j])
			fmt.Print(" " + midSep)
		}
		fmt.Print("\n")

		sep := ""
		if i == len(s) - 1 {
			sep = botSep
		} else {
			sep = midSep
		}

		fmt.Print(sep)
		for range s[i] {
			fmt.Print(hor)
			fmt.Print(sep)
		}
		fmt.Print("\n")
	}
}
