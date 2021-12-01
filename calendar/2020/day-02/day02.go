package main

import (
	"advent-of-go/utils/files"
	"strconv"
	"strings"
)

type Rule struct {
	keyLetter string
	minInstance int
	maxInstance int
}
type Input struct {
	rule Rule
	password string
}

func main() {
	input := files.ReadFile(02, 2020, "\n")
	lines := ParseInput(input)
	println(solvePart1(lines))
	println(solvePart2(lines))
}

func ParseInput(input []string) []Input {
	lines := make([]Input, len(input))
	for i, line := range input {
		lines[i] = ParseLine(line)
	}
	return lines
}

func ParseLine(input string) Input {
	parts := strings.Split(input, " ")

	rulesParts := strings.Split(parts[0], "-")
	minInstance, _ := strconv.Atoi(rulesParts[0])
	maxInstance, _ := strconv.Atoi(rulesParts[1])

	return Input {
		rule: Rule {
			keyLetter: parts[1][:1],
			minInstance: minInstance,
			maxInstance: maxInstance,
		},
		password: parts[2],
	}
}

func (input *Input) IsValidForPart1() bool {
	count := strings.Count(input.password, input.rule.keyLetter)
	return count >= input.rule.minInstance && count <= input.rule.maxInstance
}

func (input *Input) IsValidForPart2() bool {
	minCharacter := input.password[input.rule.minInstance - 1:input.rule.minInstance]
	maxCharacter := input.password[input.rule.maxInstance - 1:input.rule.maxInstance]
	return (minCharacter == input.rule.keyLetter || maxCharacter == input.rule.keyLetter) && minCharacter != maxCharacter;
}

func solvePart1(input []Input) int {
	result := 0

	for _, toTest := range input {
		if toTest.IsValidForPart1() {
			result++
		}
	}

	return result
}

func solvePart2(input []Input) int {
	result := 0

	for _, toTest := range input {
		if toTest.IsValidForPart2() {
			result++
		}
	}

	return result
}
