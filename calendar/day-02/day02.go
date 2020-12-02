package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"advent-of-go-2020/utils/slices"
	"strings"
)

func main() {
	input := files.ReadFile(2, "\n")
	pwds := slices.Filter(input, isPasswordValid)
	println(len(pwds))
}

func isPasswordValid(currentElement string) bool {
	var numbers, char, pwd, lo, hi string
	slices.Unpack(strings.Split(currentElement, " "), &numbers, &char, &pwd)
	slices.Unpack(strings.Split(numbers, "-"), &lo, &hi)

	numChars := strings.Count(pwd, char[:len(char) - 1])
	return numChars >= conv.ToInt(lo) && numChars <= conv.ToInt(hi)
}
