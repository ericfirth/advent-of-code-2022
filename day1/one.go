package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)


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

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
			log.Fatalf("readLines: %s", err)
	}

	calories := []int{}
	currentElf := 0
	currentCalories := 0
	
	for _, line := range lines {
		if line == "" {
			calories = append(calories, currentCalories)

			currentCalories = 0
			currentElf = currentElf + 1
			continue
		} else {
			calories, err := strconv.Atoi(line)
			if err == nil {
				currentCalories = currentCalories + calories
			}
			continue
		}
	}

	sort.Ints(calories)

	topCalories := calories[len(calories)-1]
	topThreeCalories :=  topCalories + calories[len(calories)-2] + calories[len(calories)-3]

	fmt.Println("Top calories: ", topCalories)
	fmt.Println("Total of top three calories ", topThreeCalories)
}
