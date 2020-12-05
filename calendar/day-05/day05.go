package main

import (
	"advent-of-go-2020/utils/files"
	"fmt"
)

func main() {
	input, hi := files.ReadFile(5, "\n"), 0

	for _, pattern := range input {
		curr := solve(pattern)
		if curr > hi {
			hi = curr
		}
	}

	fmt.Printf("Highest seat id: %d", hi)
}

func solve(pattern string) int {
	var row, col int
	binarySearch(0, 127, pattern[:7], "B", &row)
	binarySearch(0, 7, pattern[7:], "R", &col)

	return row * 8 + col
}

func binarySearch(lo int, hi int, rows string, char string, row *int) {
	for _, curr := range rows {
		if string(curr) == char {
			lo, *row = lo + ((hi - lo) / 2 + 1), hi
		} else {
			hi, *row = hi - ((hi - lo) / 2 + 1), lo
		}
	}
}
