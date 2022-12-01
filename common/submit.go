package common

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Submit(year int, day int, level int, answer int) {
	fmt.Printf("Will submit for %d/%d/%d: %d\n", year, day, level, answer)
	payload := bytes.NewReader([]byte(fmt.Sprintf("level=%d&answer=%d", level, answer)))
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)
	req, err := http.NewRequest(http.MethodPost, url, payload)
	Check(err)
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{}
	response, err := client.Do(req)
	Check(err)
	body, err := io.ReadAll(response.Body)
	Check(err)
	print := false
	for _, row := range strings.Split(string(body), "\n") {
		if strings.HasPrefix(row, "</main>") {
			print = false
		} else if strings.HasPrefix(row, "<main>") {
			print = true
		} else if print {
			fmt.Println(row)
		}
	}
}
