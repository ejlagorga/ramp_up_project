// Package sudoku comtains functions for reading, writing, and solving sudoku puzzles
package sudoku

import (
	"testing"
)

var easy = Sudoku{
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

var medium = Sudoku{
	{6,0,0,0,0,2,0,0,9},
	{0,0,0,0,0,7,4,0,0},
	{0,1,9,0,0,0,5,8,0},
	{4,8,0,0,0,3,0,0,5},
	{0,0,7,0,0,0,8,0,0},
	{9,0,0,5,0,0,0,1,4},
	{0,7,2,0,0,0,3,6,0},
	{0,0,6,1,0,0,0,0,0},
	{1,0,0,2,0,0,0,0,7},
}

var hard = Sudoku{
	{3,0,0,1,0,0,8,0,4},
	{0,5,6,8,3,0,0,0,0},
	{0,0,0,0,6,0,0,3,0},
	{0,9,0,0,0,0,0,0,0},
	{0,1,3,5,0,4,9,7,0},
	{0,0,0,0,0,0,0,8,0},
	{0,6,0,0,7,0,0,0,0},
	{0,0,0,0,4,8,5,2,0},
	{7,0,4,0,0,1,0,0,6},
}

var evil = Sudoku{
	{0,7,0,8,9,0,3,0,0},
	{0,0,1,0,0,0,7,2,0},
	{0,0,0,0,4,0,0,0,0},
	{5,0,0,0,0,6,2,8,0},
	{3,0,0,0,0,0,0,0,4},
	{0,2,6,1,0,0,0,0,7},
	{0,0,0,0,2,0,0,0,0},
	{0,6,5,0,0,0,4,0,0},
	{0,0,9,0,7,8,0,5,0},
}

var impossible = Sudoku{
	{0,7,0,0,0,6,0,0,0},
	{9,0,0,0,0,0,0,4,1},
	{0,0,8,0,0,9,0,5,0},
	{0,9,0,0,0,7,0,0,2},
	{0,0,3,0,0,0,8,0,0},
	{4,0,0,8,0,0,0,1,0},
	{0,8,0,3,0,0,9,0,0},
	{1,6,0,0,0,0,0,0,7},
	{0,0,0,5,0,0,0,8,0},
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
		got := easy.validateNumber(c.n, c.x, c.y)
		if got != c.want {
			t.Errorf("ValidateNumber(%v,%v,%v) == %v", c.n, c.x, c.y, c.want)
		}
	}
}

func TestSolve(t *testing.T) {
	cases := []struct {
		puzzle Sudoku
		name, algo string
		want bool
	}{
		// dfs tests
		{easy, "easy", "dfs", true},
		{medium, "medium", "dfs", true},
		{hard, "hard", "dfs", true},
		{evil, "evil", "dfs", true},
		{impossible, "impossible", "dfs", false},

		// integer programing tests
		{easy, "easy", "integer", true},
		{medium, "medium", "integer", true},
		{hard, "hard", "integer", true},
		{evil, "evil", "integer", true},
		{impossible, "impossible", "integer", false},
	}
	for _, c := range cases {
		got := true
		c.puzzle.Solve(c.algo)

		if c.puzzle == nil {
			got = false
		}

		if got != c.want {
			t.Errorf("%s was incorrectly identified", c.name)
		}
	}

}