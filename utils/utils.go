package utils

import (
	"strconv"
)

func ToIntSlice(slice []string) []int {
	sliceToReturn := []int{}

	for _, current := range slice {
		convertedString, err := strconv.Atoi(current)

		if err != nil {
			panic(err)
		}

		sliceToReturn = append(sliceToReturn, convertedString)
	}

	return sliceToReturn
}

func ToInt(str string) int {
	number, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return number
}

func Unpack(slice []string, vars... *string) {
	for i, str := range slice {
		*vars[i] = str
	}
}

func Filter(slice []string, filter func(string) bool) []string {
	retSlice := []string{}

	for _, element := range slice {
		if filter(element) {
			retSlice = append(retSlice, element)
		}
	}

	return retSlice
}

func Count(slice []string) int {
	return len(slice)
}