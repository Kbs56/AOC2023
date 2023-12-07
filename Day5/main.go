package main

// https://github.com/golang/go/wiki/SliceTricks

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	filePath := "testInput.txt"

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")

	seeds := getSeedNumbers(lines[0])
	fmt.Println(seeds)

	// something to hold list of mappings here
	mappings := [][]int{}

	for i := 2; i < len(lines); i++ {
		for i < len(lines) && len(lines[i]) > 0 && unicode.IsNumber(rune(lines[i][0])) {
			// fmt.Println(lines[i])
			stringArr := strings.Fields(lines[i])
			nums := makeIntArray(stringArr)
			fmt.Println(nums)
			mappings = append(mappings, nums)
			i++
		}
	}
	fmt.Println(mappings)
}

func makeIntArray(arr []string) []int {
	result := []int{}
	for _, num := range arr {
		number, _ := strconv.Atoi(num)
		result = append(result, number)
	}
	return result
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
