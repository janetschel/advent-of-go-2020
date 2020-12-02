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
		if isPasswordValid(currentElement) {
			numberOfKnownValidPasswords++
		}
	}

	fmt.Printf("Number of valid passwords: %d", numberOfKnownValidPasswords)
}

func isPasswordValid(currentElement string) bool {
	partsOfLine := strings.Split(currentElement, ":")
	numbers := strings.Split(partsOfLine[0], "-")

	lowerBound, _ := strconv.Atoi(numbers[0])
	upperBound, _ := strconv.Atoi(numbers[1][:len(numbers[1]) - 2])
	charToSearch := numbers[1][len(numbers[1]) - 1:]
	password := partsOfLine[1][1:]

	currentNumberOfChar := 0
	for _, char := range password {
		if string(char) == charToSearch {
			currentNumberOfChar++
		}

		if currentNumberOfChar > upperBound {
			return false
		}
	}

	return currentNumberOfChar >= lowerBound
}
