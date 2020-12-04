package main

import (
	"advent-of-go-2020/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(4, "\n\n")
	println(solve(input))
}

func solve(input []string) int {
	fields := []string {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	valid := 0

	for _, element := range input {
		for i, field := range fields {
			if !strings.Contains(element, field) {
				break
			} else if i == 6 {
				valid++
			}
		}
	}

	return valid
}
