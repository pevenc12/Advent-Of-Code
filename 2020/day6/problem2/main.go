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
	var member int = 0
	answer := map[rune]int{}

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			count += checkCount(answer, member)
			answer = make(map[rune]int)
			member = 0
			continue
		}
		for _, char := range line {
			_, ok := answer[char]
			if ok {
				answer[char]++
			} else {
				answer[char] = 1
			}
		}
		member++
	}
	count += checkCount(answer, member)
	fmt.Printf("The sum of those counts = %d", count)
}

func checkCount(_map map[rune]int, member int) int {
	ans := 0
	for _, value := range _map {
		if value == member {
			ans++
		}
	}
	return ans
}
