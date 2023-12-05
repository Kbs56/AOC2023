package main

import (
	"fmt"
	"slices"
)

func main() {
	sum := 0
	containsWinningNumbers := false
	// winningNumbers := []int{41, 48, 83, 86, 17}
	// myNumbers := []int{83, 86, 6, 31, 17, 9, 48, 53}
	winningNumbers := []int{13, 32, 20, 16, 61}
	myNumbers := []int{61, 30, 68, 82, 17, 32, 24, 19}

	fmt.Println(winningNumbers)
	fmt.Println(myNumbers)

	matchedNumbers := []int{}
	for _, num := range myNumbers {
		if slices.Contains(winningNumbers, num) {
			matchedNumbers = append(matchedNumbers, num)
		}
	}
	fmt.Println(matchedNumbers)

	score := 1
	if len(matchedNumbers) > 0 {
		containsWinningNumbers = true
		for i := 0; i < len(matchedNumbers)-1; i++ {
			score *= 2
		}
	}

	if containsWinningNumbers {
		sum += score
		containsWinningNumbers = false
	}
	fmt.Println(sum)
}
