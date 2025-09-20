package main

func (s Sudoku) Solve() Sudoku {
	result := EmptySudoku()

	var addOne func(Sudoku)

	addOne = func(curr Sudoku) {
		if !curr.IsValid() {
			return
		}

		recursed := false

		for i, row := range curr {
			for j, val := range row {
				if val == 0 {
					recursed = true
					for _, num := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
						curr[i][j] = num
						addOne(curr)
					}
				}
			}
		}

		if !recursed {
			result = curr
		}
	}

	addOne(s)

	return result
}
