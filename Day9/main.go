package main

import (
	"fmt"
)

// var testArr = []int{0, 3, 6, 9, 12, 15}
// var testArr = []int{1, 3, 6, 10, 15, 21}
var testArr = []int{10, 13, 16, 21, 30, 45}

func main() {
	test := getDiff(testArr, len(testArr))
	fmt.Println(test)
}

func getDiff(arr []int, length int) []int {
	if checkAllZero(arr) {
		return arr
	}
	retVals := []int{}
	for i := 0; i < length-1; i++ {
		diff := arr[i+1] - arr[i]
		retVals = append(retVals, diff)
	}
	return getDiff(retVals, len(retVals))
}

func checkAllZero(arr []int) bool {
	for _, num := range arr {
		if num != 0 {
			return false
		}
	}
	return true
}
