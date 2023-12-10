package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	HIGH_CARD       int = 0
	ONE_PAIR            = 1
	TWO_PAIR            = 2
	THREE_OF_A_KIND     = 3
	FULL_HOUSE          = 4
	FOUR_OF_A_KIND      = 5
	FIVE_OF_A_KIND      = 6
)

type Hand struct {
	cards    []string
	bid      int
	strength int
}

var cards = "AKQJT98765432"

func main() {
	file, _ := os.Open("input.txt")

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
		sortHands(allHands)
	}

	sum := 0
	for i := 0; i < len(allHands); i++ {
		sum += allHands[i].bid * (i + 1)
	}
	fmt.Println(sum)
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
	handStrength := parseHandMap(handMap)
	return handStrength
}

func parseHandMap(handMap map[string]int) int {
	switch len(handMap) {
	case 1:
		return FIVE_OF_A_KIND
	case 2:
		for _, v := range handMap {
			if v == 3 {
				return FULL_HOUSE
			}
		}
		return FOUR_OF_A_KIND
	case 3:
		for _, v := range handMap {
			if v == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	case 4:
		return ONE_PAIR
	case 5:
		return HIGH_CARD
	default:
		return 0
	}
}

func sortHands(allHands []Hand) {
	sort.Slice(allHands, func(i, j int) bool {
		if allHands[i].strength != allHands[j].strength {
			return allHands[i].strength < allHands[j].strength
		}
		for idx := range allHands[i].cards {
			if allHands[i].cards[idx] != allHands[j].cards[idx] {
				return strings.Index(
					cards,
					allHands[i].cards[idx],
				) > strings.Index(
					cards,
					allHands[j].cards[idx],
				)
			} else {
				continue
			}
		}
		// Will never be reached due to input
		return allHands[i].strength < allHands[j].strength
	})
}
