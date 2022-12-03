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

	overlappingChars := []string{}
	for _, line := range lines {
		overlappingChars = append(overlappingChars, getCommonCharsFromLine(line)...)
	}

	sumForAnswerOne := 0
	for _, char := range overlappingChars {
		sumForAnswerOne += getValueOfChar(char)
	}
	fmt.Println("Sum of overlapping chars: ", sumForAnswerOne)

	groupsOfThree := getGroupsOfThree(lines)
	fmt.Println("num groups should be: ", len(lines)/3, " and is: ", len(groupsOfThree))
	sumForAnswerTwo := 0
	for _, group := range groupsOfThree {
		sumForAnswerTwo += getValueOfChar(getCommonCharFromGroupOfThree(group))
	}

	fmt.Println("Sum of common chars: ", sumForAnswerTwo)
}

func getCommonCharFromGroupOfThree(group [3]string) string {
	var commonChar string
	for _, char := range group[0] {
		if strings.Count(group[1], string(char)) >= 1 && strings.Count(group[2], string(char)) >= 1 {
			commonChar = string(char)
		}
	}
	return commonChar
}

func getGroupsOfThree(lines []string) [][3]string {
	var groupsOfThree [][3]string
	var group [3]string
	fmt.Println(lines)
	for i, line := range lines {
		group[i%3] = line
		if i%3 == 2 {
			groupsOfThree = append(groupsOfThree, group)
			group = [3]string{}
		}
	}
	return groupsOfThree
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
