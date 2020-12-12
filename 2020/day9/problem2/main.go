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

	var invalid int = 0
	arr := []int{}
	totalArr := []int{}

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		totalArr = append(totalArr, num)
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
				invalid = num
				break
			}
		}
	}

	var temp int = 0
	var left int = 0
	var right int = -1
	arr = []int{}
	for temp != invalid {
		if temp < invalid {
			right++
			temp += totalArr[right]
		} else if temp > invalid {
			left++
			temp -= totalArr[left-1]
		}
	}

	maxN, minN := totalArr[left], totalArr[left]
	for i := left; i <= right; i++ {
		if totalArr[i] > maxN {
			maxN = totalArr[i]
		}
		if totalArr[i] < minN {
			minN = totalArr[i]
		}
	}

	fmt.Printf("The answer is %d.", maxN+minN)
}
