package main

import (
	"advent-of-go/utils/files"
	"math"
	"strconv"
)

type snElement struct {
	value int
	level int
}

type snailNumber []snElement

func main() {
	input := files.ReadFile(18, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	numbers := parseInput(input)
	sum := numbers[0]
	for i := 1; i < len(numbers); i++ {
		sum = add(sum, numbers[i])
	}
	return magnitude(sum)
}

func solvePart2(input []string) int {
	numbers := parseInput(input)
	largestMagnitude := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j {
				sum := add(numbers[i], numbers[j])
				magnitude := magnitude(sum)
				if magnitude > largestMagnitude {
					largestMagnitude = magnitude
				}
			}
		}
	}
	return largestMagnitude
}

func add(number1 snailNumber, number2 snailNumber) snailNumber {
	sum := make(snailNumber, len(number1))
	copy(sum, number1)
	sum = append(sum, number2...)
	for i := range sum {
		sum[i].level++
	}
	return reduce(sum)
}

func reduce(sn snailNumber) snailNumber {
	reduced := make(snailNumber, len(sn))
	copy(reduced, sn)
	action := nextAction(reduced)
	for action != nil {
		reduced = action(reduced)
		action = nextAction(reduced)
	}
	return reduced
}

func parseInput(input []string) []snailNumber {
	sn := []snailNumber{}
	for i := range input {
		sn = append(sn, parseSnailNumber(input[i]))
	}
	return sn
}

func parseSnailNumber(input string) snailNumber {
	level, number := -1, snailNumber{}

	for i := 0; i < len(input); i++ {
		current := input[i : i+1]
		if current == "[" {
			level++
		} else if current == "]" {
			level--
		} else if current != "," {
			val, _ := strconv.Atoi(current)
			number = append(number, snElement{value: val, level: level})
		}
	}

	return number
}

func nextAction(sn snailNumber) func(snailNumber) snailNumber {
	shouldSplit := false
	for i := range sn {
		if sn[i].level >= 4 {
			return explode
		}
		if sn[i].value >= 10 {
			shouldSplit = true
		}
	}

	if shouldSplit {
		return split
	}
	return nil
}

func explode(sn snailNumber) snailNumber {
	newNumber, exploded := snailNumber{}, 0
	for i := 0; i < len(sn); i++ {
		elem := sn[i]
		if elem.level >= 4 && exploded < 2 {
			if i > 0 && exploded == 0 {
				newNumber[i-1].value += elem.value
			}
			if exploded == 0 {
				newNumber = append(newNumber, snElement{level: elem.level - 1, value: 0})
			}
			if i < len(sn)-1 && exploded == 1 {
				sn[i+1].value += elem.value
			}

			exploded++
		} else {
			newNumber = append(newNumber, elem)
		}
	}
	return newNumber
}

func split(sn snailNumber) snailNumber {
	newNumber := snailNumber{}
	hasSplit := false
	for _, elem := range sn {
		if elem.value >= 10 && !hasSplit {
			hasSplit = true
			roundDown := elem.value / 2
			newNumber = append(newNumber, snElement{level: elem.level + 1, value: roundDown})
			newNumber = append(newNumber, snElement{level: elem.level + 1, value: elem.value - roundDown})
		} else {
			newNumber = append(newNumber, elem)
		}
	}
	return newNumber
}

func magnitude(sn snailNumber) int {
	num := make(snailNumber, len(sn))
	copy(num, sn)
	for l := 1; l <= 4; l++ {
		length, level := len(num), 4-l
		for n := 0; n < length-1; n++ {
			if num[n].level == level && num[n+1].level == level {
				mag := (3 * num[n].value) + (2 * num[n+1].value)
				num[n].value = mag
				num[n].level--
				num[n+1].level = math.MaxInt
			}
		}

		newNum := snailNumber{}
		for n := 0; n < len(num); n++ {
			if num[n].level != math.MaxInt {
				newNum = append(newNum, num[n])
			}
		}
		num = newNum
	}
	return num[0].value
}
