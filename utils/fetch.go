package utils

import (
	"advent-of-go-2020/secrets"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateFile(day int, filePath string) {
	puzzleInput := makeRequest(day)

	err := ioutil.WriteFile(filePath, []byte(puzzleInput), 0755)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("INFO: File successfully created")
	}
}

func makeRequest(day int) string {
	url := fmt.Sprintf("https://adventofcode.com/2020/day/%v/input", day)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Cookie", secrets.Session)

	client := http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
