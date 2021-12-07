package slices

import (
	"fmt"
	"regexp"
	"strings"
)

func Unpack(slice []string, vars ...*string) {
	for i, str := range slice {
		*vars[i] = str
	}
}

func ParseLine(line string, splitOn string, vars ...*string) {
	regex := regexp.MustCompile(splitOn)
	Unpack(regex.Split(line, -1), vars...)
}

func Filter(slice []string, filter func(string) bool) []string {
	retSlice := make([]string, 0)

	for _, element := range slice {
		if filter(element) {
			retSlice = append(retSlice, element)
		}
	}

	return retSlice
}

func Contains(slice []string, word string) bool {
	for _, element := range slice {
		if element == word {
			return true
		}
	}

	return false
}

func Max(input []int) int {
	max := 0
	for _, element := range input {
		if element > max {
			max = element
		}
	}

	return max
}

func Frame(slice []string) []string {
	framed := make([]string, len(slice)+1)
	padding := strings.Repeat(".", len(slice[0])+2)

	framed = append(framed, padding)
	framed[0] = padding

	for i := 1; i < len(slice)+1; i++ {
		framed[i] = fmt.Sprintf(".%s.", slice[i-1])
	}

	return framed
}

func Equals(first []string, second []string) bool {
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func CountCharInSlice(slice []string, char string) int {
	numChars := 0
	for _, element := range slice {
		numChars += strings.Count(element, char)
	}

	return numChars
}
