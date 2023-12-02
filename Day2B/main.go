package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dice struct {
	Count int
	Color string
}

func main() {
	// file, err := os.Open("input.txt")
	// file, err := os.Open("smallInput.txt")
	file, err := os.Open("testInput.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		rounds := getDiceInGame(line)
		dieMap := checkValidRounds(rounds)
		fmt.Println(dieMap)
		// TODO:
		// multiply all the counts in the map
		// add the product to sum
	}
}

func checkValidRounds(rounds []string) map[string]int {
	diceMap := map[string]int{"red": 0, "green": 0, "blue": 0}
	for _, round := range rounds {
		die := strings.Split(strings.TrimSpace(round), ", ")
		for _, dice := range die {
			count, _ := strconv.Atoi(strings.Split(dice, " ")[0])
			color := strings.Split(dice, " ")[1]
			dice := &Dice{Count: count, Color: color}
			if dice.Count > diceMap[dice.Color] {
				diceMap[dice.Color] = dice.Count
			}
		}
	}
	return diceMap
}

func getDiceInGame(line string) []string {
	dice := []string{}
	_, after, _ := strings.Cut(line, ": ")
	new := strings.Split(after, ";")
	for _, str := range new {
		dice = append(dice, str)
	}
	return dice
}
