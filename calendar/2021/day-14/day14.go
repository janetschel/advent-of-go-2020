package main

import (
	"advent-of-go/utils/files"
	"math"
	"strings"
)

func main() {
	input := files.ReadFile(14, 2021, "\n\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int64 {
	template, rules := parseInput(input)
	initial := template
	pairs := makePairs(template)
	for i := 0; i < 10; i++ {
		pairs = pairInsert(pairs, rules)
	}
	return scorePairs(pairs, initial)
}

func solvePart2(input []string) int64 {
	template, rules := parseInput(input)
	initial := template
	pairs := makePairs(template)
	for i := 0; i < 40; i++ {
		pairs = pairInsert(pairs, rules)
	}

	return scorePairs(pairs, initial)
}

func scorePolymer(polymer string) int {
	frequency := map[rune]int{}
	for _, char := range polymer {
		frequency[char]++
	}

	mostFrequent, leastFrequent := 0, math.MaxInt
	for _, f := range frequency {
		if f > mostFrequent {
			mostFrequent = f
		}
		if f < leastFrequent {
			leastFrequent = f
		}
	}
	return mostFrequent - leastFrequent
}

func scorePairs(pairs map[string]int64, initial string) int64 {
	frequencies := map[string]int64{}
	for pair, frequency := range pairs {
		frequencies[pair[1:]] += frequency
	}
	frequencies[initial[:1]]++

	mostFrequent, leastFrequent := int64(0), int64(math.MaxInt64)
	for _, f := range frequencies {
		if f > mostFrequent {
			mostFrequent = f
		}
		if f < leastFrequent {
			leastFrequent = f
		}
	}
	return mostFrequent - leastFrequent
}

func performPairInsertions(template string, rules map[string]string) string {
	newString := ""
	for i := 1; i < len(template); i++ {
		insert := rules[template[i-1:i+1]]
		newString += template[i-1:i] + insert
	}
	newString += template[len(template)-1:]
	return newString
}

func parseInput(input []string) (string, map[string]string) {
	rules := map[string]string{}
	template := input[0]
	subs := strings.Split(input[1], "\n")
	for i := range subs {
		parts := strings.Split(subs[i], " -> ")
		rules[parts[0]] = parts[1]
	}
	return template, rules
}

func makePairs(template string) map[string]int64 {
	pairs := map[string]int64{}
	for i := 1; i < len(template); i++ {
		pairs[template[i-1:i+1]]++
	}
	return pairs
}

func pairInsert(pairs map[string]int64, rules map[string]string) map[string]int64 {
	newPairs := map[string]int64{}
	for pair, frequency := range pairs {
		insert := rules[pair]

		newPairs[pair[:1]+insert] += frequency
		newPairs[insert+pair[1:]] += frequency
	}
	return newPairs
}
