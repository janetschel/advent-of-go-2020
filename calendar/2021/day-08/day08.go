package main

import (
	"advent-of-go/utils/files"
	"math"
	"sort"
	"strings"
)

func main() {
	input := files.ReadFile(8, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	for i := range input {
		parts := strings.Fields(strings.Split(input[i], "|")[1])
		for j := range parts {
			length := len(parts[j])
			if length == 2 || length == 3 || length == 4 || length == 7 {
				result++
			}
		}
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	for i := range input {
		value := decodeInput(input[i])
		result += value
	}

	return result
}

func decodeInput(input string) int {
	parts := strings.Split(input, "|")
	decoder := buildDecoderMap(strings.Fields(parts[0]))
	value := 0
	digits := strings.Fields(parts[1])
	for i, digit := range digits {
		value += int(math.Pow10(len(digits)-1-i)) * decoder[sortString(digit)]
	}
	return value
}

func buildDecoderMap(input []string) map[string]int {
	decoder := map[string]int{}
	one, four := "", ""
	for i := range input {
		current := sortString(input[i])
		length := len(current)
		if length == 2 {
			decoder[current] = 1
			one = current
		} else if length == 3 {
			decoder[current] = 7
		} else if length == 4 {
			decoder[current] = 4
			four = current
		} else if length == 7 {
			decoder[current] = 8
		}
	}

	for i := range input {
		current := sortString(input[i])
		length := len(current)

		oneCount, fourCount := 0, 0
		for _, char := range current {
			if strings.ContainsRune(one, char) {
				oneCount++
			}
			if strings.ContainsRune(four, char) {
				fourCount++
			}
		}

		if length == 5 {
			if oneCount == 2 {
				decoder[current] = 3
			} else if fourCount == 3 {
				decoder[current] = 5
			} else {
				decoder[current] = 2
			}
		} else if length == 6 {
			if oneCount == 1 {
				decoder[current] = 6
			} else if fourCount == 4 {
				decoder[current] = 9
			} else {
				decoder[current] = 0
			}
		}
	}
	
	return decoder
}

func sortString(input string) string {
	inputAsSlice := []string{}
	for i := range input {
		inputAsSlice = append(inputAsSlice, input[i:i+1])
	}
	sort.Strings(inputAsSlice)
	return strings.Join(inputAsSlice, "")
}
