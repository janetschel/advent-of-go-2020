package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(16, 2019, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) string {
	signal := input[0]
	for i := 0; i < 100; i++ {
		signal = fft(signal)
	}
	return signal[:8]
}

func solvePart2(input []string) string {
	offset, _ := strconv.Atoi(input[0][:7])
	// brute force won't work here
	// there must be a pattern that can be applied
	signal := strings.Repeat(input[0], 10000)
	signalMap := map[string]string{}

	for i := 0; i < 100; i++ {
		newSignal, cached := signalMap[signal]
		if !cached {
			newSignal = fft(signal)
			signalMap[signal] = newSignal
		}
		signal = newSignal
	}

	return signal[offset:offset+8]
}

func fft(input string) string {
	pattern := []int{ 0, 1, 0, -1 }
	output := ""
	for i := range input {
		sum := 0
		for j, char := range input {
			sum += int(char - '0') * pattern[(j+1)/(i+1)%4]
		}
		output += fmt.Sprint(maths.Abs(sum % 10))
	}
	return output
}
