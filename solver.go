package main

import "fmt"

// Sudoku ...
type Sudoku [][]int

// Print prints the sudoku
func (s Sudoku) Print() {
	for _, line := range s {
		for _, num := range line {
			if num == 0 {
				fmt.Print("- ")
			} else {
				fmt.Printf("%d ", num)
			}
		}
		fmt.Println()
	}
}

// ValidNums returns the valid nums
func (s Sudoku) ValidNums(x, y int) []int {
	cols := []int{}
	rows := []int{}
	sqr := []int{}

	for i := 0; i < len(s); i++ {

	}

	return
}

// Solve ...
func (s Sudoku) Solve() Sudoku {
	if len(s) != len(s[0]) {
		panic("must be square")
	}

	for x := 0; x < len(s); x++ {
		for y := 0; y < len(s); y++ {
			// we don't care
			if s[x][y] != 0 {
				continue
			}

			num := s.solve(x, y)
			s[x][y] = num
		}
	}

	return s
}

func (s Sudoku) solve(x, y int) int {
	nums := s.ValidNums(x, y)

	// no valid nums
	if len(nums) == 0 {
		return -1
	}

	num := -1
outer:
	for _, num = range nums {
		// assume it's the current number
		s[x][y] = num
		for i := x; i < len(s); i++ {
			for j := y + 1; j < len(s); j++ {
				// we don't care
				if s[i][j] != 0 {
					continue
				}

				guess := s.solve(i, j)
				if guess == -1 {
					// we fucked up
					goto outer
				}
			}
		}
	}

	return num
}
