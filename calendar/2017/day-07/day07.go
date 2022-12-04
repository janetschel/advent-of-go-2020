package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/sets"
	"advent-of-go/utils/slices"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(07, 2017, "\n")
	part1 := solvePart1(input)
	println(part1)
	println(solvePart2(input, part1))
}

func solvePart1(input []string) string {
	allInput := strings.Join(input, "\n")
	pattern := regexp.MustCompile("[a-z]+")
	matches := pattern.FindAllString(allInput, -1)

	programCounts := make(map[string]int)
	for _, match := range matches {
		programCounts[match]++
	}

	for program, count := range programCounts {
		if count == 1 {
			return program
		}
	}

	return ""
}

func solvePart2(input []string, base string) (int, string) {
	children, weights := parse(input)

	current := base
	target := 0
	
	for !isBalanced(current, children, weights) {
		current, target = findUnbalanced(current, children, weights)
	}
	
	difference := target - getTotalWeight(current, children, weights)

	return difference + weights[current], current
}

func parse(input []string) (map[string][]string, map[string]int) {
	children, weights := make(map[string][]string), make(map[string]int)
	
	discPattern := regexp.MustCompile("[a-z]+")
	weightPattern := regexp.MustCompile("[0-9]+")

	for _, disc := range input {
		discs := discPattern.FindAllString(disc, -1)
		parent, child := discs[0], discs[1:]
		weightStr := weightPattern.FindString(disc)
		weight, _ := strconv.Atoi(weightStr)
		children[parent] = child
		weights[parent] = weight
	}

	return children, weights
}

func isBalanced(disc string, children map[string][]string, weights map[string]int) bool {
	childWeights := sets.New()
	for _, child := range children[disc] {
		childWeights.Add(fmt.Sprint(getTotalWeight(child, children, weights)))
	}
	return childWeights.Size() == 1
}

func findUnbalanced(disc string, children map[string][]string, weights map[string]int) (string, int) {
	childWeightsMap := make(map[string]int)
	for _, child := range children[disc] {
		childWeightsMap[child] = getTotalWeight(child, children, weights)
	}
	childWeights := make([]int, 0, len(childWeightsMap))
	for _, weight := range childWeightsMap {
		childWeights = append(childWeights, weight)
	}
	targetWeight := slices.Mode(childWeights)
	unbalanced := ""
	for child, weight := range childWeightsMap {
		if weight != targetWeight {
			unbalanced = child
		}
	}

	return unbalanced, targetWeight
}

func getTotalWeight(disc string, children map[string][]string, weights map[string]int) int {
	childSum := 0
	for _, child := range children[disc] {
		childSum += getTotalWeight(child, children, weights)
	}
	return childSum + weights[disc]
}
