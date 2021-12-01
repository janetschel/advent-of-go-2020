package main

import (
	"advent-of-go/utils/files"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(04, 2020, "\n\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	for _, passport := range input {
		if isValidForPart1(passport) {
			result++
		}
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	for _, passport := range input {
		if isValidForPart2(passport) {
			result++
		}
	}

	return result
}

func isValidForPart1(passport string) bool {
	requiredKeys := []string { "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid" }
	for _, key := range requiredKeys {
		if (!strings.Contains(passport, key + ":")) {
			return false
		}
	}
	return true
}

func isValidForPart2(passport string) bool {
	if (!isValidForPart1(passport)) {
		return false
	}

	passport = strings.ReplaceAll(passport, "\n", " ")
	pairs := strings.Split(passport, " ")
	for _, pair := range pairs {
		kvp := strings.Split(pair, ":")
		if !validateValue(kvp[1], kvp[0]) {
			return false
		}
	}

	return true
}

func validateValue(value string, key string) bool {
	// fmt.Printf("%s : %s\n", key, value)
	switch key {
	case "byr":
		if (len(value) != 4) {
			return false
		}
		num, _ := strconv.Atoi(value)
		return num >= 1920 && num <= 2002
	case "iyr":
		if (len(value) != 4) {
			return false
		}
		num, _ := strconv.Atoi(value)
		return num >= 2010 && num <= 2020
	case "eyr":
		if (len(value) != 4) {
			return false
		}
		num, _ := strconv.Atoi(value)
		return num >= 2020 && num <= 2030
	case "hgt":
		unit := value[len(value) - 2:]
		num, err := strconv.Atoi(value[:len(value) - 2])
		if (err != nil) {
			return false
		}
		switch unit {
		case "in":
			return num >= 59 && num <= 76
		case "cm":
			return num >= 150 && num <= 193
		default:
			return false
		}
	case "hcl":
		if (value[:1] != "#") {
			return false
		}
		_, err := strconv.ParseUint(value[1:], 16, 64)
		if (err != nil) {
			return false
		}
	case "ecl":
		validEyeColors := []string { "amb", "blu", "brn", "gry", "grn", "hzl", "oth" }
		for i, _ := range validEyeColors {
			if validEyeColors[i] == value {
				return true
			}
		}
		return false
	case "pid":
		if (len(value) != 9) {
			return false
		}
		_, err := strconv.Atoi(value)
		if (err != nil) {
			return false
		}
		return true
	default:
		return true
	}
	return true
}
