package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Reads content of the input file and returns it in an array, split by the specified delimiter
// If the input file does not exist, it will be created
func ReadInputFile(day int, delimiter string) []string {
	currentDay := strconv.Itoa(day)

	if len(currentDay) == 1 {
		currentDay = "0" + currentDay
	}

	filePath := fmt.Sprintf("calendar/day-%v/puzzle-input.txt", currentDay)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		CreateFile(day, filePath)
	} else {
		log.Fatal("INFO: File already exists.. Will not create new one")
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(file)
	slicedContent := strings.Split(fileContent, delimiter)

	return slicedContent
}
