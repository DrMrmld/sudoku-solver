package main

import (
	"os"
	"fmt"
)

func main() {
	rows := os.Args[1:]
	sudoku, err := NewSudoku(rows)
	if err != nil {
		fmt.Println(err)
		return
	}

	sudoku.Print()
	fmt.Println(sudoku.IsValid())
}
