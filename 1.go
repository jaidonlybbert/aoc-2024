package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lList, rList, err := parse_input()

	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	compute_distances(lList, rList)
	compute_similarity(lList, rList)
}

func compute_similarity(lList []int, rList []int) {
	similarity := 0
	occurrences := make(map[int]int)

	for _, val := range lList {
		occurrences[val] = 0
	}

	for _, val := range rList {
		occurrences[val] = 0
	}

	for _, val := range rList {
		occurrences[val] += 1
	}

	for _, val := range lList {
		similarity += val * occurrences[val]
	}

	fmt.Println("Similarity score:", similarity)
}

func compute_distances(lList []int, rList []int) {
	sort.Ints(lList)
	sort.Ints(rList)

	sum := 0
	for i := 0; i < len(lList); i++ {
		sum += int(math.Abs(float64(lList[i] - rList[i])))
	}

	fmt.Println("Distance sum:", sum)
}

func parse_input() ([]int, []int, error) {

	// Open the file
	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []int{}, []int{}, err
	}
	defer file.Close()

	// Declare a slice to hold the numbers
	var lList []int
	var rList []int

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line by tabs
		line := scanner.Text()
		fields := strings.Fields(line)

		rval, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Println("Error parsing digit", err)
			return []int{}, []int{}, err
		}

		lval, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("Error parsing digit", err)
			return []int{}, []int{}, err
		}

		lList = append(lList, lval)
		rList = append(rList, rval)
	}

	// Check if any error occurred while reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return []int{}, []int{}, err
	}

	return lList, rList, nil
}

