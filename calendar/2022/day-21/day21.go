package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(21, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	monkeys := parseInput(input)
	return getValue("root", monkeys)
}

func solvePart2(input []string) int {	
	monkeys := parseInput(input)
	// these increments are too big and the comparison is backwards for the sample
	i, increment := 100000000000, 100000000000
	for i > 0 {
		monkeys["humn"][0] = fmt.Sprint(i)
		v1 := getValue(monkeys["root"][0], monkeys)
		v2 := getValue(monkeys["root"][2], monkeys)
		if v1 == v2 {
			return i
		} else if v2 < v1 {
			i += increment
		} else {
			i -= increment
			increment /= 10
		}
	}
	return -1
}

func getValue(monkey string, monkeys map[string][]string) int {
	m := monkeys[monkey]
	if len(m) == 1 {
		value, _ := strconv.Atoi(m[0])
		return value
	} else if m[1] == "+" {
		return getValue(m[0], monkeys) + getValue(m[2], monkeys)
	} else if m[1] == "-" {
		return getValue(m[0], monkeys) - getValue(m[2], monkeys)
	} else if m[1] == "/" {
		return getValue(m[0], monkeys) / getValue(m[2], monkeys)
	} else {
		return getValue(m[0], monkeys) * getValue(m[2], monkeys)
	}
}

func parseInput(input []string) map[string][]string {
	monkeys := map[string][]string{}
	for _, line := range input {
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return strings.ContainsRune(" :", r)
		})
		monkeys[parts[0]] = parts[1:]
	}
	return monkeys
}
