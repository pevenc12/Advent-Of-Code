package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// read the input file
	f, _ := os.Open("./input")
	defer f.Close()

	seats := [][]rune{}
	next := [][]rune{}

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune{}
		for _, char := range line {
			if char == 'L' {
				row = append(row, '#')
			} else {
				row = append(row, '.')
			}
		}
		next = append(next, row)
	}

	rows := len(next)
	cols := len(next[0])

	for i := 0; i < rows; i++ {
		row := []rune{}
		for j := 0; j < cols; j++ {
			row = append(row, '.')
		}
		seats = append(seats, row)
	}

	// iterate
	for !check(seats, next) {
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				seats[i][j] = next[i][j]
			}
		}
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				adj := findocc(seats, i, j, rows, cols)
				if i == 1 && j == 1 {
				}
				switch seats[i][j] {
				case 'L':
					if adj == 0 {
						next[i][j] = '#'
					} else {
						next[i][j] = 'L'
					}
				case '#':
					if adj >= 4 {
						next[i][j] = 'L'
					} else {
						next[i][j] = '#'
					}
				default:
					continue
				}
			}
		}
	}

	var res int = 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if seats[i][j] == '#' {
				res++
			}
		}
	}

	fmt.Printf("The answer is %d.", res)
}

func check(s1 [][]rune, s2 [][]rune) bool {
	var same bool = true
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1[0]); j++ {
			if s1[i][j] != s2[i][j] {
				same = false
				break
			}
		}
		if !same {
			break
		}
	}
	return same
}

func findocc(seats [][]rune, i int, j int, rows int, cols int) int {
	occ := 0
	if i-1 >= 0 && j-1 >= 0 && seats[i-1][j-1] == '#' {
		occ++
	}
	if i-1 >= 0 && seats[i-1][j] == '#' {
		occ++
	}
	if i-1 >= 0 && j+1 < cols && seats[i-1][j+1] == '#' {
		occ++
	}
	if j-1 >= 0 && seats[i][j-1] == '#' {
		occ++
	}
	if j+1 < cols && seats[i][j+1] == '#' {
		occ++
	}
	if i+1 < rows && j-1 >= 0 && seats[i+1][j-1] == '#' {
		occ++
	}
	if i+1 < rows && seats[i+1][j] == '#' {
		occ++
	}
	if i+1 < rows && j+1 < cols && seats[i+1][j+1] == '#' {
		occ++
	}
	return occ
}
