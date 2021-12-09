package main

import (
	"strconv"
	"strings"
	"tblue-aoc-2021/utils/files"
)

type journalEntry struct {
	signalPatterns, output []string
}

func main() {
	input := files.ReadFile(8, 2021, "\n", false)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	entries := parseInput(input)

	return countEasySymbols(entries)
}

func solvePart2(input []string) int {
	entries := parseInput(input)

	return getOutputSum(entries)
}

func parseInput(input []string) []journalEntry {
	entries := make([]journalEntry, len(input))
	for i, val := range input {
		halves := strings.Split(val, " | ")
		entries[i] = journalEntry{
			signalPatterns: strings.Split(halves[0], " "),
			output:         strings.Split(halves[1], " "),
		}
	}
	return entries
}

func countEasySymbols(entries []journalEntry) int {
	count := 0
	for _, val := range entries {
		for _, o := range val.output {
			segmetCount := len(o)
			if segmetCount == 2 || segmetCount == 4 || segmetCount == 3 || segmetCount == 7 {
				count++
			}
		}
	}
	return count
}

func getOutputSum(entries []journalEntry) int {
	sum := 0
	for _, entry := range entries {
		outputString := ""
		for _, o := range entry.output {
			digit := "2"
			segmetCount := len(o)
			if segmetCount == 2 {
				digit = "1"
			}
			if segmetCount == 4 {
				digit = "4"
			}
			if segmetCount == 3 {
				digit = "7"
			}
			if segmetCount == 7 {
				digit = "8"
			}
			if segmetCount == 6 {
				digit = "6"
				//find 4
				fourVal := ""
				sevenVal := ""
				for _, val := range entry.signalPatterns {
					if len(val) == 4 {
						fourVal = val
					}
					if len(val) == 3 {
						sevenVal = val
					}
				}
				matchingSegment4Count := 0
				matchingSegment7Count := 0
				for _, r := range o {
					if strings.ContainsRune(fourVal, r) {
						matchingSegment4Count++
					}
					if strings.ContainsRune(sevenVal, r) {
						matchingSegment7Count++
					}
				}
				if matchingSegment4Count == 4 {
					digit = "9"
				} else if matchingSegment7Count == 3 {
					digit = "0"
				}

			}
			if segmetCount == 5 {
				//find 4
				fourVal := ""
				sevenVal := ""
				for _, val := range entry.signalPatterns {
					if len(val) == 4 {
						fourVal = val
					}
					if len(val) == 3 {
						sevenVal = val
					}
				}
				matchingSegment4Count := 0
				matchingSegment7Count := 0
				for _, r := range o {
					if strings.ContainsRune(fourVal, r) {
						matchingSegment4Count++
					}
					if strings.ContainsRune(sevenVal, r) {
						matchingSegment7Count++
					}
				}
				if matchingSegment7Count == 3 {
					digit = "3"
				} else if matchingSegment4Count == 3 {
					digit = "5"
				}
			}
			outputString += digit
		}
		output, _ := strconv.Atoi(outputString)
		sum += output
	}
	return sum
}
