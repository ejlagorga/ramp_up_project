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
	//no number present	
	if puzzle[x][y] != 0 {
		return false
	}
	
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