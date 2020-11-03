package utils

import (
	"advent-of-go-2020/secrets"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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
