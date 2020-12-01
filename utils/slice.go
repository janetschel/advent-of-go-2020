package utils

import "strconv"

func ToIntSlice(slice []string) []int {
	sliceToReturn := []int{}

	for _, current := range slice {
		convertedString, _ := strconv.Atoi(current)
		sliceToReturn = append(sliceToReturn, convertedString)
	}

	return sliceToReturn
}