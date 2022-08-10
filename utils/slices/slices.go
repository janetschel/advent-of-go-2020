package slices

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func Unpack(slice []string, vars ...*string) {
	for i, str := range slice {
		*vars[i] = str
	}
}

func ParseLine(line string, splitOn string, vars ...*string) {
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

func ContainsGeneric(slice interface{}, item interface{}) bool {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		panic("Slice is not a slice")
	}

	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func Max(input []int) int {
	max := 0
	for _, element := range input {
		if element > max {
			max = element
		}
	}

	return max
}

func Min(input []int) int {
	min := 0
	for _, element := range input {
		if element < min {
			min = element
		}
	}

	return min
}

func Frame(slice []string) []string {
	framed := make([]string, len(slice)+1)
	padding := strings.Repeat(".", len(slice[0])+2)

	framed = append(framed, padding)
	framed[0] = padding

	for i := 1; i < len(slice)+1; i++ {
		framed[i] = fmt.Sprintf(".%s.", slice[i-1])
	}

	return framed
}

func Equals(first []string, second []string) bool {
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func CountCharInSlice(slice []string, char string) int {
	numChars := 0
	for _, element := range slice {
		numChars += strings.Count(element, char)
	}

	return numChars
}

func Sum(slice []int) int {
	sum := 0
	for i := range slice {
		sum += slice[i]
	}
	return sum
}

func Swap(slice interface{}, i int, j int) {
	if reflect.TypeOf(slice).Kind() == reflect.Slice {
		reflect.Swapper(slice)(i, j)
	}
}

func Fill(value int, count int, arr *[]int) {
	for i := 0; i <= count; i++ {
		*arr = append(*arr, value)
	}
}

func GeneratePermutations(items []int) [][]int {
	length := len(items)

	initial, itemsCopy := make([]int, length), make([]int, length)
	copy(initial, items)
	copy(itemsCopy, items)

	permutations := [][]int{initial}

	indexes := make([]int, length)

	i := 0
	for i < length {
		if indexes[i] < i {
			if i%2 == 0 {
				Swap(itemsCopy, 0, i)
			} else {
				Swap(itemsCopy, indexes[i], i)
			}
			permutation := make([]int, length)
			copy(permutation, itemsCopy)
			permutations = append(permutations, permutation)
			indexes[i] = indexes[i] + 1
			i = 0
		} else {
			indexes[i] = 0
			i++
		}
	}

	return permutations
}

func GenerateCombinationsLengthN(items []int, n int) [][]int {
	length := len(items)
	itemsCopy := make([]int, length)
	copy(itemsCopy, items)

	if length == 0 || n > length || n == 0 {
		return [][]int{{}}
	} else if n == length {
		initial := make([]int, length)
		copy(initial, itemsCopy)
		return [][]int{initial}
	}

	if n == length {
		combinations := [][]int{}
		for _, element := range itemsCopy {
			combinations = append(combinations, []int{element})
			return combinations
		}
	}

	first := itemsCopy[0]
	nMinusOneCombinations := GenerateCombinationsLengthN(itemsCopy[1:], n-1)
	for i := range nMinusOneCombinations {
		nMinusOneCombinations[i] = append([]int{first}, nMinusOneCombinations[i]...)
	}
	return append(nMinusOneCombinations, GenerateCombinationsLengthN(itemsCopy[1:], n)...)
}

func GenerateAllCombinations(items []int) [][]int {
	combinations := [][]int{}
	for n := 0; n <= len(items); n++ {
		combinations = append(combinations, GenerateCombinationsLengthN(items, n)...)
	}
	return combinations
}

// IndexOf returns the index of the selected item or -1 if not present
func IndexOf(item string, slice []string) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}
