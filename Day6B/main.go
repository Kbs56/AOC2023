package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, _ := os.ReadFile("../Day6/input.txt")

	lines := strings.Split(string(contents), "\n")

	time := parseLine(lines[0])
	distance := parseLine(lines[1])

	numCombos := calcWinningCombos(time, distance)

	fmt.Println(numCombos)
}

func parseLine(line string) int {
	_, lineAfter, _ := strings.Cut(line, ":")
	lineStr := strings.Fields(lineAfter)
	num := strings.Join(lineStr, "")
	combinedNumb, _ := strconv.Atoi(num)
	return combinedNumb
}

func calcWinningCombos(time, distance int) int {
	totalWins := 0
	for i := 1; i < time; i++ {
		timeRemaining := time - i
		distanceToTravel := i * timeRemaining
		if distanceToTravel > distance {
			totalWins++
		}
	}
	return totalWins
}
