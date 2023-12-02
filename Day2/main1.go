package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("testInput.txt")
	file, err := os.Open("smallInput.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// fmt.Println(line)
		gameID := getGameID(line)
		fmt.Println(gameID)

		cubes := getDiceInGame(line)

		for _, str := range cubes {
			fmt.Println(strings.TrimSpace(str))
		}
	}
}

func getGameID(line string) int {
	gameID := line[5]
	gameIDNum, _ := strconv.Atoi(string(gameID))
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
