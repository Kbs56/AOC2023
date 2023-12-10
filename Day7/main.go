package main

import "fmt"

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
	// Raead file and parse cards line by line
	testMap := make(map[string]int)
	testStr := "ATJKQQ"
	for i := 0; i < len(testStr); i++ {
		val, ok := testMap[string(testStr[i])]
		if ok {
			fmt.Println("yes", val)
			testMap[string(testStr[i])] = val + 1
		} else {
			fmt.Println("no", val)
			testMap[string(testStr[i])] = 1
		}
	}
	fmt.Println(testMap)
}

// Method to get frequency of cards in each hand and assign strength

// Sort the Hands (asc) into []Hand
// Figure out how to settle tiebreakers
// multiply each hands bid by index + 1
