// Package sudoku comtains functions for reading, writing, and solving sudoku puzzles
package sudoku

import (
	"testing"
)

var puzzle = Sudoku{
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

func TestValidateNumber(t *testing.T) {
	cases := []struct {
		n, x, y uint8
		want bool
	}{
		//row collision
		{9, 1, 2, false},
		//col collision
		{1, 0, 1, false},
		//box1 collision 
		{6, 0, 1, false},
		//box5 collision 
		{3, 3, 3, false},
		//box9 collision 
		{6, 6, 6, false},
		//no collision
		{5, 4, 4, true},
	}
	for _, c := range cases {
		got := puzzle.validateNumber(c.n, c.x, c.y)
		if got != c.want {
			t.Errorf("ValidateNumber(%v,%v,%v) == %v", c.n, c.x, c.y, c.want)
		}
	}
}