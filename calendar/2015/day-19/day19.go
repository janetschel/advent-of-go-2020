package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"strings"
)

type replacement struct {
	toReplace   string
	replaceWith string
}

func main() {
	input := files.ReadFile(19, 2015, "\n\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	startingString := input[1]
	replacements := parseReplacements(strings.Split(input[0], "\n"))

	newStrings := replace([]string{startingString}, replacements)

	return len(newStrings)
}

// do not run for the sample input; it is not deterministic enough
func solvePart2(input []string) int {
	targetString := input[1]
	replacements := parseReversedReplacements(strings.Split(input[0], "\n"))

	current := targetString
	steps := 0
	for current != "e" {
		for rep := range replacements {
			if strings.Contains(current, replacements[rep].toReplace) {
				steps++
				current = strings.Replace(current, replacements[rep].toReplace, replacements[rep].replaceWith, 1)
			}
		}
	}

	return steps
}

func parseReplacements(input []string) []replacement {
	replacements := []replacement{}
	for i := range input {
		parts := strings.Split(input[i], " => ")
		replacements = append(replacements, replacement{toReplace: parts[0], replaceWith: parts[1]})
	}
	return replacements
}

func parseReversedReplacements(input []string) []replacement {
	replacements := []replacement{}
	for i := range input {
		parts := strings.Split(input[i], " => ")
		replacements = append(replacements, replacement{toReplace: parts[1], replaceWith: parts[0]})
	}
	return replacements
}

func replace(input []string, replacements []replacement) []string {
	newStrings := []string{}
	for str := range input {
		for r := range replacements {
			candidates := makeReplacement(input[str], replacements[r])
			for _, candidate := range candidates {
				if !slices.Contains(newStrings, candidate) {
					newStrings = append(newStrings, candidate)
				}
			}
		}
	}
	return newStrings
}

func makeReplacement(input string, r replacement) []string {
	newStrings := []string{}
	for i := range input {
		if strings.Contains(input[i:], r.toReplace) {
			candidate := input[:i] + strings.Replace(input[i:], r.toReplace, r.replaceWith, 1)
			if !slices.Contains(newStrings, candidate) {
				newStrings = append(newStrings, candidate)
			}
		}
	}
	return newStrings
}
