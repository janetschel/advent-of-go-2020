package main

import (
	"advent-of-go/utils/files"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(24, 2017, "\n")
	solution := solve(input)
	println(solution.strength)
	println(solution.strengthForLongest)
}

type component struct {
	a, b int
	used bool
}

type maxValues struct {
	length int
	strength int
	strengthForLongest int
}

func solve(input []string) maxValues {
	components := parseInput(input)
	maxes := build(components, 0, 0, 0, &maxValues{0,0,0})
	return maxes
}

func solvePart2(input []string) int {
	result := 0



	return result
}

func build(components []component, pins int, strength int, length int, maxes *maxValues) maxValues {
	if strength > maxes.strength {
		maxes.strength = strength
	}
	if length > maxes.length {
		maxes.length = length
	}

	if length == maxes.length && strength > maxes.strengthForLongest  {
		maxes.strengthForLongest = strength
	}

	for i := 0; i < len(components); i++ {
		c := &components[i]
		if !c.used && (c.a == pins || c.b == pins) {
			c.used = true
			usedPins := c.a
			if c.a == pins {
				usedPins = c.b
			}
			build(components, usedPins, strength + c.a + c.b, length + 1, maxes)
			c.used = false
		}
	}
	return *maxes
}

func parseInput(input []string) []component {
	components := make([]component, len(input))
	for i, c := range input {
		pins := strings.Split(c, "/")
		a, _ := strconv.Atoi(pins[0])
		b, _ := strconv.Atoi(pins[1])

		components[i] = component{ a: a, b: b, used: false}
	}
	return components
}
