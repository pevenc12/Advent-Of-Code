package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	color   string
	contain map[string]bool
}

func main() {
	// read the input file
	f, _ := os.Open("./input")
	defer f.Close()

	answer := map[string]bool{}
	nodeMap := map[string]node{}

	// scan line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		parent := ""
		child := []string{}
		for i, word := range words {
			if word == "bag" || word == "bag," || word == "bag." || word == "bags" || word == "bags," || word == "bags." {
				if i == 2 {
					parent = words[i-2] + " " + words[i-1]
				} else {
					child = append(child, words[i-2]+" "+words[i-1])
				}
			}
			_, ok := nodeMap[parent]
			if !ok {
				newNode := node{
					color:   parent,
					contain: map[string]bool{},
				}
				nodeMap[parent] = newNode
			}
			for _, c := range child {
				if c == "no other" {
					continue
				}
				n, ok := nodeMap[c]
				if !ok {
					newNode := node{
						color:   c,
						contain: map[string]bool{parent: true},
					}
					nodeMap[c] = newNode
				} else {
					n.contain[parent] = true
				}
			}
		}
	}

	// BFS
	var q []string
	q = append(q, "shiny gold")

	for len(q) != 0 {
		var name string
		name, q = q[0], q[1:]
		n, ok := nodeMap[name]
		if ok {
			if n.color != "shiny gold" {
				answer[n.color] = true
			}
			for key := range n.contain {
				q = append(q, key)
			}
		}
	}

	fmt.Printf("total color number is %d.", len(answer))
}
