# Sudoku Solver

This tiny program is a solution for the [sudoku problem](https://github.com/01-edu/public/tree/master/subjects/sudoku) in Golang.

To test the program, run it with a set of valid rows (9) of the sudoku as the arguments:

```
❯ go run . "34.91..2." ".96.8..41" "..8.2..7." ".6..57.39" "1.2.6.7.." "97..3..64" "45.2.8..6" ".8..9..5." "6.3..189."
```

If the input is valid, the program will print the solution for the sudoku.

(All characters other than 1-9 are treated as an empty cell. Each row must be 9 characters long. Using unicode characters should throw an error.)

In case of an invalid input, the program prints an error message and quits:

```
❯ go run . ".7....28." ".2...6.57" "8654729.." "..925..64" ".4..19.7." "7.8..4..9" "3..7..698" "..79.1..." "59..28.39"
Entered sudoku is not valid.

❯ go run . ".932..8." "27.3.85.." ".8.73.254" "9758...31" "....74.6." "6.45.38.7" "7....2.48" "32.4...7." "..8.579.."
Failed to create a sudoku: Row at position 0 is invalid

❯ go run . "not" "a" "sudoku"
Failed to create a sudoku: Invalid array of rows
```
