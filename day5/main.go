package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type FoodRange struct {
	start int
	end   int
}

func readFoodRanges(filename string) ([]FoodRange, []int) {

	// 1. Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err) // Use log.Fatalf for concise error handling
	}
	// Ensure the file is closed after the function exits
	defer file.Close()

	// 2. Create a scanner
	scanner := bufio.NewScanner(file)

	var lines []FoodRange
	var freshFoodIds []int
	//scan the first half of the file into a 1d array of FoodRange structs
	//check for a blank line and then switch to scanning the second half of the file into an array of ints that represent fresh food ids
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		//split by '-' to get start and end of FoodRange
		lineSplit := strings.Split(line, "-")
		start, _ := strconv.Atoi(lineSplit[0])
		end, _ := strconv.Atoi(lineSplit[1])
		lines = append(lines, FoodRange{start: start, end: end})
	}
	//scan the second half of the file into an array of ints that represent fresh food ids
	for scanner.Scan() {
		line := scanner.Text()
		id, _ := strconv.Atoi(line)
		freshFoodIds = append(freshFoodIds, id)
	}

	return lines, freshFoodIds

}

func main() {

	foodRanges, freshFoodIds := readFoodRanges("input.txt")
	fmt.Println(foodRanges)
	fmt.Println(freshFoodIds)
	fmt.Println("pt1:", pt1(foodRanges, freshFoodIds))
	fmt.Println("pt2:", pt2(foodRanges))
}

func mergeFoodRanges(foodRanges []FoodRange) []FoodRange {
	//sort the ranges by start with slices.sortFunc
	slices.SortFunc(foodRanges, func(a, b FoodRange) int {
		return a.start - b.start
	})

	var mergedRanges []FoodRange
	for _, foodRange := range foodRanges {
		if len(mergedRanges) == 0 || mergedRanges[len(mergedRanges)-1].end < foodRange.start {
			mergedRanges = append(mergedRanges, foodRange)
		} else {
			mergedRanges[len(mergedRanges)-1].end = max(mergedRanges[len(mergedRanges)-1].end, foodRange.end)
		}
	}
	return mergedRanges
}

func pt1(foodRanges []FoodRange, freshFoodIds []int) int {

	foodRanges = mergeFoodRanges(foodRanges)

	totalFreshFoodIds := 0
	//loop through freshFoodIds and check if they are in any of the foodRanges
	for _, freshFoodId := range freshFoodIds {
		for _, foodRange := range foodRanges {
			if freshFoodId >= foodRange.start && freshFoodId <= foodRange.end {
				totalFreshFoodIds++
			}
		}
	}
	return totalFreshFoodIds
}

func pt2(foodRanges []FoodRange) int {
	foodRanges = mergeFoodRanges(foodRanges)

	//loop through all the merged ranges and just count the how many numbers are in each range, inclusive
	totalFoodRanges := 0
	for _, foodRange := range foodRanges {
		totalFoodRanges += foodRange.end - foodRange.start + 1
	}
	return totalFoodRanges
}
