package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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
	//scan entire file into a 2d array of single character strings
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Split(line, ""))
	}

	return lines

}

func main() {

	input := readInput("input.txt")
	fmt.Println("pt1:", pt1(input))
	fmt.Println("pt2:", pt2(input))
}

func pt1(input [][]string) int {
	totalAccessibleForklifts := 0

	// Define all 8 directions as [row_offset, col_offset]
	directions := [][2]int{
		{-1, 0},  // up
		{1, 0},   // down
		{0, -1},  // left
		{0, 1},   // right
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, -1},  // down-left
		{1, 1},   // down-right
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "@" {
				adjacentForkliftCount := 0
				for _, dir := range directions {
					neighborRow, neighborCol := i+dir[0], j+dir[1]
					// Check bounds and character
					if neighborRow >= 0 && neighborRow < len(input) && neighborCol >= 0 && neighborCol < len(input[i]) {
						if input[neighborRow][neighborCol] == "@" {
							adjacentForkliftCount++
						}
					}
				}

				if adjacentForkliftCount < 4 {
					totalAccessibleForklifts++
				}
			}
		}
	}
	return totalAccessibleForklifts
}

func pt2(input [][]string) int {
	totalRemovedRolls := 0
	currentRemovedRolls := 0
	shouldKeepRemoving := true

	// Define all 8 directions as [row_offset, col_offset]
	directions := [][2]int{
		{-1, 0},  // up
		{1, 0},   // down
		{0, -1},  // left
		{0, 1},   // right
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, -1},  // down-left
		{1, 1},   // down-right
	}

	for shouldKeepRemoving {
		currentRemovedRolls = 0
		for i := 0; i < len(input); i++ {
			for j := 0; j < len(input[i]); j++ {
				if input[i][j] == "@" {
					adjacentForkliftCount := 0
					for _, dir := range directions {
						neighborRow, neighborCol := i+dir[0], j+dir[1]
						// Check bounds and character
						if neighborRow >= 0 && neighborRow < len(input) && neighborCol >= 0 && neighborCol < len(input[i]) {
							if input[neighborRow][neighborCol] == "@" {
								adjacentForkliftCount++
							}
						}
					}

					if adjacentForkliftCount < 4 {
						//modify the character at the current index to mark it for removal in a future loop
						input[i][j] = "X"
						totalRemovedRolls++
						currentRemovedRolls++
					}
				}
			}
		}

		//change all marked for removal characters to periods so they're skipped in the next iteration.
		for i := 0; i < len(input); i++ {
			for j := 0; j < len(input[i]); j++ {
				if input[i][j] == "X" {
					input[i][j] = ""
				}
			}
		}

		if currentRemovedRolls == 0 {
			shouldKeepRemoving = false
		}
	}

	return totalRemovedRolls
}
