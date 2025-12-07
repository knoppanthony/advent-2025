package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		// Split by whitespace to handle multi-digit numbers and varying spaces
		fields := strings.Fields(line)
		if len(fields) > 0 {
			lines = append(lines, fields)
		}
	}

	return lines
}

func readInputRaw(filename string) [][]string {
	// Read input preserving all characters including spaces
	// This is needed for pt2 where column positions matter

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines [][]string

	for scanner.Scan() {
		line := scanner.Text()
		// Split into individual characters
		chars := strings.Split(line, "")
		lines = append(lines, chars)
	}

	return lines
}

func main() {
	input := readInput("input.txt")

	fmt.Println("pt1:", pt1(input))
	fmt.Println("pt2:", pt2(input))
}

func pt1(input [][]string) int {
	// Each column in the 2D array represents a math problem
	// The last row contains operators, rows above contain numbers
	// Iterate through each column from bottom to top

	if len(input) == 0 {
		return 0
	}

	result := 0

	// Find the maximum number of columns across all rows
	numCols := 0
	for _, row := range input {
		if len(row) > numCols {
			numCols = len(row)
		}
	}

	// Iterate through each column
	for col := 0; col < numCols; col++ {
		var columnSum int
		var operator string
		firstNum := true

		// Iterate through each row from bottom to top for this column
		for row := len(input) - 1; row >= 0; row-- {
			// Skip if this column doesn't exist in this row
			if col >= len(input[row]) {
				continue
			}

			value := input[row][col]

			// The last row contains the operator
			if row == len(input)-1 {
				operator = value
				continue
			}

			// Try to convert to number
			num, err := strconv.Atoi(value)
			if err != nil {
				// Skip non-numeric values
				continue
			}

			// Initialize with the first number
			if firstNum {
				columnSum = num
				firstNum = false
				continue
			}

			// Apply the operator
			switch operator {
			case "*":
				columnSum *= num
			case "+":
				columnSum += num
			}
		}

		// Add this column's result to the total
		result += columnSum
	}

	return result
}

func pt2(inputIgnored [][]string) int {
	// Read raw input to preserve spacing
	input := readInputRaw("input.txt")

	if len(input) == 0 {
		return 0
	}

	// Find max columns
	numCols := 0
	for _, row := range input {
		if len(row) > numCols {
			numCols = len(row)
		}
	}

	result := 0

	// Process from right to left, finding problems separated by space columns
	col := numCols - 1

	for col >= 0 {
		// Skip space columns
		isSpaceCol := true
		for row := 0; row < len(input)-1; row++ {
			if col < len(input[row]) && input[row][col] != " " && input[row][col] != "" {
				isSpaceCol = false
				break
			}
		}

		if isSpaceCol {
			col--
			continue
		}

		// Found end of a problem - find where it starts (to the left)
		problemEnd := col
		problemStart := col

		for problemStart >= 0 {
			isSpaceCol := true
			for row := 0; row < len(input)-1; row++ {
				if problemStart < len(input[row]) && input[row][problemStart] != " " && input[row][problemStart] != "" {
					isSpaceCol = false
					break
				}
			}

			if isSpaceCol {
				problemStart++
				break
			}
			problemStart--
		}

		if problemStart < 0 {
			problemStart = 0
		}

		// Get the operator for this problem
		var operator string
		for c := problemStart; c <= problemEnd; c++ {
			if c < len(input[len(input)-1]) {
				op := input[len(input)-1][c]
				if op == "+" || op == "*" {
					operator = op
					break
				}
			}
		}

		if operator == "" {
			col = problemStart - 1
			continue
		}

		// For each column in this problem, read top-to-bottom to form a number
		// Process columns right-to-left
		var numbers []int
		for c := problemEnd; c >= problemStart; c-- {
			digitStr := ""
			for row := 0; row < len(input)-1; row++ {
				if c < len(input[row]) && input[row][c] != " " && input[row][c] != "" {
					digitStr += input[row][c]
				}
			}

			if digitStr != "" {
				num, err := strconv.Atoi(digitStr)
				if err == nil {
					numbers = append(numbers, num)
				}
			}
		}

		// Apply the operator to all numbers
		if len(numbers) > 0 {
			problemResult := numbers[0]
			for i := 1; i < len(numbers); i++ {
				switch operator {
				case "*":
					problemResult *= numbers[i]
				case "+":
					problemResult += numbers[i]
				}
			}
			result += problemResult
		}

		// Move to next problem
		col = problemStart - 1
	}

	return result
}
