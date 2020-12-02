package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"advent-of-go-2020/utils/slices"
	"advent-of-go-2020/utils/str"
	"strings"
)

func main() {
	input := files.ReadFile(2, "\n")
	pwds := slices.Filter(input, isPasswordValidCharPos)
	println(slices.Count(pwds))
}

func isPasswordValidCharPos(currentElement string) bool {
	var numbers, char, pwd, lo, hi string
	slices.Unpack(strings.Split(currentElement, " "), &numbers, &char, &pwd)
	slices.Unpack(strings.Split(numbers, "-"), &lo, &hi)

	i, errFirst := str.CharAt(pwd, conv.ToInt(lo) - 1)
	j, errSecond := str.CharAt(pwd, conv.ToInt(hi) - 1)

	return errFirst == nil && errSecond == nil && (i == char[:len(char) - 1]) != (j == char[:len(char) - 1])
}
