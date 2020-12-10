package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"advent-of-go-2020/utils/slices"
)

func main() {
	input := conv.ToIntSlice(files.ReadFile(10, "\n"))
	possible := make(map[int][]int)
	possible[0] = []int{1, 2, 3}

	for _, element := range input {
		possible[element] = []int{element + 3, element + 2, element + 1}
	}

	result := connections(possible, make(map[int]int), slices.Max(input) + 3,0)
	println("Solution part 2:", result)
}

func connections(possible map[int][]int, memo map[int]int, target int, currPos int) int {
	if value, seen := memo[currPos]; seen {
		return value
	}

	value := 0
	for _, current := range possible[currPos] {
		if current != target {
			value += connections(possible, memo, target, current)
			continue
		}

		value += 1
	}

	memo[currPos] = value
	return value
}