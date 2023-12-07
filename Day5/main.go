package main

// https://github.com/golang/go/wiki/SliceTricks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("testInput.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	seeds := getSeedNumbers(line)
	fmt.Println(seeds)

	// Scan the rest of the lines
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			if unicode.IsNumber(rune(line[0])) {
				fmt.Println(line)
			}
		}
	}
}

func getSeedNumbers(line string) []int {
	result := []int{}
	_, after, _ := strings.Cut(line, ": ")
	stringArr := strings.Fields(after)
	for _, num := range stringArr {
		number, _ := strconv.Atoi(num)
		result = append(result, number)
	}
	return result
}
