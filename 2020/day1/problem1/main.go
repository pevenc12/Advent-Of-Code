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

	diff := map[int]bool{}
	res := 0

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		_, ok := diff[n]
		if ok {
			res = n * (2020 - n)
			break
		} else {
			diff[2020-n] = true
		}
	}

	fmt.Printf("answer is %d.", res)
}
