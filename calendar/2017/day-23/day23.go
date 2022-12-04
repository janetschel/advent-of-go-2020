package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"strconv"
	"strings"
)

type operation string
const (
	set operation = "set"
	sub = "sub"
	multiply = "mul"
	jumpNotZero = "jnz"
)

func main() {
	input := files.ReadFile(23, 2017, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	registers := map[string]int{}
	increment := 1
	for i := 0; i < len(input); i += increment {
		increment = 1
		parts := strings.Fields(input[i])
		register := parts[1]

		y := 0
		if len(parts) > 2 {
			var isRegister bool
			y, isRegister = registers[parts[2]]
			if !isRegister {
				y, _ = strconv.Atoi(parts[2])
			}
		}
		switch operation(parts[0]) {
		case set:
			registers[register] = y
		case sub:
			registers[register] -= y
		case multiply:
			registers[register] *= y
			result++
		case jumpNotZero:
			value, isRegister := registers[register]
			if !isRegister {
				value, _ = strconv.Atoi(register)
			}
			if value != 0 {
				increment = y
			}
		}
	}

	return result
}

func solvePart2(input []string) int {
	a,b,c,g,h := 1,0,0,0,0

	b = 99
	c = b

	if a != 0 {
		b = (b * 100) + 100000
		c = b + 17000
	}
	for g != 0 || h == 0 {
		if !maths.IsPrime(b) {
			h++
		}
		g = b - c
		b += 17
	}

	return h
}
