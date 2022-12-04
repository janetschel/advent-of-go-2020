package main

import (
	"advent-of-go/utils/files"
	"fmt"
)

func main() {
	input := files.ReadFile(19, 2017, "\n")
	part1, part2 := followRoute(input)
	println(part1)
	println(part2)
}

func followRoute(input []string) (string, int) {
	dx, dy := 0, 1
	x, y := 0, 0

	whitespace := ' '
	for i, character := range input[0] {
		if character != whitespace {
			x = i
			break
		}
	}

	letters := ""
	steps := 0
	println(fmt.Sprintf("dimensions: %d x %d", len(input[y]), len(input)))
	for (y >= 0 && y < len(input)) && (x >= 0 && x < len(input[y]) ) {
		println(fmt.Sprintf("Processing %d, %d (%s) [%s]", x, y, string(input[y][x]), letters))
		current := rune(input[y][x])
		letter, isLetter := parseLetter(current)
		if isLetter {
			letters += letter
		} else if current == whitespace {
			println("Dead end")
			return letters, steps
		} else if current == '+' {
			if dx != 0 && y < len(input) - 1 && rune(input[y+1][x]) != whitespace {
				println("turning down")
				dx = 0
				dy = 1
			} else if dx != 0 && y > 0 && rune(input[y-1][x]) != whitespace {
				println("turning up")
				dx = 0
				dy = -1
			} else if dy != 0 && x < len(input[y]) - 1 && rune(input[y][x+1]) != whitespace {
				println("turning right")
				dy = 0
				dx = 1
			} else if dy != 0 && x > 0 && rune(input[y][x-1]) != whitespace {
				println("turning left")
				dy = 0
				dx = -1
			} else {
				println("Error, no new direction found")
				return letters, steps
			}
		}
		x += dx
		y += dy
		steps++
	}


	return letters, steps
}

func parseLetter(character rune) (string, bool) {
	if (character >= 'A' && character <= 'Z') ||
		(character >= 'a' && character <= 'z') {
			return string(character), true
		}
	return "", false
}
