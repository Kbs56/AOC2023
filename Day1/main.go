package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("smallInput.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		res := dummyFunc(line)
		fmt.Println(res)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func dummyFunc(line string) string {
	p1 := 0

	return string(line[p1])
}
