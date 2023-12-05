package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	arr := [][]string{}
	for scanner.Scan() {
		lines := scanner.Text()
		line := strings.Split(lines, "")
		arr = append(arr, line)
	}

	sum := []int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if !isDigit(arr[i][j]) && !isDot(arr[i][j]) {
				if isDigit(arr[i][j-1]) {
					number := walkLeft(arr[i], j-1)
					sum = append(sum, number)
				}
				if isDigit(arr[i][j+1]) {
					number := walkRight(arr[i], j+1)
					sum = append(sum, number)
				}
				if isDigit(arr[i-1][j]) {
					leftPointer := j
					rightPointer := j
					for isDigit(arr[i-1][leftPointer]) && isDigit(arr[i-1][rightPointer]) && leftPointer != 0 && rightPointer != len(arr)-1 {
						leftPointer--
						rightPointer++
						if !isDigit(arr[i-1][leftPointer]) && isDigit(arr[i-1][rightPointer]) {
							for isDigit(arr[i-1][rightPointer]) && rightPointer != len(arr)-1 {
								rightPointer++
							}
						}
						if isDigit(arr[i-1][leftPointer]) && !isDigit(arr[i-1][rightPointer]) {
							for isDigit(arr[i-1][leftPointer]) && leftPointer != 0 {
								leftPointer--
							}
						}
					}
					num := extractNumber(arr, i-1, leftPointer+1, rightPointer-1)
					sum = append(sum, num)
				} else if isDigit(arr[i-1][j-1]) && isDigit(arr[i-1][j+1]) {
					leftNumber := walkLeft(arr[i-1], j-1)
					sum = append(sum, leftNumber)
					rightNumber := walkRight(arr[i-1], j+1)
					sum = append(sum, rightNumber)
				} else if isDigit(arr[i-1][j-1]) {
					num := walkLeft(arr[i-1], j-1)
					sum = append(sum, num)
				} else if isDigit(arr[i-1][j+1]) {
					num := walkRight(arr[i-1], j+1)
					sum = append(sum, num)
				}
				if isDigit(arr[i+1][j]) {
					leftPointer := j
					rightPointer := j
					for isDigit(arr[i+1][leftPointer]) && isDigit(arr[i+1][rightPointer]) && leftPointer != 0 && rightPointer != len(arr)-1 {
						leftPointer--
						rightPointer++
						if !isDigit(arr[i+1][leftPointer]) && isDigit(arr[i+1][rightPointer]) {
							for isDigit(arr[i+1][rightPointer]) && rightPointer != len(arr)-1 {
								rightPointer++
							}
						}
						if isDigit(arr[i+1][leftPointer]) && !isDigit(arr[i+1][rightPointer]) {
							for isDigit(arr[i+1][leftPointer]) && leftPointer != 0 {
								leftPointer--
							}
						}
					}
					num := extractNumber(arr, i+1, leftPointer+1, rightPointer-1)
					sum = append(sum, num)
				} else if isDigit(arr[i+1][j-1]) && isDigit(arr[i+1][j+1]) {
					leftNumber := walkLeft(arr[i+1], j-1)
					sum = append(sum, leftNumber)
					rightNumber := walkRight(arr[i+1], j+1)
					sum = append(sum, rightNumber)
				} else if isDigit(arr[i+1][j-1]) {
					num := walkLeft(arr[i+1], j-1)
					sum = append(sum, num)
				} else if isDigit(arr[i+1][j+1]) {
					num := walkRight(arr[i+1], j+1)
					sum = append(sum, num)
				}
			}
		}
	}
	finalSum := 0
	for _, num := range sum {
		finalSum += num
	}
	fmt.Println(finalSum)
}

func isDigit(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

func isDot(s string) bool {
	if s == "." {
		return true
	}
	return false
}

func walkLeft(arr []string, j int) int {
	numArr := []string{}
	pointer := j
	for isDigit(arr[pointer]) && pointer > 0 {
		numArr = append(numArr, arr[pointer])
		pointer--
		if pointer == 0 && isDigit(arr[pointer]) {
			numArr = append(numArr, arr[pointer])
		}
	}
	foundNumber := reverseAndReturnAsNumber(numArr)
	return foundNumber
}

func walkRight(arr []string, j int) int {
	numArr := []string{}
	pointer := j
	for isDigit(arr[pointer]) && pointer < len(arr)-1 {
		numArr = append(numArr, arr[pointer])
		pointer++
		if pointer == len(arr)-1 && isDigit(arr[pointer]) {
			numArr = append(numArr, arr[pointer])
		}
	}
	foundNumber := combineAndReturnAsNumber(numArr)
	return foundNumber
}

func extractNumber(arr [][]string, row int, leftIdx int, rightIdx int) int {
	numbers := []string{}
	for leftIdx <= rightIdx {
		numbers = append(numbers, arr[row][leftIdx])
		leftIdx++
	}
	num, _ := strconv.Atoi(strings.Join(numbers, ""))
	return num
}

func combineAndReturnAsNumber(arr []string) int {
	num, _ := strconv.Atoi(strings.Join(arr, ""))
	return num
}

func reverseAndReturnAsNumber(arr []string) int {
	p1 := 0
	p2 := len(arr) - 1

	for p2 > p1 {
		temp := arr[p1]
		arr[p1] = arr[p2]
		arr[p2] = temp
		p1++
		p2--
	}
	joined := strings.Join(arr, "")
	num, _ := strconv.Atoi(joined)
	return num
}
