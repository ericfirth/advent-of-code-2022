package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/maps"
)

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	overlapping_chars := []string{}
	for _, line := range lines {
		// fmt.Println("Line: ", line)
		// firstHalf, secondHalf := getHalves(line)
		// fmt.Println(firstHalf, secondHalf)
		// fmt.Println("common chars", getCommonCharsFromLine(line))
		overlapping_chars = append(overlapping_chars, getCommonCharsFromLine(line)...)
		// fmt.Println(overlapping_chars)
	}

	sum := 0
	for _, char := range overlapping_chars {
		sum += getValueOfChar(char)
	}
	fmt.Println("Sum of overlapping chars: ", sum)
}

func getValueOfChar(char string) int {
	map_of_chars := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}
	return map_of_chars[char]
}

func getCommonCharsFromLine(line string) []string {
	mymap := make(map[string]bool)
	firstHalf, secondHalf := getHalves(line)
	for _, char := range firstHalf {
		fmt.Println("Char: ", string(char), "Second Half: ", secondHalf, "Count: ", strings.Count(secondHalf, string(char)))
		if strings.Count(secondHalf, string(char)) >= 1 {
			mymap[string(char)] = true
		}
	}
	return maps.Keys(mymap)
}

func getHalves(line string) (string, string) {
	half := len(line) / 2
	return line[:half], line[half:]
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
