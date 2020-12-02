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
		minP, _ := strconv.Atoi(rangeN[0])
		maxP, _ := strconv.Atoi(rangeN[1])
		target := strings.Split(input[1], ":")[0]
		pwd := input[2]

		if len(pwd) < maxP {
			continue
		}

		if string(pwd[minP-1]) == target && string(pwd[maxP-1]) != target {
			valid++
		} else if string(pwd[minP-1]) != target && string(pwd[maxP-1]) == target {
			valid++
		}
	}

	fmt.Printf("There are %d valid password(s).", valid)
}
