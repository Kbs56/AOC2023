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
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		rounds := getDiceInGame(line)
		dieMap := checkValidRounds(rounds)
		product := multiplyMapValues(dieMap)
		sum += product
	}
	fmt.Println(sum)
}

func multiplyMapValues(dieMap map[string]int) int {
	product := 1
	for key := range dieMap {
		product *= dieMap[key]
	}
	return product
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
	gameArray := []string{}
	_, after, _ := strings.Cut(line, ": ")
	roundsArray := strings.Split(after, ";")
	for _, round := range roundsArray {
		gameArray = append(gameArray, round)
	}
	return gameArray
}
