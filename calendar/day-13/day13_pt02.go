package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"advent-of-go-2020/utils/maths"
	"strings"
)

func main() {
	input := files.ReadFile(13, "\n")
	println("Timestamp t for bus-departures:", solvePart2(input))
}

func solvePart2(input []string) int {
	depTime, minutesApart := parseLineWithIndex(input[1])
	sum, dep := depTime[0] + minutesApart[0], depTime[0]

	for j := 1; j < len(depTime); j++ {
		for (sum + minutesApart[j]) % depTime[j] != 0 {
			sum += dep
		}

		dep = (dep * depTime[j]) / maths.Gcd(dep, depTime[j])
	}

	return sum
}

func parseLineWithIndex(input string) ([]int, []int) {
	result, times := make([]int, 0), make([]int, 0)

	for i, element := range strings.Split(input, ",") {
		if element != "x" {
			result = append(result, conv.ToInt(element))
			times = append(times, i)
		}
	}

	return result, times
}
