package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// read the input file
	f, _ := os.Open("./input")
	defer f.Close()

	var x int = 0
	var y int = 0

	forward := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	var direction = 0
	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ins, a := line[0], line[1:]
		arg, _ := strconv.Atoi(a)
		switch ins {
		case 'F':
			x += forward[direction][0] * arg
			y += forward[direction][1] * arg
		case 'N':
			y += arg
		case 'E':
			x += arg
		case 'S':
			y -= arg
		case 'W':
			x -= arg
		case 'L':
			direction = (direction - (arg / 90) + 4) % 4
		case 'R':
			direction = (direction + (arg / 90)) % 4
		}
	}

	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}

	fmt.Printf("The Manhattan Distance is %d.", x+y)
}
