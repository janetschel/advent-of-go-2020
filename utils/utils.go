package utils

import (
	"strconv"
)

func ToIntSlice(slice []string) []int {
	sliceToReturn := []int{}

	for _, current := range slice {
		convertedString, err := strconv.Atoi(current)

		if err != nil {
			panic("Can't convert to int.. Please check the validity your input AND your specified delimiter")
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