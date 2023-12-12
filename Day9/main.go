package main

import (
	"fmt"
)

var testArr = []int{10, 13, 16, 21, 30, 45}

func main() {
	testRet := [][]int{}
	test := getDiff(testArr, len(testArr), testRet)

	currentNum := 0
	for i := len(test) - 1; i > 0; i-- {
		currentNum += test[i-1][len(test[i-1])-1]
	}
	fmt.Println(currentNum)
}

func getDiff(arr []int, length int, testRet [][]int) [][]int {
	testRet = append(testRet, arr)
	if checkAllZero(arr) {
		return testRet
	}
	retVals := []int{}
	for i := 0; i < length-1; i++ {
		diff := arr[i+1] - arr[i]
		retVals = append(retVals, diff)
	}
	return getDiff(retVals, len(retVals), testRet)
}

func checkAllZero(arr []int) bool {
	for _, num := range arr {
		if num != 0 {
			return false
		}
	}
	return true
}
