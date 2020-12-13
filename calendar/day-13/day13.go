package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"advent-of-go-2020/utils/maths"
	"strings"
)

func main() {
	input := files.ReadFile(13, "\n")
	println("Earliest bus I can take:", solvePart1(input))
}

func solvePart1(input []string) int {
    earliestDep, depTimes := conv.ToInt(input[0]), parseLine(input[1])
    earliest, id := maths.MaxInt(), 0

    for _, element := range depTimes {
    	curr := 0
    	for curr <= earliestDep {
    		curr += element
		}

		if curr < earliest {
			id = element
			earliest = curr
		}
	}

    return maths.Abs(earliestDep - earliest) * id
}

func parseLine(input string) []int {
	result := make([]int, 0)

	for _, element := range strings.Split(input, ",") {
		if element != "x" {
			result = append(result, conv.ToInt(element))
		}
	}

	return result
}