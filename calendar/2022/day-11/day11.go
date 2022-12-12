package main

import (
	"advent-of-go/utils/files"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	inspections int

	items []int

	operation func(int, int) int
	operands []int

	test int
	newIfTrue int
	newIfFalse int
}

func main() {
	input := files.ReadFile(11, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	monkeys := parseInput(input)
	return solve(monkeys, 3, divide, 20)
}

func solvePart2(input []string) int {
	monkeys := parseInput(input)
	modifier := 1
	for _, m := range monkeys {
		modifier *= m.test
	}

	return solve(monkeys, modifier, modulo, 10000)
}

func solve(monkeys []*monkey, modifier int, modifierFunc func(int, int) int, rounds int) int {
	for i := 0; i < rounds; i++ {
		playRound(monkeys, modifier, modifierFunc)
	}
	
	inspections := make([]int, len(monkeys))
	for i, m := range monkeys {
		inspections[i] = m.inspections
	}

	sort.Ints(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func playRound(monkeys []*monkey, modifier int, modifierFunc func(int, int) int) {
	for _, m := range monkeys {
		for len(m.items) > 0 {
			worryLevel := m.items[0]
			m.items = m.items[1:]
			worryLevel = m.inspect(worryLevel, modifier, modifierFunc)
			dest := m.newDestination(worryLevel)
			monkeys[dest].items = append(monkeys[dest].items, worryLevel)
		}
	}
}

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func (m *monkey) inspect(worryLevel int, modifier int, modifierFunc func(int, int) int) int {
	a, b := worryLevel, worryLevel
	if len(m.operands) == 2 {
		a, b = m.operands[0], m.operands[1]
	} else if len(m.operands) == 1 {
		a = m.operands[0]
	}
	m.inspections++
	return modifierFunc(m.operation(a, b), modifier)
}

func (m monkey) newDestination(worryLevel int) int {
	if worryLevel % m.test == 0 {
		return m.newIfTrue
	}
	return m.newIfFalse
}

func divide(a int, b int) int {
	return a / b
}

func modulo(a int, b int) int {
	return a % b
}

func parseInput(input []string) []*monkey {
	monkeys := []*monkey{}
	for i := 1; i < len(input); i += 7 {
		itemsStr := strings.Split(strings.Split(input[i], ": ")[1], ", ")
		items := make([]int, len(itemsStr))
		for i, item := range itemsStr {
			items[i], _ = strconv.Atoi(item)
		}

		ops := strings.Fields(strings.Split(input[i + 1], "= ")[1])
		operands := []int{}
		op1, op1Err := strconv.Atoi(ops[0])
		op2, op2Err := strconv.Atoi(ops[2])
		if op1Err == nil {
			operands = append(operands, op1)
		}
		if op2Err == nil {
			operands = append(operands, op2)
		}

		operation := add
		if ops[1] == "*" {
			operation = multiply
		}

		test, _ := strconv.Atoi(strings.Fields(input[i+2])[3])

		newIfTrue, _ := strconv.Atoi(strings.Fields(input[i+3])[5])
		newIfFalse, _ := strconv.Atoi(strings.Fields(input[i+4])[5])

		m := monkey{ items: items, operation: operation, operands: operands, test: test, newIfTrue: newIfTrue, newIfFalse: newIfFalse }
		monkeys = append(monkeys, &m)
	}
	return monkeys
}
