package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"unicode"
)

func main() {
	input := readFile("day5/input.txt")

	fmt.Printf("Day 5 Part 1: %d units remain after fully reacting the polymer\n", partOne(input))
	fmt.Printf("Day 5 Part 2: %d is the length of the shortest polymer you can produce by removing all units of exactly one type and fully reacting the result", partTwo(input))
}

func partOne(input string) int {
	return len(reduce(input))
}

func reduce(input string) string {
	var prev rune
	// Loop over all the characters
	for i, cur := range input {
		if i >= 1 {
			// Check if the chars are not the same, but differ in capitalization
			if cur != prev && unicode.ToLower(cur) == unicode.ToLower(prev) {
				// Recursion! Do the same, but then without the current and previous char in the string
				return reduce(input[:i-1] + input[i+1:])
			}
		}
		prev = cur
	}

	return input
}

func partTwo(input string) int {

	leastCount := 100000
	for exclude := 'a'; exclude <= 'z'; exclude++ {
		count := len(reduce2(input, exclude))
		if count < leastCount {
			leastCount = count
		}
	}

	return leastCount
}

func reduce2(input string, exclude rune) string {
	var prev rune
	// Loop over all the characters
	for i, cur := range input {
		// Skip the excluded char
		lowerCur := unicode.ToLower(cur)
		if lowerCur == exclude {
			return reduce2(input[:i]+input[i+1:], exclude)
		}

		if i >= 1 {
			// Check if the chars are not the same, but differ in capitalization
			if cur != prev && lowerCur == unicode.ToLower(prev) {
				// Recursion! Do the same, but then without the current and previous char in the string
				return reduce2(input[:i-1]+input[i+1:], exclude)
			}
		}
		prev = cur
	}

	return input
}

/**
Helper method to read file in a string
*/
func readFile(fileName string) string {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return string(contents)
}
