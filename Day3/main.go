package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("flatInput.txt")
	// file, err := os.Open("input.txt")
	file, err := os.Open("testInput.txt")
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

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if _, err := strconv.Atoi(arr[i][j]); err != nil && arr[i][j] != "." {
				fmt.Println(arr[i][j])
				// check adjacent values
				// up, down, left, right & diags
			}
		}
	}
}
