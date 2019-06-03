// Package sudoku comtains functions for reading, writing, and solving sudoku puzzles
package sudoku

import (
	"fmt"
)

type Sudoku [][]uint8

func (puzzle Sudoku) Print() {
	fmt.Println("╔═══╤═══╤═══╦═══╤═══╤═══╦═══╤═══╤═══╗")

	for i, row := range puzzle {
		fmt.Print("║ ")	

		for j, _ := range row {

			if puzzle[i][j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(puzzle[i][j])
			}
			
			if (j+1) % 3 == 0 {
				fmt.Print(" ║ ")
			} else {
				fmt.Print(" │ ")
			}
		}
		
		if (i+1) % 3 == 0 {
			if i < 8 { fmt.Println("\n╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣") }
		} else {
			fmt.Println("\n╟───┼───┼───╫───┼───┼───╫───┼───┼───╢")
		}
	}

	fmt.Println("\n╚═══╧═══╧═══╩═══╧═══╧═══╩═══╧═══╧═══╝")
}

func (puzzle Sudoku) validateNumber(n, x, y uint8) bool {
	//no matching num in row
	for _, v := range puzzle[x] {
		if v == n { return false }
	}
	//no matching num in col
	for _, row := range puzzle {
		if row[y] == n { return false }
	}

	//no matching num in box
	for i := 3*(x/3); i < 3*(x/3)+3; i++ {
		for j := 3*(y/3); j < 3*(y/3)+3; j++ {
			if puzzle[i][j] == n { return false }
		}
	}

	return true
}

func dfs(puzzle Sudoku, x, y uint8) (bool, Sudoku) {
	// return true if end of puzzle reached
	if x >= 9 {
		return true, puzzle
	}
	
	// if space is filled, move to next space
	if puzzle[x][y] > 0 {
		solved, puzzle := dfs(puzzle, x+(y/8), (y+1)%9)
		return solved, puzzle
	}

	// loop over potential values
	for num := uint8(1); num <= 9; num++ {
		// if value is possible, update puzzle and call dfs
		if puzzle.validateNumber(num, x, y) {
			puzzle[x][y] = num
			// if puzzle solution is found, return
			if solved, puzzle := dfs(puzzle, x+(y/8), (y+1)%9); solved {
				return solved, puzzle
			}
		}
	}

	// if value can be put in space, backtrack
	puzzle[x][y] = 0
	return false, puzzle
}

func (puzzle *Sudoku) Solve() {
	solved, p := dfs(*puzzle, 0, 0)
	if !solved {
		*puzzle = nil
	} else {
		*puzzle = p
	}
}