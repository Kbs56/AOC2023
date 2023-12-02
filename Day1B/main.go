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
	// file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		// find first and last num, as well as their index
		// search before the firstNum index and after the lastNum index for a match to "one", "two", "three", etc...
		// if there is a match get the corresponding mapping
		// return first number/string number combined with last number/string number
		line := scanner.Text()
		fmt.Println(line)
		firstStr, firstIdx := findFirstNum(line)
		lastStr, lastIdx := findLastNum(line)
		// fmt.Println(firstStr, firstIdx)
		// fmt.Println(lastStr, lastIdx)

		// What if the index does not exist?
		// first / lastStr will be empty
		// Search for all occurences and get the first and last?

		if len(firstStr) == 0 || len(lastStr) == 0 {
			arr := fullRegexSearch(line)
			firstNumStr := arr[0]
			lastNumStr := arr[len(arr)-1]
			firstNumStr, _ = mapStringtoNum(firstNumStr)
			lastNumStr, _ = mapStringtoNum(lastNumStr)
			result, _ := strconv.Atoi(firstNumStr + lastNumStr)
			sum += result
			fmt.Println(result)
			fmt.Println(sum)
		} else {
			firstNumStr := regexSearchFront(firstIdx, line)
			lastNumStr := regexSearchBack(lastIdx, line)
			// fmt.Println(firstNumStr)
			// fmt.Println(lastNumStr)

			// Need to check here for empty string returned from regex search method
			if len(firstNumStr) == 0 && len(lastNumStr) != 0 {
				result, _ := strconv.Atoi(firstStr + lastNumStr)
				sum += result
				fmt.Println(result)
				fmt.Println(sum)
			} else if len(firstNumStr) != 0 && len(lastNumStr) == 0 {
				result, _ := strconv.Atoi(firstNumStr + lastStr)
				sum += result
				fmt.Println(result)
				fmt.Println(sum)
			} else if len(firstNumStr) == 0 && len(lastNumStr) == 0 {
				result, _ := strconv.Atoi(firstStr + lastStr)
				sum += result
				fmt.Println(result)
				fmt.Println(sum)
			} else {
				result, _ := strconv.Atoi(firstNumStr + lastNumStr)
				sum += result
				fmt.Println(result)
				fmt.Println(sum)
			}
		}
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

func fullRegexSearch(line string) []string {
	pattern := "one|two|three|four|five|six|seven|eight|nine"
	r := regexp.MustCompile(pattern)

	res := r.FindAllString(line, -1)
	return res
}

func regexSearchFront(idx int, line string) string {
	pattern := "one|two|three|four|five|six|seven|eight|nine"
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
	pattern := "one|two|three|four|five|six|seven|eight|nine"
	r, _ := regexp.Compile(pattern)
	searchStr := line[(idx + 1):]
	numStr := r.FindAllString(searchStr, -1)

	if len(numStr) != 0 {
		num, _ := mapStringtoNum(numStr[len(numStr)-1])
		return num
	}

	return ""
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
