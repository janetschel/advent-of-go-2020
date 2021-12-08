package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(16, 2015, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	return parseSueNumber(slices.Filter(input, filterPartOne)[0])
}

func solvePart2(input []string) int {
	sues := slices.Filter(input, filterPartTwo)
	return parseSueNumber(sues[0])
}

func parseSueNumber(sue string) int {
	sueParts := strings.Split(sue, " ")
	sueNumber, _ := strconv.Atoi(sueParts[1][0:len(sueParts[1]) - 1])

	return sueNumber
}

func filterPartOne(sue string) bool {
	if strings.Contains(sue, "children") && !strings.Contains(sue, "children: 3") {
		return false
	}
	if strings.Contains(sue, "cats") && !strings.Contains(sue, "cats: 7") {
		return false
	}
	if strings.Contains(sue, "samoyeds") && !strings.Contains(sue, "samoyeds: 2") {
		return false
	}
	if strings.Contains(sue, "pomeranians") && !strings.Contains(sue, "pomeranians: 3") {
		return false
	}
	if strings.Contains(sue, "akitas") && !strings.Contains(sue, "akitas: 0") {
		return false
	}
	if strings.Contains(sue, "vizslas") && !strings.Contains(sue, "vizslas: 0") {
		return false
	}
	if strings.Contains(sue, "goldfish") && !strings.Contains(sue, "goldfish: 5") {
		return false
	}
	if strings.Contains(sue, "trees") && !strings.Contains(sue, "trees: 3") {
		return false
	}
	if strings.Contains(sue, "cars") && !strings.Contains(sue, "cars: 2") {
		return false
	}
	if strings.Contains(sue, "perfumes") && !strings.Contains(sue, "perfumes: 1") {
		return false
	}
	return true
}

func filterPartTwo(sue string) bool {
	if strings.Contains(sue, "children") && !strings.Contains(sue, "children: 3") {
		return false
	}
	if strings.Contains(sue, "cats") {
		if (parseQuantity(sue, "cats") <= 7) {
			return false
		}
	}
	if strings.Contains(sue, "samoyeds") && !strings.Contains(sue, "samoyeds: 2") {
		return false
	}
	if strings.Contains(sue, "pomeranians") {
		if (parseQuantity(sue, "pomeranians") >= 3) {
			return false
		}
	}
	if strings.Contains(sue, "akitas") && !strings.Contains(sue, "akitas: 0") {
		return false
	}
	if strings.Contains(sue, "vizslas") && !strings.Contains(sue, "vizslas: 0") {
		return false
	}
	if strings.Contains(sue, "goldfish") {
		if (parseQuantity(sue, "goldfish") >= 5) {
			return false
		}
	}
	if strings.Contains(sue, "trees") {
		if (parseQuantity(sue, "trees") <= 3) {
			return false
		}
	}
	if strings.Contains(sue, "cars") && !strings.Contains(sue, "cars: 2") {
		return false
	}
	if strings.Contains(sue, "perfumes") && !strings.Contains(sue, "perfumes: 1") {
		return false
	}
	return true
}

func parseQuantity(str string, item string) int {
	parts := strings.Split(str, item + ": ")
	numStr := strings.Split(parts[1], ",")[0]
	num, _ := strconv.Atoi(numStr)
	return num
}
