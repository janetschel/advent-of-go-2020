package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"tblue-aoc-2021/utils/req"
)

// Reads content of the input file and returns it in an array, split by the specified delimiter
// If the input file does not exist, it will be created
func ReadFile(day int, year int, delimiter string) []string {
	currentDay := strconv.Itoa(day)

	if len(currentDay) == 1 {
		currentDay = "0" + currentDay
	}

	filePath := fmt.Sprintf("calendar/%v/day-%v/puzzle-input.in", year, currentDay)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		createFile(day, year, filePath)
	} else {
		fmt.Println("INFO: File already exists.. Will not create new one")
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	fileContent := string(file)
	slicedContent := strings.Split(fileContent, delimiter)

	if delimiter == "\n" {
		// fetch utils always adds a new line at the end of a file, which could lead to some problems when parsing it
		// this is why the last row is just removed if the delimiter is a newline

		return slicedContent[:len(slicedContent)-1]
	} else {
		// if the delimiter is not a newline and we split on eg. a comma, the newline will be appended to the last
		// element in the slice which then cannot be converted to an int.
		// this is the reason the last element in the slice is modified (the last char is removed
		// [which is the extra newline]) so it can be worked with

		lastElement := slicedContent[len(slicedContent)-1]
		slicedContent[len(slicedContent)-1] = lastElement[:len(lastElement)-1]

		return slicedContent
	}
}

func createFile(day int, year int, filePath string) {
	puzzleInput := req.MakeRequest(day, year)

	err := ioutil.WriteFile(filePath, []byte(puzzleInput), 0755)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("INFO: File successfully created")
	}
}
