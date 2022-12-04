package knothash

import (
	"fmt"
	"strconv"
	"strings"
)

// HashInt runs a single round of the knot hash algorithm, with the input as ints
func HashInt(input string) string {
	return hash(input, parseAsInts, 256, 1, formatInt)
}

// HashHex runs a full 256 rounds of the knot hash algorithm, with the inputs as bytes
func HashHex(input string) string {
	return hash(input, parseAsBytes, 256, 64, formatHex)
}

func hash(input string, parse func(string) []int, listSize int, rounds int, format func([]int) string) string {
	list := makeRange(0, listSize - 1)
	lengths := parse(input)

	currentPosition := 0
	skipSize := 0

	for i := 0; i < rounds; i++ {
		lengthsCopy := make([]int, len(lengths))
		copy(lengthsCopy, lengths)
		currentPosition, skipSize = runRound(list, lengthsCopy, currentPosition, skipSize)
	}

	return format(list)
}

func formatInt(arr []int) string {
	return fmt.Sprintf("%d", arr[0] * arr[1])
}

func formatHex(arr []int) string {
	hexString := ""
	for i := 0; i < len(arr) / 16; i++  {
		blockValue := 0
		for j := 0; j < 16; j++ {
			blockValue ^= arr[(16 * i) + j]
		}

		hexString += fmt.Sprintf("%02x", blockValue)
	}

	return hexString
}

func runRound(list []int, lengths []int, currentPosition int, skipSize int) (int, int) {
	for _, length := range lengths {
		for i := 0; i < length/2; i++ {
			firstIndex, lastIndex := currentPosition + i, currentPosition + (length - i) - 1
			first, last := elementAtCircular(list, firstIndex), elementAtCircular(list, lastIndex)
			setCircular(list, first, lastIndex)
			setCircular(list, last, firstIndex)
		}
		currentPosition += skipSize + length
		skipSize++
	}
	return currentPosition, skipSize
}

func parseAsInts(input string) []int {
	numStrs := strings.Split(input, ",")
	lengths := make([]int, len(numStrs))
	for i, num := range numStrs {
		value, _ := strconv.Atoi(num)
		lengths[i] = value
	}
	return lengths
}

func parseAsBytes(input string) []int {
	lengths := make([]int, len(input))
	for i, num := range input {
		lengths[i] = int(num)
	}
	return append(lengths, []int{17, 31, 73, 47, 23}...)
}

func elementAtCircular(arr []int, index int) int {
	return(arr[index % len(arr)])
}

func setCircular(arr []int, value int, index int) {
	arr[index % len(arr)] = value
}

func makeRange(from int, to int) []int {
	arr := make([]int, (to - from) + 1)
	for i := 0; i < len(arr); i++ {
		arr[i] = from + i
	}
	return arr
}