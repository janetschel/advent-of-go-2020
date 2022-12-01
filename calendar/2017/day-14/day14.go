package main

import (
	"advent-of-go/utils/files"
	knothash "advent-of-go/utils/knotHash"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(14, 2017, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	usedCount := 0

	hashInputs := buildHashInputs(input)
	for _, input := range hashInputs {
		hash := knothash.HashHex(input)
		binary := hashToBinary(hash)
		usedCount += strings.Count(binary, "1")
	}


	return usedCount
}

func solvePart2(input []string) int {
	grid := []string{}
	hashInputs := buildHashInputs(input)
	for _, input := range hashInputs {
		hash := knothash.HashHex(input)
		grid = append(grid, hashToBinary(hash))
	}

	regions := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				visit(i, j, grid)
				regions++
			}
		}
	}

	return regions
}

func visit(i, j int, grid []string) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}

	str := []rune(grid[i])
	str[j] = '0'
	grid[i] = string(str)

	visit(i+1, j, grid)
	visit(i-1, j, grid)
	visit(i, j+1, grid)
	visit(i, j-1, grid)
}

func hashToBinary(hashInput string) string {
	hash := ""

	for _, char := range hashInput {
		hex, _ := strconv.ParseInt(string(char), 16, 64)
		hash += fmt.Sprintf("%04b", hex)
	}
	return hash
}

func buildHashInputs(input []string) []string {
	hashInputs := make([]string, 128)

	keyString := input[0]
	for i := 0; i < 128; i++ {
		hashInputs[i] = fmt.Sprintf("%s-%d", keyString, i)
	}

	return hashInputs
}
