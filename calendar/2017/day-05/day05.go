package main

import (
	"advent-of-go/utils/files"
	"strconv"
)

func main() {
	input := files.ReadFile(05, 2017, "\n")

	instructions := make([]int, len(input))
	for i, ins := range input {
		value, _ := strconv.Atoi(ins)
		instructions[i] = value
	}
	println(executeInstructions(instructions, incrementPart1))
	println(executeInstructions(instructions, incrementPart2))
}

func incrementPart1(offset int) int {
	return 1
}

func incrementPart2(offset int) int {
	if offset >= 3 {
		return -1
	}
	return 1
}

func executeInstructions(input []int, increment func(int) int) int {
	instructions := make([]int, len(input))
	copy(instructions, input)
	jumps := 0

	for i := 0; i >= 0 && i <= len(instructions) - 1; {
		instruction := instructions[i]
		instructions[i] += increment(instruction)
		i += instruction
		jumps++
	}

	return jumps
}
