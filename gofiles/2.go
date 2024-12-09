package gofiles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Two() {
	numSafe, err := check_levels()
	if err != nil {
		fmt.Println("Error checking levels:", err)
		return
	}
	fmt.Println("Safe reports:", numSafe)
}

func check_levels() (int, error) {
	// Open the file
	file, err := os.Open("2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, err
	}
	defer file.Close()

	numSafe := 0

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line by white space
		line := scanner.Text()
		fields := strings.Fields(line)
		// fmt.Println(line)

		// Skip the trivial case of one value
		if len(fields) < 2 {
			numSafe += 1
			continue
		}

		// cast report to int slice
		report, err := get_report(fields)

		if err != nil {
			fmt.Println("Error getting report:", err)
			return 0, err
		}

		if test_report(report) {
			numSafe += 1
		}
	}

	// Check if any error occurred while reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0, err
	}

	return numSafe, nil
}

func get_report(strReport []string) ([]int, error) {
	var report []int
	for _, strVal := range strReport {
		val, err := strconv.Atoi(strVal)
		if err != nil {
			fmt.Println("Error casting string to int:", err)
			return []int{}, err
		}
		report = append(report, val)
	}
	return report, nil
}

func test_conditions(reactorReport []int) int {
	// Return first bad level, past end index means success
	increasing := reactorReport[1]-reactorReport[0] >= 0
	lastVal := reactorReport[0]
	for idx, val := range reactorReport[1:] {
		if increasing && val-lastVal > 0 && val-lastVal <= 3 {
			lastVal = val
			continue
		} else if !increasing && lastVal-val > 0 && lastVal-val <= 3 {
			lastVal = val
			continue
		} else {
			return idx + 1
		}
	}
	return len(reactorReport)
}

func test_report(reactorReport []int) bool {
	// Brute force testing by dropping either of the first two
	// levels, to determine monotonic increase/decrease condition
	lastIdx := len(reactorReport) - 1

	// 0 no level is bad, or last level is bad
	firstBadIdx := test_conditions(reactorReport)
	if firstBadIdx >= lastIdx {
		// fmt.Println("SAFE: No level was bad")
		return true
	}

	// 1. first level is bad
	if test_conditions(reactorReport[1:]) == lastIdx {
		// fmt.Println("SAFE: First level was bad")
		return true
	}

	// 2. second level is bad
	secondRemoved := append([]int(nil), reactorReport[0])
	secondRemoved = append(secondRemoved, reactorReport[2:]...)
	if test_conditions(secondRemoved) == lastIdx {
		// fmt.Println("SAFE: Second level was bad")
		return true
	}

	// 3. any other level is bad
	otherRemoved := append([]int(nil), reactorReport[:firstBadIdx]...)
	otherRemoved = append(otherRemoved, reactorReport[firstBadIdx+1:]...)
	if test_conditions(otherRemoved) == lastIdx {
		// fmt.Printf("SAFE: %d level was bad\n", firstBadIdx)
		return true
	}

	// fmt.Printf("UNSAFE: first bad idx %d\n", firstBadIdx)

	return false
}
