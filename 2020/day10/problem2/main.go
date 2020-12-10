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

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		arr = append(arr, num)
	}
	sort.Ints(arr)
	arr = append(arr, arr[len(arr)-1]+3)

	dp := []int{}
	for i := range arr {
		if i == 0 {
			dp = append(dp, 1)
			continue
		}
		var ways int = 0
		for j := i - 1; j >= i-3; j-- {
			if j < 0 || arr[i]-arr[j] > 3 {
				break
			}
			ways = ways + dp[j]
		}
		dp = append(dp, ways)
	}

	fmt.Printf("answer is %d.", dp[len(dp)-1])
}
