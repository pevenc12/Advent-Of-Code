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

	// scan line by line
	scanner := bufio.NewScanner(f)
	allSeats := map[int]bool{}
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
		allSeats[position] = true
	}

	for i := 1; i < 1023; i++ {
		_, ok := allSeats[i]
		if ok {
			continue
		}
		_, ok = allSeats[i-1]
		if !ok {
			continue
		}
		_, ok = allSeats[i+1]
		if !ok {
			continue
		}
		fmt.Printf("Your seat ID is %d.", i)
	}
}
