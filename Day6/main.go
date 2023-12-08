package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(contents), "\n")

	times := parseLine(lines[0])
	distances := parseLine(lines[1])

	combos := []int{}

	for i := range times {
		numCombos := calcWinningCombos(times[i], distances[i])
		combos = append(combos, numCombos)
	}

	ans := 1
	for _, num := range combos {
		ans *= num
	}
	fmt.Println(ans)
}

func convertToIntArr(arr []string) []int {
	result := []int{}
	for _, num := range arr {
		number, _ := strconv.Atoi(num)
		result = append(result, number)
	}
	return result
}

func parseLine(line string) []int {
	_, lineAfter, _ := strings.Cut(line, ":")
	lineStr := strings.Fields(lineAfter)
	nums := convertToIntArr(lineStr)
	return nums
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
