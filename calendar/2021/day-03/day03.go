package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(03, 2021, "\n")
	bitColumns := stringsFromBitColumns(input)
	println(solvePart1(input, bitColumns))
	println(solvePart2(input, bitColumns))
}

func solvePart1(input []string, bitColumns []string) int {
	bits := len(input[0])
	majorityCount := len(input) / 2
	gamma := ""
	epsilon := ""
	for i := 0; i < bits; i++ {
		if strings.Count(bitColumns[i], "0") > majorityCount {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaDec, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonDec, _ := strconv.ParseInt(epsilon, 2, 64)
	return int(gammaDec * epsilonDec)
}

func solvePart2(input []string, bitColumns []string) int {
	possibleOxygenValues, possibleCo2Values := make([]string, len(input)), make([]string, len(input))
	copy(possibleOxygenValues, input)
	copy(possibleCo2Values, input)

	bits := len(input[0])

	for bit := 0; bit < bits && (len(possibleOxygenValues) > 1 || len(possibleCo2Values) > 1); bit++ {
		if len(possibleOxygenValues) > 1 {
			o2ZeroCount := 0
			for _, o2Candidate := range possibleOxygenValues {
				if o2Candidate[bit:bit+1] == "0" {
					o2ZeroCount++
				}
			}

			possibleOxygenValues = slices.Filter(possibleOxygenValues, func(candidate string) bool {
				currentBit := candidate[bit : bit+1]
				if currentBit == "0" && o2ZeroCount > (len(possibleOxygenValues)/2) {
					return true
				} else if currentBit == "1" && o2ZeroCount <= (len(possibleOxygenValues)/2) {
					return true
				}
				return false
			})
		}

		if len(possibleCo2Values) > 1 {
			co2ZeroCount := 0
			for _, co2Candidate := range possibleCo2Values {
				if co2Candidate[bit:bit+1] == "0" {
					co2ZeroCount++
				}
			}

			possibleCo2Values = slices.Filter(possibleCo2Values, func(candidate string) bool {
				currentBit := candidate[bit : bit+1]
				if currentBit == "1" && co2ZeroCount > (len(possibleCo2Values)/2) {
					return true
				} else if currentBit == "0" && co2ZeroCount <= (len(possibleCo2Values)/2) {
					return true
				}
				return false
			})
		}
	}

	oxygenDec, _ := strconv.ParseInt(possibleOxygenValues[0], 2, 64)
	co2Dec, _ := strconv.ParseInt(possibleCo2Values[0], 2, 64)

	return int(oxygenDec * co2Dec)
}

func stringsFromBitColumns(input []string) []string {
	bits := len(input[0])
	bitColumns := make([]string, bits)

	for i := 0; i < bits; i++ {
		for _, binary := range input {
			bitColumns[i] += binary[i : i+1]
		}
	}

	return bitColumns
}
