package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"strconv"
	"strings"
)

type keyValues struct {
	divide int
	add int
	offset int
}

func main() {
	input := files.ReadFile(24, 2021, "\n")
	part1, part2 := solve(input)
	println(part1)
	println(part2)
}

func solve(input []string) (string, string) {
	digits := map[int][2]int{}
	kv := parseInput(input)
	stack := [][2]int{}
	for dig, inst := range kv {
		if inst.divide == 1 {
			stack = append([][2]int{{ dig, inst.offset}}, stack...)
		} else {
			sibling, add := stack[0][0], stack[0][1]
			stack = stack[1:]
			diff := add + inst.add
			if diff < 0 {
				digits[sibling] = [2]int{ (-1 * diff) + 1, 9}
				digits[dig] = [2]int{1, 9 + diff}
			} else {
				digits[sibling] = [2]int{1, 9 - diff}
				digits[dig] = [2]int{1 + diff, 9}
			}
		}
	}
	minStr, maxStr := "", ""
	for i := 0; i < 14; i++ {
		minStr += fmt.Sprint(digits[i][0])
		maxStr += fmt.Sprint(digits[i][1])
	}

	return maxStr, minStr
}

func parseInput(input []string) [14]keyValues {
	values := [14]keyValues{}

	for i := 0; i < 14; i++ {
		startingLine := 18 * i
		div, _ := strconv.Atoi(strings.Fields(input[startingLine+4])[2])
		check, _ := strconv.Atoi(strings.Fields(input[startingLine+5])[2])
		offset, _ := strconv.Atoi(strings.Fields(input[startingLine+15])[2])
		values[i] = keyValues{ divide: div, add: check, offset: offset }
	}

	return values
}
