package main

import (
	"strconv"
	"tblue-aoc-2021/utils/files"
)

func main() {
	input := files.ReadFile(03, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int64 {
	countOfOnes := make([]int, len(input[0]))
	maxOnes := make([]rune, len(countOfOnes))
	for i := range countOfOnes {
		countOfOnes[i] = 0
		maxOnes[i] = '1'
	}
	maxOnesValue, _ := strconv.ParseInt(string(maxOnes), 2, 64)

	for _, binNum := range input {
		for j, runeChar := range binNum {
			if runeChar == '1' {
				countOfOnes[j] += 1
			}
		}
	}

	neededToBeMax := float64(len(input)) / float64(2)
	gamma := make([]rune, len(countOfOnes))

	for i, count := range countOfOnes {
		if float64(count) > neededToBeMax {
			gamma[i] = '1'
		} else {
			gamma[i] = '0'
		}
	}

	gammaValue, _ := strconv.ParseInt(string(gamma), 2, 64)

	return gammaValue * (gammaValue ^ maxOnesValue)
}

func solvePart2(input []string) int64 {
	values := input
	oxyIndex := 0

	for len(values) > 1 {
		zerosAtIndex := []string{}
		onesAtIndex := []string{}
		for _, val := range values {
			if val[oxyIndex] == '1' {
				onesAtIndex = append(onesAtIndex, val)
			} else {
				zerosAtIndex = append(zerosAtIndex, val)
			}
		}
		if len(onesAtIndex) >= len(zerosAtIndex) {
			values = onesAtIndex
		} else {
			values = zerosAtIndex
		}
		oxyIndex++
	}
	oxygenNumber, _ := strconv.ParseInt(string(values[0]), 2, 64)

	values = input
	co2Index := 0

	for len(values) > 1 {
		zerosAtIndex := []string{}
		onesAtIndex := []string{}
		for _, val := range values {
			if val[co2Index] == '1' {
				onesAtIndex = append(onesAtIndex, val)
			} else {
				zerosAtIndex = append(zerosAtIndex, val)
			}
		}
		if len(zerosAtIndex) <= len(onesAtIndex) {
			values = zerosAtIndex
		} else {
			values = onesAtIndex
		}
		co2Index++
	}
	co2Number, _ := strconv.ParseInt(string(values[0]), 2, 64)

	return oxygenNumber * co2Number
}
