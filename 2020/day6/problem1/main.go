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

	var count int = 0
	answer := map[rune]bool{}

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println(len(answer))
			count += len(answer)
			answer = make(map[rune]bool)
			continue
		}
		for _, char := range line {
			answer[char] = true
		}
	}
	count += len(answer)
	fmt.Printf("The sum of those counts = %d", count)
}
