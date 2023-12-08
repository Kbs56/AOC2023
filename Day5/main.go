package main

import (
	"fmt"
	"os"
	"slices"
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
	filePath := "input.txt"

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
			seedRanges = append(seedRanges, seedRangeList)
			i++
		}
		i++
		if len(seedRanges) > 0 {
			mappings = append(mappings, seedRanges)
		}
	}

	locations := []int{}
	for _, seed := range seeds {
		seedMapping := findNumberInRange(seed)
		locations = append(locations, seedMapping)
	}

	fmt.Println(slices.Min(locations))
}

func findNumberInRange(seed int) int {
	currentNum := seed
	for i := 0; i < len(mappings); i++ {
		for j := 0; j < len(mappings[i]); j++ {
			dest := mappings[i][j].destination
			source := mappings[i][j].source
			length := mappings[i][j].length
			if source <= currentNum && currentNum < source+length {
				currentNum = dest + (currentNum - source)
				break
			}
		}
	}
	return currentNum
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
