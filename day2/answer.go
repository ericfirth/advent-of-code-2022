package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// "fmt"
// "log"
// "os"
// "sort"
// "strconv"

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
		fmt.Println("Line: ", line)
		oponnent_choice, result := getOpponentAndResult(line)
		fmt.Println("Oponnent chose ", oponnent_choice, " and result was ", result)
		my_choice := myShape(oponnent_choice, result)
		fmt.Println("My choice: ", my_choice)
		score = score + shape_score[my_choice] + result_score[result]
	}

	fmt.Println("Score: ", score)
}

func myShape(oponnent Choice, result Result) Choice {
	switch oponnent {
		case "Rock":
			switch result {
				case "Win":
					return "Paper"
				case "Draw":
					return "Rock"
				case "Lose":
					return "Scissors"
				}
		case "Paper":
			switch result {
				case "Win":
					return "Scissors"
				case "Draw":
					return "Paper"
				case "Lose":
					return "Rock"
				}
		case "Scissors":
			switch result {
				case "Win":
					return "Rock"
				case "Draw":
					return "Scissors"
				case "Lose":
					return "Paper"
			}
		}
		return "Rock"
	}


func getOpponentAndResult(line string) (Choice, Result) {
	mapOfChoices := map[string]Choice{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}

	mapOfResults := map[string]Result{
		"X": "Lose",
		"Y": "Draw",
		"Z": "Win",
	}

	results := strings.Split(line, " ")
	choice := mapOfChoices[results[0]]
	result := mapOfResults[results[1]]

	return choice, result
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



// Part 1

// func main() {
// 	lines, err := readLines("input.txt")
// 	if err != nil {
// 			log.Fatalf("readLines: %s", err)
// 	}

// 	shape_score := map[Choice]int{
// 		"Rock": 1,
// 		"Paper": 2,
// 		"Scissors": 3,
// 	}
// 	result_score := map[Result]int{
// 		"Win": 6,
// 		"Draw": 3,
// 		"Lose": 0,
// 	}

// 	score := 0

// 	for _, line := range lines {
// 		choices := getChoices(line)
// 		score = score + shape_score[choices[1]] + result_score[getResult(choices)]
// 	}

// 	fmt.Println("Score: ", score)
// }

// func getChoices(line string) ([2]Choice) {
// 	mapofTypes := map[string]Choice{
// 		"A": "Rock",
// 		"B": "Paper",
// 		"C": "Scissors",
// 		"X": "Rock",
// 		"Y": "Paper",
// 		"Z": "Scissors",
// 	}
// 	split := strings.Split(line, " ")
// 	choices := [2]Choice{ mapofTypes[split[0]], mapofTypes[split[1]] }

// 	return choices;
// }

// func getResult(choices [2]Choice) Result {
// 	switch choices[1] {
// 		case "Rock":
// 			switch choices[0] {
// 				case "Rock":
// 					return "Draw"
// 				case "Paper":
// 					return "Lose"
// 				case "Scissors":
// 					return "Win"
// 			}
// 		case "Paper":
// 			switch choices[0] {
// 				case "Rock":
// 					return "Win"
// 				case "Paper":
// 					return "Draw"
// 				case "Scissors":
// 					return "Lose"
// 			}
// 		case "Scissors":
// 			switch choices[0] {
// 				case "Rock":
// 					return "Lose"
// 				case "Paper":
// 					return "Win"
// 				case "Scissors":
// 					return "Draw"
// 			}
// 		}

//   return "Draw";
// }



// func readLines(path string) ([]string, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 			return nil, err
// 	}
// 	defer file.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 			lines = append(lines, scanner.Text())
// 	}
// 	return lines, scanner.Err()
// }

