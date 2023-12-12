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
)

func main() {
	contents, _ := os.ReadFile("../Day8/input.txt")

	lines := strings.Split(string(contents), "\n")
	directions := strings.TrimSpace(lines[0])

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

	turns := []int{}
	for _, node := range startingNodes {
		numOfTurns := followMapMove(node, directions)
		turns = append(turns, numOfTurns)
	}

	ans := LCM(turns)
	fmt.Println(ans)
}

func LCM(turns []int) int {
	result := turns[0]

	for i := 0; i < len(turns); i++ {
		result = result * turns[i] / GCD(result, turns[i])
	}

	return result
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func followMapMove(startingNode, directions string) int {
	turns := 0
	turnPointer := 0
	currentValue := startingNode
	for string(currentValue[2]) != "Z" {
		if string(directions[turnPointer]) == "L" {
			currentValue = nodeMap[currentValue].left
		} else {
			currentValue = nodeMap[currentValue].right
		}
		turnPointer++
		if turnPointer == len(directions) {
			turnPointer = 0
		}
		turns++
	}
	return turns
}

func checkRetVals(retVals []string) bool {
	for _, val := range retVals {
		if string(val[2]) != "Z" {
			return false
		}
	}
	return true
}
