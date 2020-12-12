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

type process struct {
	current int
	accu    int
	fixed   bool
	history map[int]bool
}

func main() {
	// read the input file
	f, _ := os.Open("./input")
	defer f.Close()

	// pos - (length - 1) of the instructions
	var pos int = 0
	var res int
	var totalIns = []instruction{}

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		operation := strings.Split(line, " ")[0]
		num := strings.Split(line, " ")[1]
		argument, _ := strconv.Atoi(num)
		totalIns = append(totalIns, instruction{op: operation, arg: argument})
		pos++
	}

	pos--

	// BFS
	q := []process{{current: 0, accu: 0, fixed: false, history: map[int]bool{}}}
	for {
		var p process
		p, q = q[0], q[1:]

		_, ok := p.history[p.current]
		if ok {
			continue
		}

		if p.current == pos {
			res = p.accu
			break
		} else if p.current > pos {
			continue
		} else if p.current < 0 {
			continue
		}

		p.history[p.current] = true

		newHis := map[int]bool{}
		for key, val := range p.history {
			newHis[key] = val
		}

		ins := totalIns[p.current]
		if ins.op == "acc" {
			q = append(q, process{current: p.current + 1, accu: p.accu + ins.arg, fixed: p.fixed, history: newHis})
		} else {
			if p.fixed {
				newP := process{current: p.current, accu: p.accu, fixed: p.fixed, history: newHis}
				switch ins.op {
				case "jmp":
					newP.current += ins.arg
				case "nop":
					newP.current++
				}
				q = append(q, newP)
			} else {
				switch ins.op {
				case "jmp":
					q = append(q, process{current: p.current + ins.arg, accu: p.accu, fixed: false, history: newHis})
					q = append(q, process{current: p.current + 1, accu: p.accu, fixed: true, history: newHis})
				case "nop":
					q = append(q, process{current: p.current + 1, accu: p.accu, fixed: false, history: newHis})
					q = append(q, process{current: p.current + ins.arg, accu: p.accu, fixed: true, history: newHis})
				}
			}
		}
	}

	// last instruction
	if ins := totalIns[pos]; ins.op == "acc" {
		res += ins.arg
	}

	fmt.Printf("acc is %d.", res)
}
