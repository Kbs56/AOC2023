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
	fmt.Println("Hello")
	fmt.Println(HIGH_CARD)
}

// Sort the Hands (asc) into []Hand
// Figure out how to settle tiebreakers
// multiply each hands bid by index + 1
