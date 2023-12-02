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
		gameID := getGameID(line)
		rounds := getDiceInGame(line)
		isValid := checkValidRounds(rounds)
		if isValid {
			sum += gameID
		}
	}
	fmt.Println(sum)
}

func checkValidRounds(rounds []string) bool {
	diceMap := map[string]int{"red": 12, "green": 13, "blue": 14}
	for _, round := range rounds {
		die := strings.Split(strings.TrimSpace(round), ", ")
		for _, dice := range die {
			count, _ := strconv.Atoi(strings.Split(dice, " ")[0])
			color := strings.Split(dice, " ")[1]
			dice := &Dice{Count: count, Color: color}
			if dice.Count > diceMap[dice.Color] {
				return false
			}
		}
	}
	return true
}

func getGameID(line string) int {
	before, _, _ := strings.Cut(line, ": ")
	gameString := strings.Split(before, " ")[1]
	gameIDNum, _ := strconv.Atoi(string(gameString))
	return gameIDNum
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
