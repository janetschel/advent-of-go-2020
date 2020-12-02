package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"advent-of-go-2020/utils/slices"
	"advent-of-go-2020/utils/str"
)

func main() {
	input := files.ReadFile(2, "\n")
	pwds := slices.Filter(input, isPasswordValidCharPos)
	println(len(pwds))
}

func isPasswordValidCharPos(currentElement string) bool {
	var i, j, char, pwd string
	slices.ParseLine(currentElement, "(-)|(:\\s)|\\s", &i, &j, &char, &pwd)

	i, errFirst := str.CharAt(pwd, conv.ToInt(i) - 1)
	j, errSecond := str.CharAt(pwd, conv.ToInt(j) - 1)

	return errFirst == nil && errSecond == nil && (i == char) != (j == char)
}
