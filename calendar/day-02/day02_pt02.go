package main

import (
	"advent-of-go-2020/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadFile(2, "\n")
	pwds := utils.Filter(input, isPasswordValidCharPos)
	num := utils.Count(pwds)

	fmt.Printf("Number of valid passwords: %d", num)
}

func isPasswordValidCharPos(currentElement string) bool {
	parts := strings.Split(currentElement, " ")

	var numbers, char, pwd, i, j string
	utils.Unpack(parts, &numbers, &char, &pwd)
	utils.Unpack(strings.Split(numbers, "-"), &i, &j)

	char = char[:len(char) - 1]
	first, errFirst := utils.CharAt(pwd, utils.ToInt(i) - 1)
	second, errSecond := utils.CharAt(pwd, utils.ToInt(j) - 1)

	return errFirst == nil && errSecond == nil && (first == char) != (second == char)
}
