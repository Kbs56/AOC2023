package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		rawLine := scanner.Text()
		_, lineAfter, _ := strings.Cut(rawLine, ": ")
		lineLeft, lineRight, _ := strings.Cut(lineAfter, " | ")
		winningNumbers := strings.Fields(lineLeft)
		myNumbers := strings.Fields(lineRight)
		cardScore := processCardScore(winningNumbers, myNumbers)
		sum += cardScore
	}
	fmt.Println(sum)
}

func processCardScore(winningNumbers, myNumbers []string) int {
	cardSum := 0
	containsWinningNumbers := false

	matchedNumbers := []string{}
	for _, num := range myNumbers {
		if slices.Contains(winningNumbers, num) {
			matchedNumbers = append(matchedNumbers, num)
		}
	}

	score := 1
	if len(matchedNumbers) > 0 {
		containsWinningNumbers = true
		for i := 0; i < len(matchedNumbers)-1; i++ {
			score *= 2
		}
	}

	if containsWinningNumbers {
		cardSum += score
		return cardSum
	}
	return 0
}
