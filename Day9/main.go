package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		data := strings.Fields(text)
		numberArr := []int{}
		for _, num := range data {
			convertedNumber, _ := strconv.Atoi(num)
			numberArr = append(numberArr, convertedNumber)
		}
		// Parse input into arrays for each line
		// Need to do this for each array in the input
		predictionArr := [][]int{}
		patternedArr := getDiff(numberArr, len(numberArr), predictionArr)

		currentNum := 0
		for i := len(patternedArr) - 1; i > 0; i-- {
			currentNum += patternedArr[i-1][len(patternedArr[i-1])-1]
		}
		// Add current num to a sum
		fmt.Println(currentNum)
		sum += currentNum
	}
	fmt.Println(sum)
}

func getDiff(arr []int, length int, predictionArr [][]int) [][]int {
	predictionArr = append(predictionArr, arr)
	if checkAllZero(arr) {
		return predictionArr
	}
	retVals := []int{}
	for i := 0; i < length-1; i++ {
		diff := arr[i+1] - arr[i]
		retVals = append(retVals, diff)
	}
	return getDiff(retVals, len(retVals), predictionArr)
}

func checkAllZero(arr []int) bool {
	for _, num := range arr {
		if num != 0 {
			return false
		}
	}
	return true
}
