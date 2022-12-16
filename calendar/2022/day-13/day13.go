package main

import (
	"advent-of-go/utils/files"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

func main() {
	input := files.ReadFile(13, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	packets, sum := parseInput(input), 0
	for i := 0; i < len(packets) - 1; i += 2 {
		if comparePackets(packets[i], packets[i+1]) >= 0 {
			sum += (i / 2) + 1
		}
	}
	return sum
}

func solvePart2(input []string) int {
	divider1, divider2 := "[[2]]", "[[6]]"
	input = append(input, divider1, divider2)
	packets := parseInput(input)

	sort.Slice(packets, func(i int, j int) bool {
		return comparePackets(packets[i], packets[j]) > 0
	})

	decoderKey := 1
	for i, item := range packets {
		formattedItem := fmt.Sprint(item)
		if formattedItem == divider1 || formattedItem == divider2 {
			decoderKey *= i + 1
		}
	}

	return decoderKey
}

// returns 1 if left is larger, -1 if right is larger, 0 if equivalent
func comparePackets(left []any, right []any) int {
	for i := 0; i < len(left); i++ {

		// right list is longer
		if i == len(right) {
			return -1
		}

		// the JSON.Unmarshall function uses floats for all number types although the input is always ints
		integerType := reflect.Float64.String()
		typeA, typeB := reflect.TypeOf(left[i]), reflect.TypeOf(right[i])
		if typeA.Name() == integerType && typeB.Name() == integerType {
			// both integers
			l, r := left[i].(float64), right[i].(float64)
			if l != r {
				if l < r {
					return 1
				}
				return -1
			}
		} else if typeA == typeB {
			// both lists
			nested := comparePackets(left[i].([]any), right[i].([]any))
			if nested != 0 {
				return nested
			}
		} else {
			// one integer, one list
			var nested int
			if typeA.Name() == integerType {
				nested = comparePackets([]any{left[i]}, right[i].([]any))
			} else {
				nested = comparePackets(left[i].([]any), []any{right[i]})
			}
			if nested != 0 {
				return nested
			}
		}
	}

	// left list is longer, packets are ordered
	if len(left) < len(right) {
		return 1
	}

	// reached the end of the list, left and right are equivalent
	return 0
}

func parseInput(input []string) [][]any {
	packets := [][]any{}

	for _, p := range input {
		if p != "" {
			var packet []any
			json.Unmarshal([]byte(p), &packet)
			packets = append(packets, packet)
		}
	}
	
	return packets
}
