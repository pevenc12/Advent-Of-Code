package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// read the input file
	f, _ := os.Open("./input")
	defer f.Close()

	var valid int = 0
	var hasCid bool = false
	var field int = 0

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if string(line) == "" {
			if field == 8 || (field == 7) && hasCid == false {
				valid++
			}
			hasCid = false
			field = 0
			continue
		}
		pairs := strings.Split(line, " ")
		for _, pair := range pairs {
			key := strings.Split(pair, ":")[0]
			if key == "cid" {
				hasCid = true
			}
			field++
		}
	}

	if field == 8 || (field == 7) && hasCid == false {
		valid++
	}

	fmt.Printf("There are %d valid passport(s).", valid)
}
