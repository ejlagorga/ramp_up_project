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