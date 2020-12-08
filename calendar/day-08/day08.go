package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"go/types"
	"strings"
)

func main() {
	input := files.ReadFile(8, "\n")
	println("Acc before execution of a line for a second time:", solve(input))
}

func solve(input []string) int {
	acc := 0
	instructions := make(map[int]types.Nil)

	for i := 0; i < len(input); i++ {
		split := strings.Split(input[i], " ")

		if _, has := instructions[i]; has {
			return acc
		}

		instructions[i] = types.Nil{}

		delta, pointer := run(split[0], conv.ToInt(split[1]))
		acc += delta
		i += pointer
	}

	return acc
}

func run(in string, number int) (int, int) {
	if in == "acc" {
		return number, 0
	}

	if in == "jmp" {
		return 0, number - 1
	}

	// nop
	return 0, 0
}
