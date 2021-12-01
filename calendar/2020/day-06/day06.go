package main

import (
	"advent-of-go/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(06, 2020, "\n\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(groups []string) int {
	result := 0

	for i := range groups {
		distinct := make(map[byte]bool)
		group := strings.ReplaceAll(groups[i], "\n", "")
		for j := range group {
			distinct[group[j]] = true
		}
		result += len(distinct)
	}

	return result
}

func solvePart2(groups []string) int {
	result := 0

	for i := range groups {
		distinct := make(map[string]int)
		group := groups[i]
		for _, person := range strings.Split(group, "\n") {
			for j := 0; j < len(person); j++ {
				distinct[person[j:j+1]]++
			}
		}
		for _, value := range distinct {
			if (value == strings.Count(group, "\n") + 1) {
				result++
			}
		}
	}

	return result
}
