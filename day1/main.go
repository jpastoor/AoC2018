package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := parseLines("day1/input.txt")

	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}

func partOne(input []string) int {
	frequency := 0
	for _, line := range input {
		parsedInt, _ := strconv.Atoi(line)
		frequency += parsedInt
	}

	return frequency
}

func partTwo(input []string) int {
	frequency := 0
	reachedFrequencies := make(map[int]int)
	for {
		for _, line := range input {
			parsedInt, _ := strconv.Atoi(line)
			frequency += parsedInt

			if _, exists := reachedFrequencies[frequency]; !exists {
				reachedFrequencies[frequency] = 1
			} else {
				// reached for the second time
				return frequency
			}
		}
	}
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
