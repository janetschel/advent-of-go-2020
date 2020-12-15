package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"advent-of-go-2020/utils/maths"
	"strings"
)

func main() {
	input := files.ReadFile(15, "\n")
	split := strings.Split(input[0], ",")

	println("Solution part 1:", solve(split, 2020))
	println("Solution part 2:", solve(split, 30000000)) // does take a while
}

func solve(input []string, max int) int {
	nums, spoken := conv.ToIntSlice(input), make(map[int][]int)

	for i, element := range nums {
		spoken[element] = []int{i + 1}
	}

	lastSpoken := nums[len(nums) - 1]
	for i := len(nums); i < max; i++ {
		curr, total := spoken[lastSpoken], 0

		if len(curr) <= 1 {
			spoken[0] = append(spoken[0], i + 1)
		} else {
			total = maths.Abs(curr[len(curr) - 1] - curr[len(curr) - 2])
			spoken[total] = append(spoken[total], i + 1)
		}

		lastSpoken = total
	}

	return lastSpoken
}
