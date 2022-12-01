package main

import (
	"advent-of-go/utils/files"
)

func main() {
	input := files.ReadFile(9, 2017, "\n")
	part1, part2 := solve(input)
	println(part1)
	println(part2)
}

func solve(input []string) (int, int) {
	stream := input[0]

	groupStack := []byte{}
	garbageStack := []byte{}
	groups, score, garbage := 0, 0, 0

	for i := 0; i < len(stream); i++ {
		current := stream[i]

		if current == '!' {
			i++
		} else if len(garbageStack) == 0 && current == '<' {
			garbageStack = append(garbageStack, current)
		} else if len(garbageStack) > 0 && current == '>' {
			garbageStack = garbageStack[1:]
		} else if len(garbageStack) == 0 &&  current == '{' {
			groupStack = append(groupStack, current)
		} else if len(garbageStack) == 0 &&  current == '}' {
			score += len(groupStack)
			groupStack = groupStack[1:]
			groups++
		} else if len(garbageStack) > 0 {
			garbage++
		}
	}

	return score, garbage
}
