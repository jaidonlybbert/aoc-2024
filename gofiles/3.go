package gofiles

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Three() {
	// part 1
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	matchesStr, err := get_matching_strs(re)
	if err != nil {
		fmt.Println("Error parsing corrupted program", err)
		return
	}
	sum, err := process_matches_pt1(matchesStr)
	if err != nil {
		fmt.Println("Error computing sums:", err)
		return
	}
	fmt.Println("Corrupted Sum Pt1:", sum)

	// part 2
	re = regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	matchesStr, err = get_matching_strs(re)
	if err != nil {
		fmt.Println("Error parsing corrupted program", err)
		return
	}
	sum, err = process_matches_pt2(matchesStr)
	if err != nil {
		fmt.Println("Error computing sums:", err)
		return
	}
	fmt.Println("Corrupted Sum Pt2:", sum)
}

func process_matches_pt1(matchesStr []string) (int, error) {
	sum := 0
	re := regexp.MustCompile(`[0-9]{1,3}`)
	for _, val := range matchesStr {
		nums := re.FindAllString(val, 2)
		val1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("Error casting value:", err)
			return 0, err
		}
		val2, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println("Error casting value:", err)
			return 0, err
		}
		sum += val1 * val2
	}

	return sum, nil
}

func process_matches_pt2(matchesStr []string) (int, error) {
	sum := 0
	re := regexp.MustCompile(`[0-9]{1,3}`)
	redo := regexp.MustCompile(`do\(\)`)
	redont := regexp.MustCompile(`don't\(\)`)
	isActive := true
	for _, val := range matchesStr {
		if redo.MatchString(val) {
			isActive = true
		} else if redont.MatchString(val) {
			isActive = false
		} else if isActive {
			nums := re.FindAllString(val, 2)
			val1, err := strconv.Atoi(nums[0])
			if err != nil {
				fmt.Println("Error casting value:", err)
				return 0, err
			}
			val2, err := strconv.Atoi(nums[1])
			if err != nil {
				fmt.Println("Error casting value:", err)
				return 0, err
			}
			sum += val1 * val2
		}
	}

	return sum, nil
}

func get_matching_strs(re *regexp.Regexp) ([]string, error) {
	// Open the file
	file, err := os.Open("3.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []string(nil), err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	muls := []string(nil)
	for scanner.Scan() {
		// Split the line by white space
		line := scanner.Text()
		muls = append(muls, re.FindAllString(line, -1)...)
	}

	// Check if any error occurred while reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return []string(nil), err
	}
	return muls, nil
}
