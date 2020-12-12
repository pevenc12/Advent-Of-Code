package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	op  string
	arg int
}

func main() {
	// read the input file
	f, _ := os.Open("./input")
	defer f.Close()

	var res int = 0
	var pos int = 0
	var totalIns = []instruction{}
	history := map[int]bool{}

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		operation := strings.Split(line, " ")[0]
		num := strings.Split(line, " ")[1]
		argument, _ := strconv.Atoi(num)
		totalIns = append(totalIns, instruction{op: operation, arg: argument})
	}

	for {
		_, ok := history[pos]
		if ok {
			break
		}
		history[pos] = true
		ins := totalIns[pos]
		switch ins.op {
		case "nop":
			pos++
		case "acc":
			res += ins.arg
			pos++
		case "jmp":
			pos += ins.arg
		}
	}

	fmt.Printf("acc is %d.", res)
}
