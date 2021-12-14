package main

import (
	"advent-of-go/utils/files"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(25, 2015, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	rowTarget, colTarget := parseInput(input)
	currentCode, multiplier, divisor, code := 20151125, 252533, 33554393, 0

	row, col := 1, 1
	for code == 0 {
		// work diagonally
		if row == 1 {
			row = col + 1
			col = 1
		} else {
			row--
			col++
		}

		currentCode = (currentCode * multiplier) % divisor

		if row == rowTarget && col == colTarget {
			code = currentCode
		}
	}
	return code
}

func solvePart2(input []string) int {
	result := 0

	return result
}

func parseInput(input []string) (int, int) {
	parts := strings.Fields(input[0])
	rowStr, colStr := parts[15][:len(parts[15])-1], parts[17][:len(parts[17])-1]
	row, _ := strconv.Atoi(rowStr)
	col, _ := strconv.Atoi(colStr)
	return row, col
}
