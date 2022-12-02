package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "fmt"
	// "log"
	// "os"
	// "sort"
	// "strconv"
)

type Choice string
func (c Choice) Valid() bool {
    switch (c){
        case "Rock", "Paper", "Scissors":
            return true
       default:
           return false
    }
}

type Result string
func (r Result) Valid() bool {
    switch (r){
        case "Win", "Lose", "Draw":
            return true
       default:
           return false
    }
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
			log.Fatalf("readLines: %s", err)
	}

	shape_score := map[Choice]int{
		"Rock": 1,
		"Paper": 2,
		"Scissors": 3,
	}
	result_score := map[Result]int{
		"Win": 6,
		"Draw": 3,
		"Lose": 0,
	}

	score := 0

	for _, line := range lines {
		choices := getChoices(line)
		fmt.Println("Line", line)
		fmt.Println("choices: ", choices)
		fmt.Println("result: ", getResult(choices))
		fmt.Println("Shape Score: ", shape_score[choices[1]])
		fmt.Println("Result Score: ", result_score[getResult(choices)])

		score = score + shape_score[choices[1]] + result_score[getResult(choices)]
	}

	fmt.Println("Score: ", score)
}

func getChoices(line string) ([2]Choice) {
	mapofTypes := map[string]Choice{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}
	split := strings.Split(line, " ")
	choices := [2]Choice{ mapofTypes[split[0]], mapofTypes[split[1]] }

	return choices;
}

func getResult(choices [2]Choice) Result {
	switch choices[1] {
		case "Rock":
			switch choices[0] {
				case "Rock":
					return "Draw"
				case "Paper":
					return "Lose"
				case "Scissors":
					return "Win"
			}
		case "Paper":
			switch choices[0] {
				case "Rock":
					return "Win"
				case "Paper":
					return "Draw"
				case "Scissors":
					return "Lose"
			}
		case "Scissors":
			switch choices[0] {
				case "Rock":
					return "Lose"
				case "Paper":
					return "Win"
				case "Scissors":
					return "Draw"
			}
		}

  return "Draw";
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

