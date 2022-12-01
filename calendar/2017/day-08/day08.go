package main

import (
	"advent-of-go/utils/files"
	"strconv"
	"strings"
)

type operation string
const (
	increase operation = "inc"
	decrease = "dec"
)

type comparator string
const (
	greaterThan comparator = ">"
	greaterThanOrEqualTo = ">="
	lessThan = "<"
	lessThanOrEqualTo = "<="
	equalTo = "=="
	notEqualTo = "!="
)

type instruction struct {
	destination string
	operation operation
	value int
	operand1 string
	comparator comparator
	compareValue int
}

func main() {
	input := files.ReadFile(8, 2017, "\n")
	part1, part2 := solve(input)
	println(part1)
	println(part2)
}

func solve(input []string) (int, int) {
	registers := make(map[string]int)

	instructions := parse(input)
	maxMemory := 0
	for _, instruction := range instructions {
		modifier := instruction.value
		if instruction.operation == decrease {
			modifier *= -1
		}
		if (instruction.comparator == greaterThan && registers[instruction.operand1] > instruction.compareValue) ||
			(instruction.comparator == greaterThanOrEqualTo && registers[instruction.operand1] >= instruction.compareValue) ||
			(instruction.comparator == lessThan && registers[instruction.operand1] < instruction.compareValue) ||
			(instruction.comparator == lessThanOrEqualTo && registers[instruction.operand1] <= instruction.compareValue) ||
			(instruction.comparator == equalTo && registers[instruction.operand1] == instruction.compareValue) ||
			(instruction.comparator == notEqualTo && registers[instruction.operand1] != instruction.compareValue) {
			registers[instruction.destination] += modifier
			if registers[instruction.destination] > maxMemory {
				maxMemory = registers[instruction.destination]
			}
		}
	}

	maxTerminal := 0
	for _, value := range registers {
		if value > maxTerminal {
			maxTerminal = value
		}
	}

	return maxTerminal, maxMemory
}

func parse(input []string)  []instruction {
	instructions := make([]instruction, len(input))

	for i, inst := range input {
		parts := strings.Fields(inst)
		value, _ := strconv.Atoi(parts[2])
		compareValue, _ := strconv.Atoi(parts[6])
		instruction := instruction{
			destination: parts[0],
			operation: operation(parts[1]),
			value: value,
			operand1: parts[4],
			comparator: comparator(parts[5]),
			compareValue: compareValue,
		}
		instructions[i] = instruction
	}

	return instructions
}
