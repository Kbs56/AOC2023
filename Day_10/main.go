package main

import (
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {
	contents, _ := os.ReadFile("testInput.txt")

	lines := strings.Split(string(contents), "\n")

	input := [][]string{}

	for i := 0; i < len(lines)-1; i++ {
		arrLine := strings.Split(lines[i], "")
		input = append(input, arrLine)
	}

	sXPos, sYPos := findStart(input)
	fmt.Println("Starting X Y:", sXPos, sYPos)
	fmt.Println("Starting Char:", input[sXPos][sYPos])

	queue := list.New()

	queue.PushBack(10)
	queue.PushBack(50)

	elem := queue.Front()
	fmt.Println(elem.Value)
}

func findStart(arr [][]string) (int, int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == "S" {
				return i, j
			}
		}
	}
	return 0, 0
}

/*
	| is a vertical pipe connecting north and south.
	- is a horizontal pipe connecting east and west.
	L is a 90-degree bend connecting north and east.
	J is a 90-degree bend connecting north and west.
	7 is a 90-degree bend connecting south and west.
	F is a 90-degree bend connecting south and east.
	. is ground; there is no pipe in this tile.
	S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
*/
