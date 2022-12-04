package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(06, 2017, "\n")
	part1, part2 := solve(input)
	println(part1)
	println(part2)
}

func solve(input []string) (int, int) {
	banksStrings := strings.Fields(input[0])
	banks := make([]int, len(banksStrings))
	for i, b := range banksStrings {
		bank, _ := strconv.Atoi(b)
		banks[i] = bank
	}

	configurations := make(map[string]int)
	cycles := 0
	lastCycle, contains := 0, false
	for ; !contains; cycles++ {
		lastCycle, contains = configurations[getKey(banks)]
		if contains {
			return cycles, cycles - lastCycle
		}
		configurations[getKey(banks)] = cycles
		toRedistribute := slices.Max(banks)
		largestBank := slices.IndexOfInt(toRedistribute, banks)
		banks[largestBank] = 0
		for i := 1; i <= toRedistribute; i++ {
			banks[(largestBank + i) % len(banks)]++
		}
	}
	return cycles, cycles - lastCycle
}

func getKey(banks []int) string {
	banksStrings := make([]string, len(banks))
	for i, blocks := range banks {
		banksStrings[i] = fmt.Sprintf("%d", blocks)
	}
	return strings.Join(banksStrings, ",")
}
