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

// IsRepeatingNumber returns true if s is made of exactly
// two copies of the same digit sequence.
func IsRepeatingNumber(s string) bool {
	n := len(s)

	// Must be even length to split evenly into 2 parts
	if n%2 != 0 {
		return false
	}

	half := n / 2
	first := s[:half]
	second := s[half:]

	return first == second
}

func IsMultiRepeatingNumber(s string) bool {
	n := len(s)

	// Need at least two repeats of something => length >= 2*1 = 2
	if n < 2 {
		return false
	}

	// Try all possible pattern sizes from 1 up to half the string
	for size := 1; size <= n/2; size++ {
		if n%size == 0 { // size must divide total length
			repeatCount := n / size
			if repeatCount >= 2 { // must repeat at least twice
				pattern := s[:size]
				if strings.Repeat(pattern, repeatCount) == s {
					return true
				}
			}
		}
	}

	return false
}

type numRange struct {
	lower int
	upper int
}

func readInput(filename string) []numRange {

	// 1. Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err) // Use log.Fatalf for concise error handling
	}
	// Ensure the file is closed after the function exits
	defer file.Close()

	// 2. Create a scanner
	scanner := bufio.NewScanner(file)

	// multiple ranges defined in the files in the following format integer-integer,integer-integer with any number of ranges
	// where the first integer is the lower bound and the second integer is the upper bound
	// load each range into a struct and return a list of ranges
	var ranges []numRange
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`([0-9]+)-([0-9]+)`)
		allMatches := re.FindAllStringSubmatch(line, -1)
		if len(allMatches) == 0 {
			log.Fatalf("no ranges found in line: %s", line)
		}
		for _, matches := range allMatches {
			if len(matches) != 3 { // Full match, lower bound, upper bound
				log.Fatalf("invalid range format in line: %s", line)
			}
			lower, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatalf("failed to parse lower bound '%s' in line '%s': %s", matches[1], line, err)
			}
			upper, err := strconv.Atoi(matches[2])
			if err != nil {
				log.Fatalf("failed to parse upper bound '%s' in line '%s': %s", matches[2], line, err)
			}
			ranges = append(ranges, numRange{lower, upper})
		}
	}

	return ranges
}

func main() {

	input := readInput("input.txt")
	fmt.Println("pt1:", pt1(input))
	fmt.Println("pt2:", pt2(input))

}

func pt1(input []numRange) int {
	//look at the lower and upper bounds of each range and every number between them.
	//Convert each number to a string and determine if the number is made of a repeating pattern
	//For example, 1188511885 is made of the pattern 11885 repeated twice.
	invalidNumberSum := 0
	for _, numRange := range input {
		for i := numRange.lower; i <= numRange.upper; i++ {
			str := strconv.Itoa(i)
			if IsRepeatingNumber(str) {
				fmt.Println(i)
				invalidNumberSum += i
			}
		}
	}

	return invalidNumberSum
}

func pt2(input []numRange) int {
	invalidNumberSum := 0
	for _, numRange := range input {
		for i := numRange.lower; i <= numRange.upper; i++ {
			str := strconv.Itoa(i)
			if IsMultiRepeatingNumber(str) {
				fmt.Println(i)
				invalidNumberSum += i
			}
		}
	}

	return invalidNumberSum
}
