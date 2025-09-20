package main

import "fmt"

type Sudoku [9][9]int

func EmptySudoku() Sudoku {
	return [9][9]int{}
}

func NewSudoku(rows []string) (Sudoku, error) {
	if rows == nil || len(rows) != 9 {
		return EmptySudoku(), fmt.Errorf("Invalid array of rows")
	}

	for i, row := range rows {
		if len(row) != 9 {
			return EmptySudoku(), fmt.Errorf("Row at position %d is invalid", i)
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

	return sudoku, nil
}

func (s *Sudoku) IsValid() bool {
	for _, row := range s {
		var seen [9]bool

		for _, val := range row {
			if val == 0 {
				continue
			}

			if seen[val - 1] {
				return false
			} else {
				seen[val - 1] = true
			}
		}
	}

	for j := range s {
		var seen [9]bool

		for i := range s {
			val := s[i][j]

			if val == 0 {
				continue
			}

			if seen[val - 1] {
				return false
			} else {
				seen[val - 1] = true
			}
		}
	}

	starts := [][]int{
		{0, 0},
		{0, 3},
		{0, 6},
		{3, 0},
		{3, 3},
		{3, 6},
		{6, 0},
		{6, 3},
		{6, 6},
	}

	diffs := [][]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 0},
		{2, 1},
		{2, 2},
	}

	for _, start := range starts {
		var seen [9]bool

		for _, diff := range diffs {
			i := start[0] + diff[0]
			j := start[1] + diff[1]

			val := s[i][j]

			if val == 0 {
				continue
			}

			if seen[val - 1] {
				return false
			} else {
				seen[val - 1] = true
			}
		}
	}

	return true
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
			if val := s[i][j]; val == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(val)
			}
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
