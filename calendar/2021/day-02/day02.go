package main

import (
	"strconv"
	"strings"
	"tblue-aoc-2021/utils/files"
)

type Direction struct {
	direction string
	amount    int
}

func main() {
	input := files.ReadFile(02, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	directions := make([]Direction, len(input))
	for i, v := range input {
		s := strings.Split(v, " ")
		intVar, _ := strconv.Atoi(s[1])
		d := Direction{direction: s[0], amount: intVar}
		directions[i] = d
	}
	vertical := 0
	horizontal := 0
	for _, d := range directions {
		if d.direction == "forward" {
			horizontal += d.amount
		}
		if d.direction == "up" {
			vertical -= d.amount
		}
		if d.direction == "down" {
			vertical += d.amount
		}
	}

	return vertical * horizontal
}

func solvePart2(input []string) int {
	directions := make([]Direction, len(input))
	for i, v := range input {
		s := strings.Split(v, " ")
		intVar, _ := strconv.Atoi(s[1])
		d := Direction{direction: s[0], amount: intVar}
		directions[i] = d
	}
	aim := 0
	vertical := 0
	horizontal := 0
	for _, d := range directions {
		if d.direction == "forward" {
			horizontal += d.amount
			vertical += (d.amount * aim)
		}
		if d.direction == "up" {
			aim -= d.amount
		}
		if d.direction == "down" {
			aim += d.amount
		}
	}

	return vertical * horizontal
}
