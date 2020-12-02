package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open a file
	f, _ := os.Open("./input")
	defer f.Close()

	var valid int = 0

	// scanner scans the file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		rangeN := strings.SplitN(input[0], "-", 2)
		minN, _ := strconv.Atoi(rangeN[0])
		maxN, _ := strconv.Atoi(rangeN[1])
		target := []rune(strings.Split(input[1], ":")[0])[0]
		pwd := input[2]
		num := 0

		for _, char := range pwd {
			if char == target {
				num++
			}
			if num > maxN {
				break
			}
		}

		if num >= minN && num <= maxN {
			valid++
		}
	}

	fmt.Printf("There are %d valid password(s).", valid)
}
