package main

import (
	"advent-of-go-2020/utils/files"
	"fmt"
)

func main() {
	input := files.ReadFile(3, "\n")
	fmt.Printf("Solution part 1: %d\n", numTrees(input, 3, 1))
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
