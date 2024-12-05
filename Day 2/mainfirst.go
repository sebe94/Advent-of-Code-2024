package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafe(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		// Difference between two consecutive numbers
		diff := numbers[i+1] - numbers[i]

		// Condition: difference greater than 3 or less than -3
		if diff > 3 || diff < -3 {
			return false
		}

		// Condition: no change (remains constant)
		if diff == 0 {
			return false
		}

		// Condition: switch from increasing to decreasing or vice versa
		if i > 0 {
			prevDiff := numbers[i] - numbers[i-1]
			if (prevDiff > 0 && diff < 0) || (prevDiff < 0 && diff > 0) {
				return false
			}
		}
	}
	return true
}

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}
	defer file.Close()

	// Read the content of the file
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading the content of the file: %v", err)
	}

	// Split the content into lines
	lines := strings.Split(string(content), "\n")

	// Variables for counting safe and unsafe reports
	var safeReportsCount int
	var unsafeReportsCount int

	// Iterate through each line and convert to numbers
	for _, line := range lines {
		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Split the line into string values
		values := strings.Fields(line)

		// Convert values to integers
		var numbers []int
		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				log.Fatalf("Invalid number in the input: %v", err)
			}
			numbers = append(numbers, num)
		}

		// Check if the current line (report) is safe or unsafe
		if isSafe(numbers) {
			safeReportsCount++ // Increment the safe reports counter
		} else {
			unsafeReportsCount++ // Increment the unsafe reports counter
		}
	}

	// Output the count of safe and unsafe reports
	fmt.Printf("There are %d safe reports.\n", safeReportsCount)
	fmt.Printf("There are %d unsafe reports.\n", unsafeReportsCount)
}
