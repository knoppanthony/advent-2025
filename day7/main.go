package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type FoodRange struct {
	start int
	end   int
}

func readInput(filename string) [][]string {

	// 1. Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err) // Use log.Fatalf for concise error handling
	}
	// Ensure the file is closed after the function exits
	defer file.Close()

	// 2. Create a scanner
	scanner := bufio.NewScanner(file)

	var lines [][]string
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		//just read everything into 2d array
		lines = append(lines, strings.Split(line, ""))
	}
	return lines
}

func main() {

	input := readInput("input.txt")
	fmt.Println(input)
	fmt.Println("pt1:", pt1(input))
	fmt.Println("pt2:", pt2(input))
}

func pt1(input [][]string) int {
	//We have a 2d array of characters. S ^ .
	//S is the starting point. We need to implement a game loop that starts by replacing the character directly south of S with a |
	//Every loop the | continues down the path, replacing any . with |
	//If we hit a ^ then instead of replacing, we put a | on the left and right side of the ^
	//Once we have multiple | we need to ensure they also progress downwards every loop
	// If multiple ^ in one line cause | to overlap, we only create one |
	//return the number of | that are in the final state
	splitCounter := 0

	//loop through everything
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {

			if input[i][j] == "S" {
				//Move onto next row and replace the character directly south of S with a |
				input[i+1][j] = "|"
				continue
			}
			//if were in the last row just break
			if i == len(input)-1 {
				break
			}

			//After the first row, we always want to be working one row down, meaning we can loop through the current row and then analyze the row below it
			if input[i][j] == "|" {
				//check whats below the laser
				if input[i+1][j] == "." {
					//if its a . then replace it with a |
					input[i+1][j] = "|"
				}
				//if its a ^ then put a | on the left and right side of the ^
				if input[i+1][j] == "^" {
					//count this as one split event
					splitCounter++
					//place the split beams on left and right
					if input[i+1][j-1] != "|" {
						input[i+1][j-1] = "|"
					}
					if input[i+1][j+1] != "|" {
						input[i+1][j+1] = "|"
					}
				}
			}
		}
	}
	return splitCounter
}

func pt2(input [][]string) int {
	// Create a deep copy of the input to avoid modifying the original
	grid := make([][]string, len(input))
	for i := range input {
		grid[i] = make([]string, len(input[i]))
		copy(grid[i], input[i])
	}

	// Find the starting position (S)
	var startRow, startCol int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "S" {
				startRow = i
				startCol = j
				break
			}
		}
	}

	// Initialize memoization map
	memo := make(map[string]int)

	// Start tracing from the position directly below S
	return traceTimeline(grid, startRow+1, startCol, memo)
}

// traceTimeline recursively follows a particle's path and counts all possible timelines
func traceTimeline(grid [][]string, row, col int, memo map[string]int) int {
	// Base case: if we've reached the bottom of the grid, this is one complete timeline
	if row >= len(grid) {
		return 1
	}

	// Base case: if we're out of bounds horizontally, this timeline ends
	if col < 0 || col >= len(grid[row]) {
		return 0
	}

	// Check memoization cache
	key := fmt.Sprintf("%d,%d", row, col)
	if val, ok := memo[key]; ok {
		return val
	}

	currentCell := grid[row][col]
	var result int

	// If we hit a splitter (^), the timeline splits into two
	switch currentCell {
	case "^":
		// Timeline splits: one goes left, one goes right
		leftTimelines := traceTimeline(grid, row+1, col-1, memo)
		rightTimelines := traceTimeline(grid, row+1, col+1, memo)
		result = leftTimelines + rightTimelines
	case ".", "|":
		// If we hit a dot (.) or pipe (|), continue straight down
		result = traceTimeline(grid, row+1, col, memo)
	default:
		// If we hit something else (like another S or out of bounds), this timeline ends
		result = 0
	}

	// Cache and return result
	memo[key] = result
	return result
}
