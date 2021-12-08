package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"strconv"
	"strings"
)

type HappinessPair struct {
	primaryPerson string
	secondaryPerson string
	happinessDelta int
}

func main() {
	input := files.ReadFile(13, 2015, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	pairs := make([]HappinessPair, len(input))
	for i := range input {
		pairs[i] = parseLine(input[i])
	}
	matrix := buildHappinessMatrix(pairs)

	return calculateMaxHappiness(matrix)
}

func solvePart2(input []string) int {
	pairs := make([]HappinessPair, len(input))
	for i := range input {
		pairs[i] = parseLine(input[i])
	}
	matrix := buildHappinessMatrix(pairs)
	matrix = append(matrix, make([]int, len(matrix)))

	for i := range matrix {
		matrix[i] = append(matrix[i], 0)
	}

	return calculateMaxHappiness(matrix)
}

func calculateMaxHappiness(matrix [][]int) int {
	indexes := make([]int, len(matrix))
	for i := range indexes {
		indexes[i] = i
	}
	permutations := slices.GeneratePermutations(indexes)
	happinesses := make([]int, len(permutations))
	for i := range permutations {
		h := calculateNetHappiness(matrix, permutations[i])
		happinesses[i] = h
	}

	return slices.Max(happinesses)
}

func buildReferenceMap(pairs []HappinessPair) map[string]int {
	peopleMap := make(map[string]int)


	for i := range pairs {
		pair := pairs[i]

		_, primaryPresent := peopleMap[pair.primaryPerson]
		if !primaryPresent {
			peopleMap[pair.primaryPerson] = len(peopleMap)
		}
		_, secondaryPresent := peopleMap[pair.secondaryPerson]
		if !secondaryPresent {
			peopleMap[pair.secondaryPerson] = len(peopleMap)
		}
	}

	return peopleMap
}

func parseLine(line string) HappinessPair {
	parts := strings.Split(line, " ")

	secondaryPerson := parts[len(parts) - 1]
	secondaryPerson = secondaryPerson[0:len(secondaryPerson) - 1]
	happinessDelta, _ := strconv.Atoi(parts[3])

	if parts[2] == "lose" {
		happinessDelta *= -1
	}

	return HappinessPair{
		primaryPerson: parts[0],
		secondaryPerson: secondaryPerson,
		happinessDelta: happinessDelta,
	}
}

func buildHappinessMatrix(pairs []HappinessPair) [][]int {
	people := buildReferenceMap(pairs)
	matrix := make([][]int, len(people))
	for i := range matrix {
		matrix[i] = make([]int, len(people))
	}

	for _, pair := range pairs {
		matrix[people[pair.primaryPerson]][people[pair.secondaryPerson]] = pair.happinessDelta
	}

	return matrix
}

func calculateNetHappiness(matrix [][]int, permutation []int) int {
	length := len(permutation)
	happiness := 0
	for i := 0; i < length ; i++ {
		happiness += matrix[permutation[i]][permutation[(i + 1) % length]]
		happiness += matrix[permutation[i]][permutation[(i - 1 + length) % length]]
	}
	return happiness
}
