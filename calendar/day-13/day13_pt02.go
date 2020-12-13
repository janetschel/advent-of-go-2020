package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(13, "\n")
	println(solvePart2(input))
}

func solvePart2(input []string) int {
	depTime, minutesApart := parseLine2(input[1])
	sum := depTime[0] + minutesApart[0]
	dep := depTime[0]

	for j := 1; j < len(depTime); j++ {
		println(minutesApart[j], depTime[j])
		for (sum + minutesApart[j]) % depTime[j] != 0 {
			sum += dep
		}
		dep = (dep * depTime[j]) / gcd(dep, depTime[j])
	}

	return sum
}

func parseLine2(input string) ([]int, []int) {
	split := strings.Split(input, ",")
	result := make([]int, 0)
	times := make([]int, 0)

	for i, element := range split {
		if element != "x" {
			result = append(result, conv.ToInt(element))
			times = append(times, i)
		}
	}

	return result, times
}

func gcd(first int, second int) int {
	var div int

	for i := 1; i <= first && i <= second; i++ {
		if first % i==0 && second % i==0 {
			div = i
		}
	}

	return div
}
