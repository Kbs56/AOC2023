package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var cardMap = make(map[int]int)

func main() {
	file, err := os.Open("../Day4/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	cardMap[1] = 1

	row := 1
	sum := 0
	for scanner.Scan() {
		rawLine := scanner.Text()
		_, lineAfter, _ := strings.Cut(rawLine, ": ")
		lineLeft, lineRight, _ := strings.Cut(lineAfter, " | ")
		winningNumbers := strings.Fields(lineLeft)
		myNumbers := strings.Fields(lineRight)
		numberOfMatches := processCard(winningNumbers, myNumbers)
		cardMap = checkMapForRow(row)
		if numberOfMatches > 0 {
			cardMap = updateMap(numberOfMatches, row, cardMap)
		} else {
			cardMap = checkMapForRow(row)
		}
		row++
	}

	for _, val := range cardMap {
		sum += val
	}
	fmt.Println(sum)
}

func checkMapForRow(row int) map[int]int {
	_, ok := cardMap[row]
	if ok {
	} else {
		cardMap[row] = 1
	}
	return cardMap
}

func updateMap(matches, row int, cardMap map[int]int) map[int]int {
	for i := row; i < matches+row; i++ {
		val, ok := cardMap[i+1]
		if ok {
			cardMap[i+1] = val + cardMap[row]
		} else {
			cardMap[i+1] = 1 + cardMap[row]
		}
	}
	return cardMap
}

func processCard(winningNumbers, myNumbers []string) int {
	matchedNumbers := []string{}
	for _, num := range myNumbers {
		if slices.Contains(winningNumbers, num) {
			matchedNumbers = append(matchedNumbers, num)
		}
	}

	if len(matchedNumbers) > 0 {
		return len(matchedNumbers)
	}
	return 0
}
