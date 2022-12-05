package main

import (
	"advent-of-go/utils/files"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(4, 2022, "\n")
	part1, part2 := solve(input)
	println(part1)
	println(part2)
}

func solve(input []string) (int, int) {
	fullOverlap, partialOverlap := 0, 0

	for _, line := range input {
		pairs := strings.Split(line, ",")
		pair1, pair2 := strings.Split(pairs[0], "-"), strings.Split(pairs[1], "-")
		a, _ := strconv.Atoi(pair1[0])
		b, _ := strconv.Atoi(pair1[1])
		c, _ := strconv.Atoi(pair2[0])
		d, _ := strconv.Atoi(pair2[1])
		if (a <= c && b >= d) || (c <= a && d >= b) {
			fullOverlap++
		}

		if a <= d && b >= c {
			partialOverlap++
		}
	}

	return fullOverlap, partialOverlap
}
