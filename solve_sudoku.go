package main

import (
	"ramp_up_project/sudoku"
)

var easy = sudoku.Sudoku{
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

func main() {
	puzzle := easy
	puzzle.Solve("integer")
	puzzle.Print()
}