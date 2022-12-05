package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"strconv"
	"strings"
)

type instruction struct {
	count int
	source int
	destination int
}

func main() {
	input := files.ReadFile(5, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) string {
	return moveCrates(input, false)
}

func solvePart2(input []string) string {
	return moveCrates(input, true)
}

func moveCrates(input []string, moveMultiple bool) string {
	stacks, instructionns := parseInput(input)
	for _, inst := range instructionns {
		toMove := make([]string, inst.count)
		copy(toMove, stacks[inst.source][0:inst.count])
		if !moveMultiple {
			toMove = slices.Reverse(toMove)
		}
		stacks[inst.source] = stacks[inst.source][inst.count:]
		newDest := append(toMove, stacks[inst.destination]...)
		stacks[inst.destination] = newDest
	}

	result := ""
	for i := 1; i <= len(stacks); i++ {
		if len(stacks[i]) > 0 {
			result += stacks[i][0]
		}
	}

	return result
}

func parseInput(input []string) (map[int][]string, []instruction) {
	stacks := map[int][]string{}
	instructionsStart := 0
	for lineNumber, line := range input {
		if line[1:2] == "1" {
			instructionsStart = lineNumber + 2
			break
		}
		for i := 0; i < len(line); i += 4 {
			char := line[i+1:i+2]
			stack := (i / 4) + 1
			if char != " " {
				_, hasStack := stacks[stack]
				if !hasStack {
					stacks[stack] = []string{}
				}
				stacks[stack] = append(stacks[stack], char)
			}
		}
	}
	instructions := []instruction{}
	for i := instructionsStart; i < len(input); i++ {
		parts := strings.Fields(input[i])
		count, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])
		instructions = append(instructions, instruction{ count: count, source: from, destination: to })
	}
	return stacks, instructions
}
