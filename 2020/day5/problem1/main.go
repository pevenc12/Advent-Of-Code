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

	var res int = 0

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := 0
		col := 0
		rowAdd := 64
		colAdd := 4
		for _, char := range line {
			switch char {
			case 'F':
				rowAdd /= 2
			case 'B':
				row += rowAdd
				rowAdd /= 2
			case 'R':
				col += colAdd
				colAdd /= 2
			case 'L':
				colAdd /= 2
			}
		}
		position := row*8 + col
		if position > res {
			res = position
		}
	}

	fmt.Printf("The highest seat ID is %d.", res)
}
