package req

import (
	"advent-of-go/secrets"
	"fmt"
	"io/ioutil"
	"net/http"
)

func MakeRequest(day int, year int) string {
	url := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Cookie", secrets.Session)

	client := http.Client{}
	response, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}
