package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main () {
	answer("sample_input.txt", "TEST RUN", 4)
	answer("input.txt", "ACTUAL RUN", 4)

	answer("sample_input.txt", "TEST RUN", 14)
	answer("input.txt", "ACTUAL RUN", 14)
}

func answer(file, beginningMessage string, numUniq int) {
	lines, err := readLines(file)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, line := range lines {
		fmt.Println(beginningMessage)
		start, end, chars := findMarker(line, numUniq)
		fmt.Println("start:", start, "end:", end, "chars:", chars)
		fmt.Println("Answer:", end, "line:", line)
	}
}

func findMarker(line string, numUniqueChars int) (int, int, string) {
	var start int
	var end int
	var chars string
	for i, _ := range line {
		nextChars := line[i:i+numUniqueChars]
		if getUniqChars(nextChars, numUniqueChars) {
			start = i
			end = i + numUniqueChars
			chars = nextChars
			break
		}
	}
	return start, end, chars
}

func getUniqChars (chars string, numChars int) bool {
	charsMap := make(map[string]bool)
	for _, char := range chars {
		charsMap[string(char)] = true
	}

	return len(charsMap) == numChars
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
