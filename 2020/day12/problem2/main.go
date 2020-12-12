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
	var wpx int = 10
	var wpy int = 1

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ins, a := line[0], line[1:]
		arg, _ := strconv.Atoi(a)
		switch ins {
		case 'F':
			x += wpx * arg
			y += wpy * arg
		case 'N':
			wpy += arg
		case 'E':
			wpx += arg
		case 'S':
			wpy -= arg
		case 'W':
			wpx -= arg
		case 'L':
			turn := (arg / 90) % 4
			switch turn {
			case 1:
				wpx, wpy = -wpy, wpx
			case 2:
				wpx, wpy = -wpx, -wpy
			case 3:
				wpx, wpy = wpy, -wpx
			}
		case 'R':
			turn := (arg / 90) % 4
			switch turn {
			case 1:
				wpx, wpy = wpy, -wpx
			case 2:
				wpx, wpy = -wpx, -wpy
			case 3:
				wpx, wpy = -wpy, wpx
			}
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
