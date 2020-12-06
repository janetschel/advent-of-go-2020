package main

import (
	"advent-of-go-2020/utils/files"
	"go/types"
	"strings"
)

func main() {
	input := files.ReadFile(6, "\n\n")
	println("Count answers part 1:", solve(input))
}

func solve(input []string) int {
	count := 0
	for _, curr := range input {
		count += answers(curr)
	}

	return count
}

func answers(group string) int{
	givenAnswers := make(map[string]types.Nil)
	count := 0

	for _, line := range strings.Split(group, "\n") {
		for _, char := range line {
			if _, present := givenAnswers[string(char)]; !present {
				givenAnswers[string(char)] = types.Nil{} // make it present
				count++
			}
		}
	}

	return count
}
