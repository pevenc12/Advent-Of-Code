package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// read the input file
	f, _ := os.Open("./input")
	defer f.Close()

	arr := []int{0}
	var diffone int = 0
	var diffthree int = 0

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		arr = append(arr, num)
	}
	sort.Ints(arr)
	arr = append(arr, arr[len(arr)-1]+3)

	for i := 0; i < len(arr)-1; i++ {
		switch arr[i+1] - arr[i] {
		case 1:
			diffone++
		case 3:
			diffthree++
		default:
			continue
		}
	}

	fmt.Printf("answer is %d.", diffone*diffthree)
}
