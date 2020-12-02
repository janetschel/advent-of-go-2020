package main

import (
	"advent-of-go-2020/utils"
	"strings"
)

func main() {
	input := utils.ReadFile(2, "\n")
	pwds := utils.Filter(input, isPasswordValidCharPos)
	println(utils.Count(pwds))
}

func isPasswordValidCharPos(currentElement string) bool {
	var numbers, char, pwd, lo, hi string
	utils.Unpack(strings.Split(currentElement, " "), &numbers, &char, &pwd)
	utils.Unpack(strings.Split(numbers, "-"), &lo, &hi)

	i, errFirst := utils.CharAt(pwd, utils.ToInt(lo) - 1)
	j, errSecond := utils.CharAt(pwd, utils.ToInt(hi) - 1)

	return errFirst == nil && errSecond == nil && (i == char[:len(char) - 1]) != (j == char[:len(char) - 1])
}
