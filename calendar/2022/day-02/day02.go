package main

import (
	"advent-of-go/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(2, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	totalScore := 0

	for _, round := range input {
		choices := strings.Fields(round)
		totalScore += scoreRound(choices[0], choices[1])
	}

	return totalScore
}

func solvePart2(input []string) int {
	totalScore := 0

	for _, round := range input {
		choices := strings.Fields(round)
		var yourChoice string
		if choices[1] == "X" {
			// lose
			if choices[0] == "A" {
				yourChoice = "Z"
			} else if choices[0] == "B" {
				yourChoice = "X"
			} else {
				yourChoice = "Y"
			}
		} else if choices[1] == "Y" {
			// draw
			if choices[0] == "A" {
				yourChoice = "X"
			} else if choices[0] == "B" {
				yourChoice = "Y"
			} else {
				yourChoice = "Z"
			}
		} else {
			// win
			if choices[0] == "A" {
				yourChoice = "Y"
			} else if choices[0] == "B" {
				yourChoice = "Z"
			} else {
				yourChoice = "X"
			}
		}
		totalScore += scoreRound(choices[0], yourChoice)
	}

	return totalScore
}

func scoreRound(opponentsChoice string, yourChoice string) int {
	choiceMap := map[string]int{ "X": 1, "Y": 2, "Z": 3}
	score := choiceMap[yourChoice]

	if opponentsChoice == "A" && yourChoice == "X" ||
	opponentsChoice == "B" && yourChoice == "Y" ||
	opponentsChoice == "C" && yourChoice == "Z" {
		return score + 3
	}

	if opponentsChoice == "A" && yourChoice == "Z" ||
		opponentsChoice == "B" && yourChoice == "X" ||
		opponentsChoice == "C" && yourChoice == "Y" {
			return score
	}

	return score + 6
}
