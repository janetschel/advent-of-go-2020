package main

import (
	"advent-of-go/utils/files"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(23, 2015, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	registers := map[string]int{"a": 0, "b": 0}
	return runProgram(input, registers)["b"]
}

func solvePart2(input []string) int {
	registers := map[string]int{"a": 1, "b": 0}
	return runProgram(input, registers)["b"]
}

func runProgram(instructions []string, r map[string]int) map[string]int {
	registers := r
	i := 0
	for i < len(instructions) {
		current := instructions[i]
		parts := strings.Fields(current)

		increment := 1

		switch parts[0] {
		case "hlf":
			registers[parts[1]] /= 2
		case "tpl":
			registers[parts[1]] *= 3
		case "inc":
			registers[parts[1]]++
		case "jmp":
			jump, _ := strconv.Atoi(parts[1])
			increment = jump
		case "jie":
			register := parts[1][:len(parts[1])-1]
			if registers[register]%2 == 0 {
				jump, _ := strconv.Atoi(parts[2])
				increment = jump
			}
		case "jio":
			register := parts[1][:len(parts[1])-1]
			if registers[register] == 1 {
				jump, _ := strconv.Atoi(parts[2])
				increment = jump
			}
		}
		i += increment
	}
	return registers
}
