package main

import (
	"advent-of-go-2020/utils/files"
	"errors"
	"fmt"
	"go/types"
)

func main() {
	input := files.ReadFile(5, "\n")
	id, err := leftover(input)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Leftover seat has id %d", id)
}

func leftover(input []string) (int, error) {
	ids := make(map[int]types.Nil)

	for _, curr := range input {
		ids[solve(curr)] = types.Nil{}
	}

	for i := 7; i < 908; i++ {
		if _, contains := ids[i]; !contains {
			return i, nil
		}
	}

	return 0, errors.New("no seats leftover")
}

