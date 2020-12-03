package main

import (
	"advent-of-go-2020/utils/files"
	"fmt"
)

// For this day this file contains the solution for part 1 and part 2, since they are almost identical
func main() {
	input := files.ReadFile(3, "\n")

	fmt.Printf("Solution part 1: %d\n", solvePart1(input))
	fmt.Printf("Solution part 2: %d", solvePart2(input))
}

func solvePart1(input []string) int {
	return numTrees(input, 3, 1)
}

func solvePart2(input []string) int {
	result := 1

	for i := 1; i <= 7; i += 2 {
		result *= numTrees(input, i, 1)
	}

	result *= numTrees(input, 1, 2)
	return result
}

func numTrees(input []string, deltax int, deltay int) int {
	trees := 0

	for y, x := 0, 0; y < len(input); y, x = y + deltay, x + deltax {
		if string(input[y][x%len(input[y])]) == "#" {
			trees++
		}
	}

	return trees
}
