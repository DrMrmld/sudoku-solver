package main

import (
	"os"
	"fmt"
)

func main() {
	rows := os.Args[1:]
	sudoku, err := NewSudoku(rows)
	if err != nil {
		fmt.Printf("Failed to create a sudoku: %s\n", err)
		return
	}

	if !sudoku.IsValid() {
		fmt.Println("Entered sudoku is not valid.")
		return
	}

	sudoku.Solve().Print()
}
