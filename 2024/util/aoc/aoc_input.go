package aoc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetInput(day int, useTestData bool, testData string) string {
	if useTestData {
		// Leading newline and random tabs have killed me every day

		testData = strings.ReplaceAll(testData, "\t", "")

		if testData[0] == '\n' {
			testData = testData[1:]
		}

		return testData
	}

	sessionId, err := os.ReadFile("../session_id")

	if err != nil {
		panic(err)
	}

	cachedFilename := fmt.Sprintf("./input_%v", day)

	cachedInput, err := os.ReadFile(cachedFilename)

	if err == nil && len(cachedInput) > 0 {
		return string(cachedInput)
	}

	fmt.Println("Warning: read of uncached puzzle input triggering HTTP request")

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2024/day/%v/input", day), nil)

	if err != nil {
		panic(err)
	}

	req.Header.Add("Cookie", "session="+string(sessionId))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 200 {
		output := raw[:len(raw)-1] // strip newline
		err := os.WriteFile(cachedFilename, output, 0666)

		if err != nil {
			panic(err)
		}

		return string(output)
	}

	return string(raw)
}
