package main

import (
	"fmt"
)

func printSudoku(puzzle [][]uint8) {
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


func main() {
	puzzle := [][]uint8{
		{3,0,2,0,0,0,0,0,0},
		{7,8,0,6,0,4,0,9,0},
		{6,4,0,0,2,5,3,0,0},
		{0,7,0,0,0,6,1,0,0},
		{0,0,0,3,0,2,0,0,0},
		{0,0,4,8,0,0,0,7,0},
		{0,0,7,9,6,0,0,1,4},
		{0,1,0,2,0,3,0,5,7},
		{0,0,0,0,0,0,6,0,9},
	}

	printSudoku(puzzle)
}
