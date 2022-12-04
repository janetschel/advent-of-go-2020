package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(02, 2017, "\n")

	matrix := parseInput(input)
	println(solvePart1(matrix))
	println(solvePart2(matrix))
}

func solvePart1(input [][]int) int {
	checksum := 0

	for _, row := range input {
		checksum += slices.Max(row) - slices.Min(row)
	}

	return checksum
}

func solvePart2(input [][]int) int {
	checksum := 0

	for _, row := range input {
		for _, value := range row {
			for _, candidate := range row {
				if value % candidate == 0 && value != candidate {
					checksum += value / candidate
				}
			}
		}
	}

	return checksum
}

func parseInput(input []string) [][]int {
	rows := make([][]int, len(input))

	for i, row := range input {
		values := strings.Fields(row)
		rows[i] = make([]int, len(values))
		for j, value := range values {
			parsed, _ := strconv.Atoi(value)
			rows[i][j] = parsed
		}
	}
	return rows
}
