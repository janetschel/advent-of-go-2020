package main

import (
	"advent-of-go-2020/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadFile(2, "\n")
	pwds := utils.Filter(input, isPasswordValid)
	num := utils.Count(pwds)

	fmt.Printf("Number of valid passwords: %d", num)
}

func isPasswordValid(currentElement string) bool {
	parts := strings.Split(currentElement, " ")

	var numbers, char, pwd, lo, hi string
	utils.Unpack(parts, &numbers, &char, &pwd)
	utils.Unpack(strings.Split(numbers, "-"), &lo, &hi)

	numChars := strings.Count(pwd, char[:len(char) - 1])
	return numChars >= utils.ToInt(lo) && numChars <= utils.ToInt(hi)
}
