package main

import (
	"advent-of-go-2020/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(6, "\n\n")
	println("Count answers part 2:", solvePart2(input))

}

func solvePart2(input []string) int {
	count := 0
	for _, curr := range input {
		count += answersPart2(curr)
	}

	return count
}

func answersPart2(group string) int{
	givenAnswers := make(map[rune]int)
	lines := strings.Split(group, "\n")

	for _, line := range lines {
		for _, question := range line {
			givenAnswers[question]++
		}
	}

	count := 0
	for question := range givenAnswers {
		if ansCount, present := givenAnswers[question]; present && ansCount == len(lines) {
			count++ // everyone answered yes to question ans
		}
	}

	return count
}
