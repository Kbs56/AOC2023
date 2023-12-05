package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("flatInput.txt")
	file, err := os.Open("input.txt")
	// file, err := os.Open("testInput.txt")
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
				fmt.Println(arr[i][j])
				if isDigit(arr[i][j-1]) {
					fmt.Println("Left number found", arr[i][j-1])
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
					fmt.Println(foundNumber)
					sum = append(sum, foundNumber)
				}
				if isDigit(arr[i][j+1]) {
					fmt.Println("Right number found", arr[i][j+1])
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
					fmt.Println(foundNumber)
					sum = append(sum, foundNumber)
				}
				if isDigit(arr[i-1][j]) {
					fmt.Println("Up number found", arr[i-1][j])
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
					fmt.Println(num)
					sum = append(sum, num)
				} else if isDigit(arr[i-1][j-1]) && isDigit(arr[i-1][j+1]) {
					fmt.Println("Diag up Buddies!!!", arr[i-1][j-1], arr[i-1][j+1], i-1, j-1, i-1, j+1)
					numArr := []string{}
					pointer := j - 1
					for isDigit(arr[i-1][pointer]) && pointer > 0 {
						fmt.Println("pointer", arr[i-1][pointer])
						numArr = append(numArr, arr[i-1][pointer])
						pointer--
						if pointer == 0 && isDigit(arr[i-1][pointer]) {
							numArr = append(numArr, arr[i-1][pointer])
						}
					}
					foundNumber := reverseAndReturnAsNumber(numArr)
					fmt.Println(foundNumber)
					sum = append(sum, foundNumber)
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
					fmt.Println(foundNumber)
					sum = append(sum, foundNumber)
				} else if isDigit(arr[i-1][j-1]) {
					fmt.Println("Diag up left number found", arr[i-1][j-1])
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
					fmt.Println(foundNumber)
					sum = append(sum, foundNumber)
				} else if isDigit(arr[i-1][j+1]) {
					fmt.Println("Diag up right number found", arr[i-1][j+1])
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
					fmt.Println(foundNumber)
					sum = append(sum, foundNumber)
				}
				if isDigit(arr[i+1][j]) {
					fmt.Println("Down number found", arr[i+1][j])
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
					fmt.Println(num)
					sum = append(sum, num)
				} else if isDigit(arr[i+1][j-1]) && isDigit(arr[i+1][j+1]) {
					fmt.Println("Diag down buddies found!!!", arr[i+1][j-1])
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
					fmt.Println(foundNumber)
					sum = append(sum, foundNumber)
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
					fmt.Println(foundNumber)
					sum = append(sum, foundNumber)
				} else if isDigit(arr[i+1][j-1]) {
					fmt.Println("Diag down left number found", arr[i+1][j-1])
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
					fmt.Println(foundNumber)
					sum = append(sum, foundNumber)
				} else if isDigit(arr[i+1][j+1]) {
					fmt.Println("Diag down right number found", arr[i+1][j+1])
					// just like right found
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
					fmt.Println(foundNumber)
					sum = append(sum, foundNumber)
				}
			}
		}
	}
	fmt.Println(sum)
	sort.Ints(sum)
	filePath := "output.txt"
	outFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(outFile)
	for _, num := range sum {
		str := strconv.Itoa(num)
		_, err := writer.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
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
