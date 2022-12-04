package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	sum := 0
	for _, line := range lines {
		range1, range2 := intoRanges(line)
		if rangesOverlap(range1, range2) {
			sum++
		}
	}
	fmt.Println("Total overlaps: ", sum)
}

func rangesOverlap(range1 []int, range2 []int) bool {
	return range2[0] >= range1[0]  && range2[1] <= range1[1] || range1[0] >= range2[0] && range1[1] <= range2[1]
}

func intoRanges(line string) ([]int, []int) {
	rangeStrings := strings.Split(line, ",")
	return intoRange(rangeStrings[0]), intoRange(rangeStrings[1])
}

func intoRange(rangeString string) []int {
	bounds := strings.Split(rangeString, "-")
	lower, _ := strconv.Atoi(bounds[0])
	upper, _ := strconv.Atoi(bounds[1])
	return []int{lower, upper}
	// s := make([]int, upper-lower+1)
	// for i := range s {
  //   s[i] = i + lower
	// }
	// return s
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
