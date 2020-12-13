package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read the input file
	f, _ := os.Open("./input")
	defer f.Close()

	var arrive int = 0
	var distance int = 100
	var res int = 0

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		a, err := strconv.Atoi(line)
		if err == nil {
			arrive = a
		} else {
			nums := strings.Split(line, ",")
			for _, num := range nums {
				n, err := strconv.Atoi(num)
				if err != nil {
					continue
				}
				d := n - (arrive % n)
				if d < distance {
					res = d * n
					distance = d
				}
			}
		}
	}

	fmt.Printf("answer is %d.", res)
}
