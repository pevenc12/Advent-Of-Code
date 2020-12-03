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

	var p1, p2, p3, p4, p5 int = 0, 0, 0, 0, 0
	var t1, t2, t3, t4, t5 int = 0, 0, 0, 0, 0
	var even bool = true

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if string(line[p1]) == "#" {
			t1++
		}
		if string(line[p2]) == "#" {
			t2++
		}
		if string(line[p3]) == "#" {
			t3++
		}
		if string(line[p4]) == "#" {
			t4++
		}
		if even == true && string(line[p5]) == "#" {
			t5++
		}
		n := len(line)
		p1 = (p1 + 1) % n
		p2 = (p2 + 3) % n
		p3 = (p3 + 5) % n
		p4 = (p4 + 7) % n
		if even != true {
			p5 = (p5 + 1) % n
		}
		even = !even
	}

	fmt.Printf("The answer is %d.", t1*t2*t3*t4*t5)
}
