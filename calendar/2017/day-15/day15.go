package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(15, 2017, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	factorA, factorB := parseInput(input)
	generatorA, generatorB := []int{factorA}, []int{factorB}

	matches := 0

	for i := 0; i < 40000000; i++ {
		generatorA, generatorB = generate(generatorA, 16807), generate(generatorB, 48271)
		lastA, lastB := generatorA[len(generatorA) - 1], generatorB[len(generatorB) - 1]
		if match(lastA, lastB) {
			matches++
		}
	}

	return matches
}

func solvePart2(input []string) int {
	factorA, factorB := parseInput(input)
	generatorA, generatorB := []int{factorA}, []int{factorB}

	matches := 0

	for i := 0; i < 5000000; i++ {
		generatorA, generatorB = generate(generatorA, 16807), generate(generatorB, 48271)
		for generatorA[len(generatorA) - 1] % 4 != 0 {
			generatorA = generate(generatorA, 16807)
		}
		for generatorB[len(generatorB) - 1] % 8 != 0 {
			generatorB = generate(generatorB, 48271)
		}
		lastA, lastB := generatorA[len(generatorA) - 1], generatorB[len(generatorB) - 1]
		if match(lastA, lastB) {
			matches++
		}
	}

	return matches
}

func generate(generator []int, factor int) []int {
	lastProduced := generator[len(generator) - 1]
	product := factor * lastProduced
	remainder := product % 2147483647
	return append(generator, remainder)
}

func match(valueA int, valueB int) bool {
	binA, binB := fmt.Sprintf("%016b", valueA), fmt.Sprintf("%016b", valueB)
	return binA[len(binA) - 16:] == binB[len(binB) - 16:]
}

func parseInput(input []string) (int, int) {
	partsA, partsB := strings.Fields(input[0]), strings.Fields(input[1])
	factorA, _ := strconv.Atoi(partsA[len(partsA) - 1])
	factorB, _ := strconv.Atoi(partsB[len(partsB) - 1])
	return factorA, factorB
}
