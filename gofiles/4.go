package gofiles

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Four() {
	lines, err := parse_input_4()
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	// Part 1
	count := count_all_matches(lines)
	fmt.Println("Total count of matches:", count)

	// Part 2
	count = count_all_x_mas(lines)
	fmt.Println("Total x-mas count:", count)
}

func parse_input_4() ([]string, error) {
	// Open the file
	file, err := os.Open("4.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []string(nil), err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	lines := []string(nil)
	for scanner.Scan() {
		// Split the line by white space
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Check if any error occurred while reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return []string(nil), err
	}
	return lines, nil
}

func count_all_x_mas(lines []string) int {
	count := 0

	for row_idx, line := range lines {
		if row_idx == 0 || row_idx == len(lines)-1 {
			continue
		}
		for col_idx, char := range line {
			if col_idx == 0 || col_idx == len(lines[0])-1 {
				continue
			}
			if char == 'A' {
				if check_x_mas(lines, row_idx, col_idx) {
					count += 1
				}
			}
		}
	}

	return count
}

func check_x_mas(lines []string, row_idx int, col_idx int) bool {
	if (lines[row_idx-1][col_idx-1] == 'M' && lines[row_idx+1][col_idx+1] == 'S' ||
		lines[row_idx-1][col_idx-1] == 'S' && lines[row_idx+1][col_idx+1] == 'M') &&
		(lines[row_idx-1][col_idx+1] == 'M' && lines[row_idx+1][col_idx-1] == 'S' ||
			lines[row_idx-1][col_idx+1] == 'S' && lines[row_idx+1][col_idx-1] == 'M') {
		return true
	}
	return false
}

func count_all_matches(lines []string) int {
	count := 0
	count += count_matches(lines)
	count += count_matches(build_vertical(lines))
	count += count_matches(build_diagonal_descending(lines))
	count += count_matches(build_diagonal_ascending(lines))
	return count
}

func count_matches(lines []string) int {
	count := 0
	fwrd_re := regexp.MustCompile(`XMAS`)
	back_re := regexp.MustCompile(`SAMX`)
	for _, val := range lines {
		count += len(fwrd_re.FindAllString(val, -1))
		count += len(back_re.FindAllString(val, -1))
	}

	return count
}

func build_vertical(lines []string) []string {
	vertical_strings := []string(nil)
	for lidx, val := range lines {
		for idx, char := range val {
			if lidx == 0 {
				vertical_strings = append(vertical_strings, string(char))
			}
			vertical_strings[idx] += string(char)
		}
	}

	return vertical_strings
}

func build_diagonal_descending(lines []string) []string {
	num_rows := len(lines)
	num_cols := len(lines[0])
	num_diagonals := num_cols + num_rows - 1
	diagonal_strings := make([]string, num_diagonals)
	// Indexing of diagonals is chosen so that the bottom
	// left character is diagonal 0 and the top right is the
	// len(diagonal_strings)-1 diagonal
	initial_offset := num_rows - 1
	for rowidx, val := range lines {
		for colidx, char := range val {
			diagonal_strings[initial_offset-rowidx+colidx] += string(char)
		}
	}

	return diagonal_strings
}

func build_diagonal_ascending(lines []string) []string {
	num_rows := len(lines)
	num_cols := len(lines[0])
	num_diagonals := num_cols + num_rows - 1
	diagonal_strings := make([]string, num_diagonals)
	// Indexing of diagonals is chosen so that the top
	// left character is diagonal 0 and the bottom right is the
	// len(diagonal_strings)-1 diagonal
	for rowidx, val := range lines {
		for colidx, char := range val {
			diagonal_strings[colidx+rowidx] += string(char)
		}
	}

	return diagonal_strings
}
