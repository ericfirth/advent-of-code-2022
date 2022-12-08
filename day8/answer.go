package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	answerOne("sample_input.txt")
	answerOne("input.txt")
	answerTwo("sample_input.txt")
	answerTwo("input.txt")
}

func answerTwo(filename string) {
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
	grid := asGrid(lines)

	var highestScenicScore int
	for y, row := range grid {
		for x := range row {
			score := calcScenicScore(grid, x, y)
			if score > highestScenicScore {
				highestScenicScore = score
			}
		}
	}
	fmt.Println("Highest scenic score:", highestScenicScore)
}

func calcScenicScore(grid [][]int, x, y int) int {
	num := grid[y][x]
	left := 0
	right := 0
	top := 0
	bottom := 0

	// Check left
	for i := x - 1; i >= 0; i-- {
		otherNum := grid[y][i]
		left++
		if otherNum >= num {
			break
		}
	}

	// Check right
	for _, otherNum := range grid[y][x+1:] {
		right++
		if otherNum >= num {
			break
		}
	}

	// Check top
	for i := y - 1; i >= 0; i-- {
		otherRow := grid[i]
		otherNum := otherRow[x]
		top++
		if otherNum >= num {
			break
		}
	}

	// Check bottom
	for _, otherRow := range grid[y+1:] {
		otherNum := otherRow[x]
		bottom++
		if otherNum >= num {
			break
		}
	}

	scenicScore := left * right * top * bottom
	return scenicScore
}

func answerOne(filename string) {
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
	
	numVisible := countVisible(asGrid(lines))
	fmt.Println("Number of visible:", numVisible)
}

func asGrid(lines []string) ([][]int) {
	grid := make([][]int, len(lines))
	for lineIndex, line := range lines {
		row := make([]int, len(line))
		for charIndex, char := range strings.Split(line, "") {
			intChar, _ := strconv.Atoi(char)
			row[charIndex] = intChar
		}
		grid[lineIndex] = row
	}
	return grid
}

func countVisible(grid [][]int) int {
	count := 0

	for x, row := range grid {
		for y := range row {
			if x == 0 || y == 0 || x == len(row) - 1 || y == len(grid) - 1 {
				count++
			} else {
				if isVisible(grid, x, y) {
					count++
				}
			}
		}
	}
	return count
}

func isVisible(grid [][]int, x, y int) bool {
	num := grid[y][x]
	visibleFromLeft := true
	visibleFromRight := true
	visibleFromTop := true
	visibleFromBottom := true

	// Check left
	for _, otherNum := range grid[y][:x] {
		if otherNum >= num {
			visibleFromLeft = false
			break
		}
	}

	// Check right
	for _, otherNum := range grid[y][x+1:] {
		if otherNum >= num {
			visibleFromRight = false
			break
		}
	}

	// Check top
	for _, otherRow := range grid[:y] {
		if otherRow[x] >= num {
			visibleFromTop = false
			break
		}
	}

	// Check bottom
	for _, otherRow := range grid[y+1:] {
		if otherRow[x] >= num {
			visibleFromBottom = false
			break
		}
	}

	isVisible := visibleFromRight || visibleFromLeft || visibleFromTop || visibleFromBottom
	return isVisible
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
