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

	offset := map[int]int{}

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		_, err := strconv.Atoi(line)
		if err == nil {
			continue
		}
		nums := strings.Split(line, ",")
		offs := 0
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err == nil {
				offset[n] = offs
			}
			offs++
		}
	}

	// generate WolframAlpha query
	for key, val := range offset {
		fmt.Printf("(t+%d) mod %d = 0,", val, key)
	}

	// solve by WolframAlpha
	// base := 1215475766514841
	offs := 327300950120029

	fmt.Printf("answer is %d.", offs)
}
