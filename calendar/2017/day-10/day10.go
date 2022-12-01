package main

import (
	"advent-of-go/utils/files"
	knothash "advent-of-go/utils/knotHash"
)

func main() {
	input := files.ReadFile(10, 2017, "\n")
	println(knothash.HashInt(input[0]))
	println(knothash.HashHex(input[0]))
}
