package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"advent-of-go-2020/utils/maths"
	"strings"
)

func main() {
	input := files.ReadFile(13, "\n")
	println(solvePart1(input))
}

func solvePart1(input []string) int {
    earliestDep := conv.ToInt(input[0])
    depTime := parseLine(input[1])

    earliest := 100000000000
    id := 0
    for _, element := range depTime {
    	curr := 0
    	for curr <= earliestDep {
    		curr += element
		}

		if curr < earliest {
			id = element
			earliest = curr
		}
	}

	minutes := earliestDep - earliest
    return maths.Abs(minutes) * id
}

func parseLine(input string) []int {
	split := strings.Split(input, ",")
	result := make([]int, 0)

	for _, element := range split {
		if element != "x" {
			result = append(result, conv.ToInt(element))
		}
	}

	return result
}