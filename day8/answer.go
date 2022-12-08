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
	// answerTwo("sample_input.txt")
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
