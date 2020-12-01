package utils

import (
	"fmt"
	"strconv"
)

func ToIntSlice(slice []string) []int {
	sliceToReturn := []int{}

	for _, current := range slice {
		fmt.Println(current)
		convertedString, err := strconv.Atoi(current)

		if err != nil {
			panic("Can't convert to int.. Please check the validity your input AND your specified delimiter")
		}

		sliceToReturn = append(sliceToReturn, convertedString)
	}

	return sliceToReturn
}