package main

import (
	"advent-of-go-2020/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadFile(2, "\n")

	numberOfKnownValidPasswords := 0

	for _, currentElement := range input {
		if isPasswordValidCharPos(currentElement) {
			numberOfKnownValidPasswords++
		}
	}

	fmt.Printf("Number of valid passwords: %d", numberOfKnownValidPasswords)
}

func isPasswordValidCharPos(currentElement string) bool {
	parts := strings.Split(currentElement, ":")
	numbers := strings.Split(parts[0], "-")

	i, _ := strconv.Atoi(numbers[0])
	j, _ := strconv.Atoi(numbers[1][:len(numbers[1]) - 2])
	char := numbers[1][len(numbers[1]) - 1:]
	pwd := parts[1][1:]

	// Indices are not zero based in input
	i--
	j--

	if i > len(pwd) {
		return false
	} else if j > len(pwd) - 1 {
		return string(pwd[i]) == char
	}

	validFirstPos := string(pwd[i]) == char && string(pwd[j]) != char
	validSecondPos := string(pwd[i]) != char && string(pwd[j]) == char

	return validFirstPos != validSecondPos // XOR
}
