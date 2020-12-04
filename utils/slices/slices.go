package slices

import "regexp"

func Unpack(slice []string, vars... *string) {
	for i, str := range slice {
		*vars[i] = str
	}
}

func ParseLine(line string, splitOn string, vars... *string) {
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
