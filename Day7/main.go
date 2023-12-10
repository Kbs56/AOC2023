package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	HIGH_CARD      int = 0
	ONE_PAIR           = 1
	TWO_PAIR           = 2
	THREE_PAIR         = 3
	FULL_HOUSE         = 4
	FOUR_OF_A_KIND     = 5
	FIVE_OF_A_KIND     = 6
)

type Hand struct {
	cards    []string
	bid      int
	strength int
}

func main() {
	file, _ := os.Open("testInput.txt")

	scanner := bufio.NewScanner(file)

	allHands := []Hand{}

	for scanner.Scan() {
		line := scanner.Text()
		handLine := strings.Split(line, " ")

		hand := strings.Split(handLine[0], "")
		bid, _ := strconv.Atoi(handLine[1])

		handStrength := calcHandStrength(hand)
		currentHand := Hand{cards: hand, bid: bid, strength: handStrength}
		allHands = append(allHands, currentHand)
	}
	fmt.Println(allHands)
}

func calcHandStrength(hand []string) int {
	handMap := make(map[string]int)
	for _, ch := range hand {
		val, ok := handMap[ch]
		if ok {
			handMap[ch] = val + 1
		} else {
			handMap[ch] = 1
		}
	}
	fmt.Println(handMap)
	handStrength := parseHandMap(handMap)
	return handStrength
}

func parseHandMap(handMap map[string]int) int {
	return 0
}

// Sort the Hands (asc) into []Hand
// Figure out how to settle tiebreakers
// multiply each hands bid by index + 1
