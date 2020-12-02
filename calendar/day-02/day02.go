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
	var i, j, char, pwd string
	slices.ParseLine(currentElement, "(-)|(:\\s)|\\s", &i, &j, &char, &pwd)

	numChars := strings.Count(pwd, char)
	return numChars >= conv.ToInt(i) && numChars <= conv.ToInt(j)
}
