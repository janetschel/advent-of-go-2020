package main

import (
	"advent-of-go-2020/utils"
	"strings"
)

func main() {
	input := utils.ReadFile(2, "\n")
	pwds := utils.Filter(input, isPasswordValid)
	println(utils.Count(pwds))
}

func isPasswordValid(currentElement string) bool {
	var numbers, char, pwd, lo, hi string
	utils.Unpack(strings.Split(currentElement, " "), &numbers, &char, &pwd)
	utils.Unpack(strings.Split(numbers, "-"), &lo, &hi)

	numChars := strings.Count(pwd, char[:len(char) - 1])
	return numChars >= utils.ToInt(lo) && numChars <= utils.ToInt(hi)
}
