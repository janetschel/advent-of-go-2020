package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/slices"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(10, 2022, "\n")
	part1, part2 := solve(input)
	println(part1)
	println(part2)
}

var noop = "noop"
var strengthCycles = []int{20, 60, 100, 140, 180, 220}
var lit, dark = "█", " "
// not my best work, this is pretty jumbled
func solve(input []string) (int, string) {
	rows := []string{""}

	x, wait, cycle, i, value, strength, row, cursor := 1, 1, 0, 0, 0, 0, 0, 0
	for wait > 0 {
		wait--
		cycle++

		if wait == 0 && i < len(input) {
			x += value
			parts := strings.Fields(input[i])
			if parts[0] != noop {
				value, _ = strconv.Atoi(parts[1])
				wait = 2
			} else {
				value = 0
				wait = 1
			}
			i++
		}

		if maths.Abs(cursor - x) <= 1 {
			rows[row] += lit
		} else {
			rows[row] += dark
		}

		cursor++
		if slices.Contains(strengthCycles, cycle) {
			strength += cycle * x
		}
		if cycle % 40 == 0 {
			cursor = 0
			row++
			rows = append(rows, "")
		}
	}

	return strength, strings.Join(rows, "\n")
}
