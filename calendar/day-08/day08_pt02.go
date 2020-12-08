package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"go/types"
	"strings"
)

func main() {
	input := files.ReadFile(8, "\n")
	println("Accumulator if the program halts successfully:", solvePart2(input))
}

func solvePart2(input []string) int {
	changes := make(map[int]types.Nil)

	for {
		acc := 0
		changed := false
		valid := true

		instructions := make(map[int]types.Nil)

		for i := 0; i < len(input); i++ {
			split := strings.Split(input[i], " ")

			if _, has := changes[i]; !changed && split[0] == "jmp" && !has {
				split[0] = "nop"
				changed = true
				changes[i] = types.Nil{}
			} else if !changed && split[0] == "nop" && !has {
				split[0] = "jmp"
				changed = true
				changes[i] = types.Nil{}
			}

			if _, has := instructions[i]; has {
				valid = false
				break
			}

			instructions[i] = types.Nil{}

			delta, pointer := runPart2(split[0], conv.ToInt(split[1]))
			acc += delta
			i += pointer
		}

		if valid {
			return acc
		}
	}
}

func runPart2(in string, number int) (int, int) {
	if in == "acc" {
		return number, 0
	}

	if in == "jmp" {
		return 0, number - 1
	}

	// nop
	return 0, 0
}