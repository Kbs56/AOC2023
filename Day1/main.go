package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
		firstNum := findFirstNum(line)
		lastNum := findLastNum(line)
		resStr := firstNum + lastNum
		resNum, err := strconv.Atoi(resStr)
		fmt.Println(resNum)
		if err != nil {
			panic(err)
		}
		sum += resNum
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func findFirstNum(line string) string {
	for _, ch := range line {
		if _, err := strconv.Atoi(string(ch)); err != nil {
		} else {
			return string(ch)
		}
	}
	return ""
}

func findLastNum(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		if _, err := strconv.Atoi(string(line[i])); err != nil {
		} else {
			return string(line[i])
		}
	}
	return ""
}
