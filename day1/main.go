package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readInput(filename string) []string {

	var lines []string
	// 1. Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err) // Use log.Fatalf for concise error handling
	}
	// Ensure the file is closed after the function exits
	defer file.Close()

	// 2. Create a scanner
	scanner := bufio.NewScanner(file)

	// 3. Iterate through the file line by line
	for scanner.Scan() {
		line := scanner.Text() // Get the current line as a string (without the newline character)
		lines = append(lines, line)
	}

	// 4. Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %s", err)
	}

	return lines
}

func main() {

	input := readInput("input.txt")
	fmt.Println(pt1(input))
	fmt.Println(pt2(input))
	fmt.Println(pt2bruteforce(input))
}

func pt2bruteforce(input []string) int {
	currentPos := 50
	numZeroPasses := 0

	for _, line := range input {
		re := regexp.MustCompile(`^([LR])(\d+)$`)
		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			log.Fatalf("invalid line format: %s", line)
		}
		direction := matches[1]
		distance, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalf("failed to parse distance: %s", err)
		}
		// Now 'direction' holds "L" or "R", and 'distance' holds the integer value.

		switch direction {
		case "R":
			//loop through distance
			for i := 0; i < distance; i++ {
				currentPos++
				if currentPos == 100 {
					numZeroPasses++
				}
				if currentPos > 99 {
					currentPos = 0
				}
			}
		case "L":
			for i := 0; i < distance; i++ {
				currentPos--
				if currentPos == 0 {
					numZeroPasses++
				}
				if currentPos < 0 {
					currentPos = 99
				}
			}
		default:
			log.Fatalf("invalid direction: %s", direction)
		}

	}
	return numZeroPasses

}

func pt2(input []string) int {
	currentPos := 50
	numZeroPasses := 0

	for _, line := range input {
		re := regexp.MustCompile(`^([LR])(\d+)$`)
		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			log.Fatalf("invalid line format: %s", line)
		}
		direction := matches[1]
		distance, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalf("failed to parse distance: %s", err)
		}
		// Now 'direction' holds "L" or "R", and 'distance' holds the integer value.

		// Helper for floor division to handle negative numbers correctly
		floorDiv := func(a, b int) int {
			if a >= 0 {
				return a / b
			}
			return (a - b + 1) / b
		}

		switch direction {
		case "R":
			// Increasing (Clockwise?)
			newPos := currentPos + distance
			// Count multiples of 100 in (currentPos, newPos]
			numZeroPasses += newPos/100 - currentPos/100
			currentPos = newPos % 100
		case "L":
			// Decreasing (Counter-Clockwise?)
			newPos := currentPos - distance
			// Count multiples of 100 in [newPos, currentPos)
			// Formula: floor((B-1)/100) - floor((A-1)/100) where B=currentPos, A=newPos
			numZeroPasses += floorDiv(currentPos-1, 100) - floorDiv(newPos-1, 100)
			currentPos = ((newPos % 100) + 100) % 100
		default:
			log.Fatalf("invalid direction: %s", direction)
		}

	}
	return numZeroPasses
}

func pt1(input []string) int {

	currentPos := 50
	numZeros := 0

	for _, line := range input {
		re := regexp.MustCompile(`^([LR])(\d+)$`)
		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			log.Fatalf("invalid line format: %s", line)
		}
		direction := matches[1]
		distance, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalf("failed to parse distance: %s", err)
		}
		// Now 'direction' holds "L" or "R", and 'distance' holds the integer value.
		switch direction {
		case "L":
			// Moving "left" (decreasing value, not turning your hand left but the lock itself) on a 0-99 lock.
			// Calculate new position after moving 'distance' to the left (decreasing).
			// The inner `(currentPos - distance) % 100` handles the wrap-around for positive results.
			// Adding `+ 100` ensures the result of the inner modulo is non-negative before the final modulo.
			// The outer `% 100` then correctly wraps the result into the 0-99 range, even if the intermediate result was negative (e.g., 50 - 60 = -10, then -10 % 100 = -10, then -10 + 100 = 90, then 90 % 100 = 90).
			currentPos = ((currentPos-distance)%100 + 100) % 100
		case "R":
			// Moving "right" (increasing value, not turning your hand right but the lock itself) on a 0-99 lock.
			// Example: 90 + 20 = 110. 110 % 100 = 10.
			// 99 + 300 = 399. 399 % 100 = 99.
			// 50 + 10976 = 11026. 11026 % 100 = 26.
			currentPos = (currentPos + distance) % 100
		default:
			log.Fatalf("invalid direction: %s", direction)
		}

		if currentPos == 0 {
			numZeros++
		}

	}
	return numZeros
}
