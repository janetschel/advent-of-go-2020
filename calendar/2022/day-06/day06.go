package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/sets"
	"strings"
)

func main() {
	input := files.ReadFile(6, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	return findMarker(input[0], 4)
}

func solvePart2(input []string) int {
	return findMarker(input[0], 14)
}

func findMarker(buffer string, markerSize int) int {
	for i := markerSize; i < len(buffer); i++ {
		set := sets.New()
		set.AddRange(strings.Split(buffer[i-markerSize:i], ""))
		if set.Size() == markerSize {
			return i
		}
	}

	return -1
}
