package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	numToMove int
	from      int
	to        int
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	stacks, instructions := getStacksAndInstructions(lines)
	fmt.Println("stacks", stacks)
	fmt.Println("instructions", instructions)
	for _, instruction := range instructions {
		fmt.Println("before:", stacks, "instruction:", instruction)
		stacks = performInstruction(stacks, instruction)
		fmt.Println("after:", stacks)
	}

	var result string
	for i := 1; i <= len(stacks); i++ {
		fmt.Println("stack", i, ":", stacks[i])
		result += stacks[i][0]
	}

	fmt.Println("Result of answer one:", result)
}

func performInstruction(stacks map[int][]string, instruction Instruction) map[int][]string {
	// copy the stack
	newStacks := make(map[int][]string)
	for k, v := range stacks {
		newStacks[k] = v
	}

	for i := 0; i < instruction.numToMove; i++ {
		char := newStacks[instruction.from][0]
		fmt.Println("char", char, "old from: ", newStacks[instruction.from], "old to: ", newStacks[instruction.to])
		newStacks[instruction.from] = newStacks[instruction.from][1:len(newStacks[instruction.from])]
		fmt.Println("new from: ", newStacks[instruction.from])
		newStacks[instruction.to] = append([]string{char}, newStacks[instruction.to]...)
		fmt.Println("new to: ", newStacks[instruction.to])
	}

	return newStacks
}

// func reverseArray(initial []string) []string {
// 	result := make([]string, len(initial))
// 	for i, v := range initial {
// 		result[len(initial)-1-i] = v
// 	}
// 	return result
// }

func makeInstructions(input string) []Instruction {
	result := []Instruction{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(line, -1)
		numToMove, _ := strconv.Atoi(nums[0])
		from, _ := strconv.Atoi(nums[1])
		to, _ := strconv.Atoi(nums[2])
		instruction := Instruction{numToMove, from, to}
		result = append(result, instruction)
	}
	return result
}

func makeStacks(input string) map[int][]string {
	result := make(map[int][]string)
	inputAsLines := strings.Split(input, "\n")
	stackNumbers := strings.Fields(inputAsLines[len(inputAsLines)-1])
	var highestStackNumber int
	for _, stackNumber := range stackNumbers {
		stackInt, _ := strconv.Atoi(stackNumber)
		result[stackInt] = []string{}
		highestStackNumber = stackInt
	}

	for _, line := range inputAsLines[:len(inputAsLines)-1] {
		chars := strings.Split(line, "")
		for i := 1; i <= highestStackNumber; i++ {
			spot := 1 + ((i - 1) * 4)
			char := chars[spot]
			if char != " " {
				result[i] = append(result[i], char)
			}
		}
	}

	return result
}

func getStacksAndInstructions(input []string) (map[int][]string, []Instruction) {
	output := strings.Join(input[:], "\n")
	outputAsArray := strings.Split(output, "\n\n")

	return makeStacks(outputAsArray[0]), makeInstructions(outputAsArray[1])
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
