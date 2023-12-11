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

var (
	nodeMap       = make(map[string]Node)
	startingNodes = []string{}
	returnNodes   = []string{}
)

func main() {
	contents, _ := os.ReadFile("testInput.txt")
	// contents, _ := os.ReadFile("../Day8/input.txt")

	lines := strings.Split(string(contents), "\n")
	directions := strings.TrimSpace(lines[0])
	fmt.Println(directions)

	for i := 2; i < len(lines)-1; i++ {
		line := strings.Split(lines[i], " = ")
		head := line[0]
		if string(head[2]) == "A" {
			startingNodes = append(startingNodes, head)
		}
		left := line[1][1:4]
		right := line[1][6:9]
		node := Node{left: left, right: right}
		nodeMap[head] = node
	}

	fmt.Println(startingNodes)

	turns := 0
	turnPointer := 0
	isDone := false
	for isDone == false {
		// for turns < 10 {
		fmt.Println("Looking:", string(directions[turnPointer]))
		for i, node := range startingNodes {
			fmt.Println("Eval:", node)
			val := followMapMove(node, string(directions[turnPointer]))
			fmt.Println("Return value:", val)
			startingNodes[i] = val
		}
		fmt.Println("Return vals:", startingNodes)
		turnPointer++
		if turnPointer == len(directions) {
			turnPointer = 0
		}
		turns++
		isDone = checkRetVals(startingNodes)
		fmt.Println("Is Done:", isDone)
		fmt.Println("Turns:", turns)
		fmt.Println()
		// isDone = true
	}
	fmt.Println(turns)
}

func followMapMove(startingNode, direction string) string {
	currentKey := startingNode
	currentValue := ""
	if direction == "L" {
		currentValue = nodeMap[currentKey].left
	} else {
		currentValue = nodeMap[currentKey].right
	}
	return currentValue
}

func checkRetVals(retVals []string) bool {
	for _, val := range retVals {
		if string(val[2]) != "Z" {
			return false
		}
	}
	return true
}
