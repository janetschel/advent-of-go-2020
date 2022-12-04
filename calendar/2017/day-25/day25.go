package main

import (
	"advent-of-go/utils/files"
	"strconv"
	"strings"
)

type stateInstructions struct {
	write int
	moveDirection int
	nextState string
}

type state map[int]stateInstructions

type blueprint struct {
	startingState string
	steps int
	states map[string]state
}

func main() {
	input := files.ReadFile(25, 2017, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	machineBlueprint := parseProgram(input)
	return runBlueprint(machineBlueprint)
}

func solvePart2(input []string) int {
	result := 0



	return result
}

func runBlueprint(bp blueprint) int {
	tape, cursor, currentState := map[int]int{}, 0, bp.startingState

	for i := 0; i < bp.steps; i++ {
		s := bp.states[currentState]
		inst := s[tape[cursor]]
		tape[cursor] = inst.write
		cursor += inst.moveDirection
		currentState = inst.nextState
	}

	return checksum(tape)
}

func checksum(tape map[int]int) int {
	sum := 0
	for _, value := range tape {
		if value == 1 {
			sum++
		}
	}
	return sum
}

func parseStateInstructions(instructions []string) stateInstructions {
	writeValue := 0
	if strings.Contains(instructions[0], "1") {
		writeValue = 1
	}
	direction := 1
	if strings.Contains(instructions[1], "left") {
		direction = -1
	}
	nextState := getStateName(instructions[2])
	return stateInstructions{
		write: writeValue,
		moveDirection: direction,
		nextState: nextState,
	}
}

func parseState(stateString []string) state {
	return map[int]stateInstructions{
		0: parseStateInstructions(stateString[1:4]),
		1: parseStateInstructions(stateString[5:9]),
	}
}

func parseProgram(input []string) blueprint {
	startingState := input[0][len(input[0]) - 2:len(input[0]) - 1]
	stepsFields := strings.Fields(input[1])
	steps, _ := strconv.Atoi(stepsFields[5])
	states := map[string]state{}

	for i := 3; i < len(input); i += 10 {
		stateName := getStateName(input[i])
		stateValue := parseState(input[i+1:i+9])
		states[stateName] = stateValue
	}

	return blueprint{
		startingState: startingState,
		steps: steps,
		states: states,
	}
}

func getStateName(input string) string {
	return input[len(input) - 2:len(input) - 1]
}
