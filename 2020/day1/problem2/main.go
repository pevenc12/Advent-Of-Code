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

	arr := []int{}
	res := 0

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		arr = append(arr, n)
	}

	fin := false
	for i, n := range arr {
		diff := map[int]bool{}
		for j := i + 1; j < len(arr); j++ {
			_, ok := diff[arr[j]]
			if ok {
				res = n * (arr[j]) * (2020 - n - arr[j])
				fin = true
				break
			} else {
				diff[2020-n-arr[j]] = true
			}
		}
		if fin {
			break
		}
	}

	fmt.Printf("answer is %d.", res)
}
