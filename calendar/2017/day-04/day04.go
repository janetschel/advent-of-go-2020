package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/sets"
	"sort"
	"strings"
)

func main() {
	input := files.ReadFile(04, 2017, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	valid := 0

	for _, password := range input {
		if isValid(password) {
			valid++
		}
	}

	return valid
}

func solvePart2(input []string) int {
	valid := 0

	for _, password := range input {
		words := strings.Fields(password)
		sortedWords := ""
		for _, word := range words {
			chars := strings.Split(word, "")
			sort.Strings(chars)
			sortedWords += " " + strings.Join(chars, "")
		}
		if isValid(sortedWords) {
			valid++
		}
	}

	return valid
}

func isValid(password string) bool {
	words := strings.Fields(password)
	uniqueWords := sets.New()
	uniqueWords.AddRange(words)
	return len(words) == uniqueWords.Size()
}
