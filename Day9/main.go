package main

import (
	"fmt"
)

var testArr = []int{10, 13, 16, 21, 30, 45}

// var testArr = []int{1, 3, 6, 10, 15, 21}

func main() {
	// Parse input into arrays for each line
	// Need to do this for each array in the input
	predictionArr := [][]int{}
	patternedArr := getDiff(testArr, len(testArr), predictionArr)

	currentNum := 0
	for i := len(patternedArr) - 1; i > 0; i-- {
		currentNum += patternedArr[i-1][len(patternedArr[i-1])-1]
	}
	// Add current num to a sum
	fmt.Println(currentNum)
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
