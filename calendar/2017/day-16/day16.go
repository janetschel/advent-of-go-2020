package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/sets"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(16, 2017, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) string {
	programs := "abcdefghijklmnop"
	moves := strings.Split(input[0], ",")
	
	return runMoves(programs, moves)
}

func solvePart2(input []string) string {
	programs, dances := "abcdefghijklmnop", 1000000000
	moves := strings.Split(input[0], ",")

	knownPrograms, cycle := sets.New(), []string{programs}
	knownPrograms.Add(programs)

	for i := 0; i < dances; i++ {
		programs = runMoves(programs, moves)
		if knownPrograms.Has(programs) {
			break
		}
		knownPrograms.Add(programs)
		cycle = append(cycle, programs)
	}
	
	return cycle[dances % knownPrograms.Size()]
}

func runMoves(programs string, moves []string) string {
	for _, move := range moves {
		moveType := move[0]
		switch moveType {
		case 's':
			n, _ := strconv.Atoi(move[1:])
			programs = programs[len(programs) - n:] + programs[0:len(programs) - n]
		case 'x':
			indexes := strings.Split(move[1:], "/")
			indexA, _ := strconv.Atoi(indexes[0])
			indexB, _ := strconv.Atoi(indexes[1])
			str := []rune(programs)
			tmp := str[indexA]
			str[indexA] = str[indexB]
			str[indexB] = tmp
			programs = string(str)
		case 'p':
			swapA, swapB := rune(move[1]), rune(move[3])
			indexA, indexB := strings.IndexRune(string(programs), swapA), strings.IndexRune(string(programs), swapB)
			str := []rune(programs)
			str[indexA] = swapB
			str[indexB] = swapA
			programs = string(str)
		}
	}

	return programs
}
