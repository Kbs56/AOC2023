package main

import (
	"fmt"
)

func main() {
	num := 52
	rnge := 48
	for i := num; i < num+rnge; i++ {
		fmt.Print(i, " ")
	}
}

// Produce map:
// function that takes in two range start numbers and a range
// put all of the numbers for each range into arrays
// map each index to each other
// How will we handle overlapping indexes if any?

/*
Destination range start | Source Range Start | Range

seed-to-soil map:
50 98 2
52 50 48

Seed -> Soil (line 15)
98, 99 (Seeds)
50, 51 (Soil)

98:50
99:51

Seed -> Soil (line 16)
50, 51, 52, 53... 97 (Seeds)
52, 53, 54, 55... 99 (Soil)

50:52
51:53
52:54
53:55
97:99
*/
