package main

import (
	"advent-of-go-2020/utils/files"
	"fmt"
)

func main() {
	input := files.ReadFile(3, "\n")
	result := 1

	// Function to solve included in part 1 (they are identical)
	for i := 1; i <= 7; i += 2 {
		result *= numTrees(input, i, 1)
	}

	result *= numTrees(input, 1, 2)

	fmt.Printf("Solution part 2: %d", result)
}
