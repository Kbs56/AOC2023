package main

// https://github.com/golang/go/wiki/SliceTricks

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type seedRange struct {
	destination int
	source      int
	length      int
}

var (
	mappings = [][]seedRange{}
	seeds    = []int{}
)

func main() {
	filePath := "testInput.txt"

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")

	seeds = getSeedNumbers(lines[0])

	for i := 1; i < len(lines); i++ {
		seedRanges := []seedRange{}
		for i < len(lines) && len(lines[i]) > 0 && unicode.IsNumber(rune(lines[i][0])) {
			stringArr := strings.Fields(lines[i])
			nums := makeIntArray(stringArr)
			seedRangeList := seedRange{destination: nums[0], source: nums[1], length: nums[2]}
			fmt.Println(seedRangeList)
			seedRanges = append(seedRanges, seedRangeList)
			i++
		}
		i++
		if len(seedRanges) > 0 {
			mappings = append(mappings, seedRanges)
		}
		fmt.Println()
	}

	fmt.Println(mappings)
}

// Need to do this part now with updated struct bizness
func findNumberInRange(seed int) int {
	for i := 0; i < len(mappings); i++ {
		for j := 0; j < len(mappings[i]); j++ {
			fmt.Println(mappings[i][j])
		}
	}
	fmt.Println()
	return 0
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
