package main

import (
	"advent-of-go/utils/files"
	"sort"
	"strings"
)

func main() {
	input := files.ReadFile(10, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	score := 0
	for _, line := range input {
		s, _ := scoreLine(line)
		score += s
	}
	return score
}

func solvePart2(input []string) int {
	scores := []int{}
	for _, line := range input {
		_, s := scoreLine(line)
		if s > 0 {
			scores = append(scores, s)
		}
	}
	sort.Ints(scores)
	middle := len(scores) / 2
	return scores[middle]
}

func scoreLine(line string) (int, int) {
	stack := []rune{}

	open := "([{<"
	close := ")]}>"

	for _, char := range line {
		if strings.ContainsRune(open, char) {
			stack = append(stack, char)
		} else if strings.ContainsRune(close, char) {
			top := stack[len(stack)-1]
			if char == []rune(close)[strings.IndexRune(open, top)] {
				stack = stack[:len(stack)-1]
			} else {
				return corruptedRuneScores[char], 0
			}
		}
	}

	score := 0
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		score *= 5
		score += missingRuneScores[[]rune(close)[strings.IndexRune(open, top)]]
	}
	return 0, score
}

var corruptedRuneScores = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var missingRuneScores = map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
