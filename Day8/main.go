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
	contents, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(contents), "\n")
	directions := strings.TrimSpace(lines[0])

	firstKey := "AAA"

	nodeMap := make(map[string]Node)
	for i := 2; i < len(lines)-1; i++ {
		line := strings.Split(lines[i], " = ")
		head := line[0]
		left := line[1][1:4]
		right := line[1][6:9]
		node := Node{left: left, right: right}
		nodeMap[head] = node
	}

	turns := 0
	turnPointer := 0
	currentKey := firstKey
	currentValue := ""
	for currentValue != "ZZZ" {
		if string(directions[turnPointer]) == "L" {
			currentValue = nodeMap[currentKey].left
			currentKey = currentValue
		} else {
			currentValue = nodeMap[currentKey].right
			currentKey = currentValue
		}
		turnPointer++
		if turnPointer == len(directions) {
			turnPointer = 0
		}
		turns++
	}
	fmt.Println(turns)
}
