package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput(filename string) []string {

	// 1. Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err) // Use log.Fatalf for concise error handling
	}
	// Ensure the file is closed after the function exits
	defer file.Close()

	// 2. Create a scanner
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines

}

func main() {

	input := readInput("input.txt")

	fmt.Println("pt1:", pt1(input))
	fmt.Println("pt2:", pt2(input))

}

func findLargestConcatenated(line string) int {
	//Find the largest digit that isn't the last in the string and assign to first digit
	//Then scan for the largest digit after the index of the first digit.
	var firstDigit int = 0
	var secondDigit int = 0
	var firstDigitIndex int = 0
	fmt.Println(line)

	//get first largest digit, one short of the len of the line
	for i := 0; i < len(line)-1; i++ {
		currentNum, _ := strconv.Atoi(string(line[i]))
		if currentNum > firstDigit {
			firstDigit = currentNum
			firstDigitIndex = i
		}
	}

	//get second largest digit, starting from the index one after the first digit index
	for i := firstDigitIndex + 1; i < len(line); i++ {
		currentNum, _ := strconv.Atoi(string(line[i]))
		if currentNum > secondDigit {
			secondDigit = currentNum
		}
	}
	//fmt.Println(firstDigit)
	//fmt.Println(secondDigit)
	//return the integer value of combining the first digit and second digit by just concatenating them
	return firstDigit*10 + secondDigit

}

func pt1(input []string) int {
	totalOutputVoltage := 0
	for _, line := range input {
		totalOutputVoltage += findLargestConcatenated(line)
	}
	return totalOutputVoltage
}

func findLargest12Digit(line string) string {
	n := len(line)
	target := 12

	// dp[i][k] = largest k-digit number we can form using digits from positions 0 to i-1
	// We store as strings to avoid integer overflow and for easy lexicographic comparison
	dp := make([][]string, n+1)
	for i := range dp {
		dp[i] = make([]string, target+1)
	}

	// Base case: empty string for 0 digits
	for i := 0; i <= n; i++ {
		dp[i][0] = ""
	}

	// Fill the DP table
	for i := 1; i <= n; i++ {
		currentDigit := string(line[i-1])

		for k := 1; k <= target; k++ {
			// Option 1: Don't include current digit
			dp[i][k] = dp[i-1][k]

			// Option 2: Include current digit (if we have enough previous digits)
			if k == 1 {
				// First digit: just compare with previous best single digit
				if dp[i][k] == "" || currentDigit > dp[i][k] {
					dp[i][k] = currentDigit
				}
			} else if dp[i-1][k-1] != "" {
				// We can form a k-digit number by appending current digit to (k-1)-digit number
				candidate := dp[i-1][k-1] + currentDigit

				// Compare lexicographically (works because all strings have same length)
				if dp[i][k] == "" || candidate > dp[i][k] {
					dp[i][k] = candidate
				}
			}
		}
	}

	return dp[n][target]
}

func pt2(input []string) int {
	totalOutputVoltage := 0
	for _, line := range input {
		result := findLargest12Digit(line)
		if result != "" {
			// Convert string to int (assuming it fits in int)
			num, _ := strconv.Atoi(result)
			totalOutputVoltage += num
		}
	}
	return totalOutputVoltage
}
