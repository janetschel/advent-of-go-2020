package main


/*
import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"advent-of-go-2020/utils/slices"
)

func main() {
	input := files.ReadFile(10, "\n")
	println(solvePart1(conv.ToIntSlice(input)))
}

func solvePart1(input []int) int {
	input = append(input, slices.Max(input) + 3)
	current, oneSteps, threeSteps, low := 0, 0, 0, 0

	for current != slices.Max(input) {
		current, low = current + low, 0
		for _, element := range input {
			if element == current + 1 {
				low = 1
			} else if element == current + 2 && low != 1 {
				low = 2
			} else if element == current + 3 && low == 0 {
				low = 3
			}
		}

		if low == 1 {
			oneSteps++
		} else if low == 3 {
			threeSteps++
		}
	}

	return oneSteps * threeSteps
}
*/
