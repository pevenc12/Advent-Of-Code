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

	var position int = 0
	var tree int = 0

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if string(line[position]) == "#" {
			tree++
		}
		position = (position + 3) % len(line)
	}

	fmt.Printf("There are %d tree(s).", tree)
}
