package common

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func ReadFile(year int, day int, filePath string) string {
	currentDay := strconv.Itoa(day)
	if len(currentDay) == 1 {
		currentDay = "0" + currentDay
	}
	if filePath == "" {
		filePath = fmt.Sprintf("%d/%v/puzzle.txt", year, currentDay)
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		createFile(year, day, filePath)
	} else {
		fmt.Println("INFO: File already exists.. Will not create new one")
	}
	file, err := os.ReadFile(filePath)
	Check(err)
	return string(file)
}

func createFile(year int, day int, filePath string) {
	puzzleInput := makeRequest(year, day)
	err := os.WriteFile(filePath, []byte(puzzleInput), 0755)
	Check(err)
	fmt.Println("INFO: File successfully created")
}

func makeRequest(year int, day int) string {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	Check(err)
	req.Header.Set("Cookie", cookie)
	client := http.Client{}
	response, err := client.Do(req)
	Check(err)
	body, err := io.ReadAll(response.Body)
	Check(err)
	return strings.TrimSuffix(string(body), "\n")
}
