package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// file, err := os.Open("smallInput.txt")
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// find first and last num, as well as their index
		// search before the firstNum index and after the lastNum index for a match to "one", "two", "three", etc...
		// if there is a match get the corresponding mapping
		// return first number/string number combined with last number/string number
		line := scanner.Text()
		firstStr, firstIdx := findFirstNum(line)
		lastStr, lastIdx := findFirstNum(line)
		fmt.Println(firstStr, firstIdx)
		fmt.Println(lastStr, lastIdx)

		firstNumStr := regexSearchFront(firstIdx, line)
		lastNumStr := regexSearchBack(lastIdx, line)
		fmt.Println(firstNumStr)
		fmt.Println(lastNumStr)

		result, _ := strconv.Atoi(firstNumStr + lastNumStr)
		fmt.Println(result)
	}
}

func findFirstNum(line string) (string, int) {
	for i, ch := range line {
		if _, err := strconv.Atoi(string(ch)); err != nil {
		} else {
			return string(ch), i
		}
	}
	return "", 0
}

func findLastNum(line string) (string, int) {
	for i := len(line) - 1; i >= 0; i-- {
		if _, err := strconv.Atoi(string(line[i])); err != nil {
		} else {
			return string(line[i]), i
		}
	}
	return "", 0
}

func regexSearchFront(idx int, line string) string {
	pattern := `\b(one|two|three|four|five|six|seven|eight|nine)\b`
	r, _ := regexp.Compile(pattern)
	searchStr := line[0:idx]
	numStr := r.FindString(searchStr)

	if len(numStr) != 0 {
		num, _ := mapStringtoNum(numStr)
		return num
	}

	return numStr
}

func regexSearchBack(idx int, line string) string {
	pattern := `\b(one|two|three|four|five|six|seven|eight|nine)\b`
	r, _ := regexp.Compile(pattern)
	searchStr := line[(idx + 1):]
	numStr := r.FindString(searchStr)

	if len(numStr) != 0 {
		num, _ := mapStringtoNum(numStr)
		return num
	}

	return numStr
}

func mapStringtoNum(str string) (string, error) {
	mappings := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	result, exists := mappings[str]

	if exists {
		return result, nil
	}

	return "", fmt.Errorf("No mapping")
}
