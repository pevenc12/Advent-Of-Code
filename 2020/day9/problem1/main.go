package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "strings"
)

func main() {
	// read the input file
	f, _ := os.Open("./input")
	defer f.Close()

	var res int = 0
	arr := []int{}

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		if len(arr) < 25 {
			arr = append(arr, num)
			continue
		} else {
			diff := map[int]bool{}
			exist := false
			for _, n := range arr {
				_, ok := diff[num-n]
				if !ok {
					diff[n] = true
				} else {
					exist = true
					break
				}
			}
			if exist {
				arr = arr[1:]
				arr = append(arr, num)
			} else {
				res = num
				break
			}
		}
	}

	fmt.Printf("The first invalid number is %d.", res)
}
