package main

import "os"

func main() {
	rows := os.Args[1:]
	sudoku := NewSudoku(rows)

	sudoku.Print()
}
