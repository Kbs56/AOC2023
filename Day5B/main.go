package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type rangeSeed struct {
	start int
	end   int
}

type seq struct {
	start int
	end   int
	diff  int
}

type set = []seq

func main() {
	val := Day5_part2()
	fmt.Println(val)
}

func scanSeedsv2(line string) []rangeSeed {
	numbers := getNumbers(line)
	seeds := make([]rangeSeed, 0)

	i := 0
	for i < len(numbers) {

		seeds = append(seeds, rangeSeed{
			start: numbers[i],
			end:   numbers[i] + numbers[i+1] - 1,
		})

		i = i + 2
	}
	return seeds
}

func scanSeq(line string) seq {
	numbers := getNumbers(line)

	source := numbers[1]
	dest := numbers[0]
	size := numbers[2]

	return seq{
		start: source,
		end:   source + size - 1,
		diff:  dest - source,
	}
}

func getSeed(result int, sets []set) int {
	curr := result

	i := len(sets) - 1

	for i >= 0 {
		var seqSelected seq
		found := false

		j := len(sets[i]) - 1

		for j >= 0 {
			seq := sets[i][j]

			if curr >= seq.start+seq.diff && curr <= seq.end+seq.diff {
				seqSelected = seq
				found = true
			}

			j--
		}

		if found {
			curr = curr - seqSelected.diff
		}

		i--
	}

	return curr
}

func getSets(lines []string) []set {
	sets := make([]set, 0)

	for i, line := range lines {
		if i == 0 || len(line) == 0 {
			continue
		}

		if strings.Contains(line, "map:") {
			sets = append(sets, make([]seq, 0))
			continue
		}

		lastSetIdx := len(sets) - 1

		sets[lastSetIdx] = append(sets[lastSetIdx], scanSeq(line))
	}

	return sets
}

func GetInput(day int, test bool) []string {
	dat, err := os.ReadFile("../Day5/input.txt")
	check(err)

	lines := strings.Split(string(dat), "\n")

	return lines
}

func getNumbers(line string) []int {
	re := regexp.MustCompile(`\d+`)

	f := re.FindAllStringIndex(line, -1)

	numbers := make([]int, 0)

	for _, match := range f {
		valueStr := line[match[0]:match[1]]
		val, err := strconv.Atoi(valueStr)
		check(err)

		numbers = append(numbers, val)
	}

	return numbers
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Day5_part2() int {
	lines := GetInput(5, false)

	seeds := scanSeedsv2(lines[0])
	sets := getSets(lines)

	i := 0
	for {
		seedSolution := getSeed(i, sets)

		for _, seed := range seeds {
			if seedSolution >= seed.start && seedSolution <= seed.end {
				return i
			}
		}

		i++
	}
}
