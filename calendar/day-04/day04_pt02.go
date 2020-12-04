package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"advent-of-go-2020/utils/slices"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(4, "\n\n")
	println(solvePart2(input))
}

func solvePart2(input []string) int {
	fields := []string {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	valid := 0

	for _, element := range input {
		for i, field := range fields {
			if !strings.Contains(element, field) {
				break
			} else if i == 6 && validate(element) {
				valid++
			}
		}
	}

	return valid
}

func validate(element string) bool {
	parts := make(map[string]string)

	for _, curr := range strings.Fields(element) {
		part := strings.Split(curr, ":")
		parts[part[0]] = part[1]
	}

	byr, iyr, eyr := conv.ToInt(parts["byr"]), conv.ToInt(parts["iyr"]), conv.ToInt(parts["eyr"])
	if (byr < 1920 || byr > 2002) || (iyr < 2010 || iyr > 2020) || (eyr < 2020 || eyr > 2030) {
		return false
	}

	// Go does not support line-splitting in if-statements. Since my lines would be too long, I concat them with else-if
	unit, hgt := parts["hgt"][len(parts["hgt"]) - 2:], conv.ToInt(parts["hgt"][:len(parts["hgt"]) - 2])
	if unit != "cm" && unit != "in" {
		return false
	} else if (unit == "cm" && (hgt < 150 ||  hgt > 193)) || (unit == "in" && (hgt < 59 || hgt > 76)) {
		return false
	}else if matched, err := regexp.MatchString("#([0-9a-f]){6}", parts["hcl"]); !matched || err != nil {
		return false
	} else if _, err := strconv.Atoi(parts["pid"]); err != nil || len(parts["pid"]) != 9 {
		return false
	}

	return slices.Contains([]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}, parts["ecl"])
}
