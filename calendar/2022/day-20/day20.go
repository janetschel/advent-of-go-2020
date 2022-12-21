package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/slices"
	"strconv"
)

type fileContents struct {
	value int
	originalIndex int
}

func main() {
	input := files.ReadFile(20, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	file := parseInput(input)
	sequence := mix(file, 1)
	zeroIndex := slices.IndexOfInt(0, sequence)
	return sequence[(zeroIndex + 1000) % len(sequence)] +
		sequence[(zeroIndex + 2000) % len(sequence)] +
		sequence[(zeroIndex + 3000) % len(sequence)]
}

func solvePart2(input []string) int {
	file := parseInput(input)
	for i := range file {
		file[i] *= 811589153
	}
	sequence := mix(file, 10)
	zeroIndex := slices.IndexOfInt(0, sequence)
	return sequence[(zeroIndex + 1000) % len(sequence)] +
		sequence[(zeroIndex + 2000) % len(sequence)] +
		sequence[(zeroIndex + 3000) % len(sequence)]
}

func mix(file []int, times int) []int {
	sequence := make([]fileContents, len(file))
	for i, value := range file {
		sequence[i] = fileContents{ value: value, originalIndex: i }
	}
	for n := 0; n < times; n++ {
		for i, value := range file {
			if value != 0 {
				currentIndex := findByOriginalIndex(sequence, i)
				destination := (currentIndex + value) % (len(sequence) - 1)	
				if destination <= 0 {
					destination = len(sequence) - maths.Abs(destination) - 1
				}
				newArr := []fileContents{}
				if currentIndex < destination {
					for i := 0; i < currentIndex; i++ {
						newArr = append(newArr, sequence[i])
					}
					newArr = append(newArr, sequence[currentIndex+1:destination+1]...)
					newArr = append(newArr, sequence[currentIndex])
					newArr = append(newArr, sequence[destination+1:]...)
				} else {
					for i := 0; i < destination; i++ {
						newArr = append(newArr, sequence[i])
					}
					newArr = append(newArr, sequence[currentIndex])
					newArr = append(newArr, sequence[destination:currentIndex]...)
					newArr = append(newArr, sequence[currentIndex+1:]...)
				}
				sequence = newArr
			}
		}
	}
	return toIntSlice(sequence)
 }

func parseInput(input []string) []int {
	numbers := make([]int, len(input))
	for i, str := range input {
		value, _ := strconv.Atoi(str)
		numbers[i] = value
	}
	return numbers
}

func findByOriginalIndex(contents []fileContents, originalIndex int) int {
	for i, c := range contents {
		if c.originalIndex == originalIndex {
			return i
		}
	}
	return -1
}

func toIntSlice(contents []fileContents) []int {
	intSlice := make([]int, len(contents))
	for i, c := range contents {
		intSlice[i] = c.value
	}
	return intSlice
}
