package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := parseLines("day2/input.txt")

	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}

func partOne(input []string) int {

	occTwo := 0
	occThree := 0

	for _, line := range input {
		// Occurrences per letter
		occPerLetter := make(map[rune]int)
		for _, char := range line {
			occPerLetter[char]++
		}

		// Transform to amount of occurrences
		occ := make(map[int]int)
		for _, amount := range occPerLetter {
			occ[amount]++
		}

		if _, exists := occ[2]; exists {
			occTwo++
		}

		if _, exists := occ[3]; exists {
			occThree++
		}
	}

	return occTwo * occThree
}

func partTwo(input []string) string {
	// Loop over both a and b so we get all combination
	for _, a := range input {
		for _, b := range input {
			if a == b {
				continue
			}

			// Count the amount of characters that a and b diff from eachother
			diff := 0
			bRunes := []rune(b)
			var diffI int
			for i, char := range a {
				if bRunes[i] != char {
					diff++
					diffI = i
				}
			}

			if diff == 1 {
				// Remove the character that contained the diff
				return a[:diffI]+a[diffI+1:]
			}
		}
	}

	return ""
}

/**
Helper method to convert file to lines
 */
func parseLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var contents []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return contents
}
