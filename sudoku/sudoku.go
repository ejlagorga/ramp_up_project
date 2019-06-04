// Package sudoku comtains functions for reading, writing, and solving sudoku puzzles
package sudoku

import (
	"fmt"
	"github.com/lukpank/go-glpk/glpk"
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

func integer(puzzle Sudoku) (bool, Sudoku) {
	// initial new linear programing problem
	solver := glpk.New()
	// set optimization goal?
	solver.SetObjDir(glpk.MAX)

	// count the number of hints in the puzzle
	hints := 0
	for i, row := range puzzle {
		for j, _ := range row {
			if puzzle[i][j] > 0 {
				hints++
			}
		}
	}

	// create #(hints) + 4*9*9 empty constrains
	constraints := hints + 4*9*9
	solver.AddRows(constraints)
	// set each constraint equal to 1, effectively making the problem and binary programing problem
	for i := 1; i <= constraints; i++ {
		solver.SetRowBnds(i, glpk.FX, 1, 1)
	}

	// add 9x9x9 tensor of variables representing indicator vars for each cell/value pair
	variables := 9*9*9
	solver.AddCols(variables)
	for i := 1; i <= variables; i++ {
		// declare all variables as binary vars i.e. {0,1}
		solver.SetColKind(i, glpk.BV)
		// set col boundries?
		solver.SetColBnds(i, glpk.DB, 0, 1)
		// set objective function coefficients?
		solver.SetObjCoef(i, 0)
	}

	ia := make([]int32, 100000)
	ja := make([]int32, 100000)
	ar := make([]float64, 100000)

	p := 1;
	rno := int32(1);
	
	// hint constraints
    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if(puzzle[r][c] > 0) {
				ia[p] = rno
				rno++
				ja[p] = int32(r*9*9 + c*9 + int(puzzle[r][c]))
				ar[p] = 1.0
                p++
            }
        }
    }
    
	// Single value constraints
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            for k := 0; k < 9; k++ {
				ia[p] = rno
				ja[p] = int32(i*9*9 + j*9 + k + 1)
				ar[p] = 1.0
                p++
            }
            rno++
        }
    }

	// row constraints
    for i := 0; i < 9; i++ {
    	for k := 0; k < 9; k++ {
    		for j := 0; j < 9; j++ {
				ia[p] = rno
				ja[p] = int32(i*9*9 + j*9 + k + 1)
				ar[p] = 1.0
                p++
            }
            rno++
        }
	}

	// col constraints
 	for j := 0; j < 9; j++ {
    	for k := 0; k < 9; k++ {
    		for i := 0; i < 9; i++ {
				ia[p] = rno
				ja[p] = int32(i*9*9 + j*9 + k + 1)
				ar[p] = 1.0
                p++
            }
            rno++
        }

	}

	// box constraints
    for I := 0; I < 9; I+=3 {
    	for J := 0; J < 9; J+=3 {
    		for k := 0; k < 9; k++ {
                for i := I; i<I+3; i++ {
                    for j := J; j<J+3; j++ {
						ia[p] = rno
						ja[p] = int32(i*9*9 + j*9 + k + 1)
						ar[p] = 1.0
                        p++
                    }
                }
                rno++
            }
        }
	}
	
	// load constraints into solver
	solver.LoadMatrix(ia[:p], ja[:p], ar[:p])

	// set solver parameters for integer optimization
	parameters := glpk.NewIocp()
	parameters.SetPresolve(true)
	if error := solver.Intopt(parameters); error != nil {
		return false, puzzle
	}

	// recreate 2d representation from tensor form
	for i := 0; i < 9; i ++ {
		for j := 0; j < 9; j ++ {
			for k := 0; k < 9; k++ {
				if solver.MipColVal(i + 9*j + 81*k + 1) == 1 {
					puzzle[i][j] = uint8(k+1)
				}
			}
		}
	} 

	return true, puzzle
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

func (puzzle *Sudoku) Solve(algo string) {
	if algo == "dfs" {
		solved, p := dfs(*puzzle, 0, 0)
		if !solved { *puzzle = nil } else { *puzzle = p }
	} else if algo == "integer" {
		solved, p := integer(*puzzle)
		if !solved { *puzzle = nil } else { *puzzle = p }
	} else {
		fmt.Print("Try \"sudoku.Solve(dfs)\" or \"sudoku.Solve(dfs)\" " )
		return
	}
}