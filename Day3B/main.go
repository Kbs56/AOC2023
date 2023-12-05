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

	newSum := []int{}
	symbolNumbers := []int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if !isDigit(arr[i][j]) && !isDot(arr[i][j]) {
				if isDigit(arr[i][j-1]) {
					numArr := []string{}
					pointer := j - 1
					for isDigit(arr[i][pointer]) && pointer > 0 {
						numArr = append(numArr, arr[i][pointer])
						pointer--
						if pointer == 0 && isDigit(arr[i][pointer]) {
							numArr = append(numArr, arr[i][pointer])
						}
					}
					foundNumber := reverseAndReturnAsNumber(numArr)
					symbolNumbers = append(symbolNumbers, foundNumber)
				}
				if isDigit(arr[i][j+1]) {
					numArr := []string{}
					pointer := j + 1
					for isDigit(arr[i][pointer]) && pointer < len(arr)-1 {
						numArr = append(numArr, arr[i][pointer])
						pointer++
						if pointer == len(arr)-1 && isDigit(arr[i][pointer]) {
							numArr = append(numArr, arr[i][pointer])
						}
					}
					foundNumber := combineAndReturnAsNumber(numArr)
					symbolNumbers = append(symbolNumbers, foundNumber)
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
					symbolNumbers = append(symbolNumbers, num)
				} else if isDigit(arr[i-1][j-1]) && isDigit(arr[i-1][j+1]) {
					numArr := []string{}
					pointer := j - 1
					for isDigit(arr[i-1][pointer]) && pointer > 0 {
						numArr = append(numArr, arr[i-1][pointer])
						pointer--
						if pointer == 0 && isDigit(arr[i-1][pointer]) {
							numArr = append(numArr, arr[i-1][pointer])
						}
					}
					foundNumber := reverseAndReturnAsNumber(numArr)
					symbolNumbers = append(symbolNumbers, foundNumber)
					numArr = []string{}
					pointer = j + 1
					for isDigit(arr[i-1][pointer]) && pointer < len(arr)-1 {
						numArr = append(numArr, arr[i-1][pointer])
						pointer++
						if pointer == len(arr)-1 && isDigit(arr[i-1][pointer]) {
							numArr = append(numArr, arr[i-1][pointer])
						}
					}
					foundNumber = combineAndReturnAsNumber(numArr)
					symbolNumbers = append(symbolNumbers, foundNumber)
				} else if isDigit(arr[i-1][j-1]) {
					numArr := []string{}
					pointer := j - 1
					for isDigit(arr[i-1][pointer]) && pointer > 0 {
						numArr = append(numArr, arr[i-1][pointer])
						pointer--
						if pointer == 0 && isDigit(arr[i-1][pointer]) {
							numArr = append(numArr, arr[i-1][pointer])
						}
					}
					foundNumber := reverseAndReturnAsNumber(numArr)
					symbolNumbers = append(symbolNumbers, foundNumber)
				} else if isDigit(arr[i-1][j+1]) {
					numArr := []string{}
					pointer := j + 1
					for isDigit(arr[i-1][pointer]) && pointer < len(arr)-1 {
						numArr = append(numArr, arr[i-1][pointer])
						pointer++
						if pointer == len(arr)-1 && isDigit(arr[i-1][pointer]) {
							numArr = append(numArr, arr[i-1][pointer])
						}
					}
					foundNumber := combineAndReturnAsNumber(numArr)
					symbolNumbers = append(symbolNumbers, foundNumber)
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
					symbolNumbers = append(symbolNumbers, num)
				} else if isDigit(arr[i+1][j-1]) && isDigit(arr[i+1][j+1]) {
					numArr := []string{}
					pointer := j - 1
					for isDigit(arr[i+1][pointer]) && pointer > 0 {
						numArr = append(numArr, arr[i+1][pointer])
						pointer--
						if pointer == 0 && isDigit(arr[i+1][pointer]) {
							numArr = append(numArr, arr[i+1][pointer])
						}
					}
					foundNumber := reverseAndReturnAsNumber(numArr)
					symbolNumbers = append(symbolNumbers, foundNumber)
					numArr = []string{}
					pointer = j + 1
					for isDigit(arr[i+1][pointer]) && pointer < len(arr)-1 {
						numArr = append(numArr, arr[i+1][pointer])
						pointer++
						if pointer == len(arr)-1 && isDigit(arr[i+1][pointer]) {
							numArr = append(numArr, arr[i+1][pointer])
						}
					}
					foundNumber = combineAndReturnAsNumber(numArr)
					symbolNumbers = append(symbolNumbers, foundNumber)
				} else if isDigit(arr[i+1][j-1]) {
					numArr := []string{}
					pointer := j - 1
					for isDigit(arr[i+1][pointer]) && pointer > 0 {
						numArr = append(numArr, arr[i+1][pointer])
						pointer--
						if pointer == 0 && isDigit(arr[i+1][pointer]) {
							numArr = append(numArr, arr[i+1][pointer])
						}
					}
					foundNumber := reverseAndReturnAsNumber(numArr)
					symbolNumbers = append(symbolNumbers, foundNumber)
				} else if isDigit(arr[i+1][j+1]) {
					numArr := []string{}
					pointer := j + 1
					for isDigit(arr[i+1][pointer]) && pointer < len(arr)-1 {
						numArr = append(numArr, arr[i+1][pointer])
						pointer++
						if pointer == len(arr)-1 && isDigit(arr[i+1][pointer]) {
							numArr = append(numArr, arr[i+1][pointer])
						}
					}
					foundNumber := combineAndReturnAsNumber(numArr)
					symbolNumbers = append(symbolNumbers, foundNumber)
				}
			}
			if len(symbolNumbers) == 2 {
				gearRatio := symbolNumbers[0] * symbolNumbers[1]
				newSum = append(newSum, gearRatio)
			}
			symbolNumbers = nil
		}
	}
	finalSum := 0
	for _, num := range newSum {
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
