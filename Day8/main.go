package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

func main() {
	contents, _ := os.ReadFile("testInput1.txt")
	// contents, _ := os.ReadFile("testInput2.txt")
	// contents, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(contents), "\n")

	fmt.Println(string(lines[0]))

	// populate map
	nodeMap := make(map[string]Node)
	for i := 2; i < len(lines)-1; i++ {
		line := strings.Split(lines[i], " = ")
		head := line[0]
		left := line[1][1:4]
		right := line[1][6:9]
		node := Node{left: left, right: right}
		nodeMap[head] = node
	}
}

// while left/right node val (depending on current instruction) != ZZZ
// look up map value and children (while we aren't lookin at ZZZ and pointer != len (lines[0]))
// if pointing at the last element reset pointer to beginning of lines[0]
// Keep track of number of turns to find ZZZ
