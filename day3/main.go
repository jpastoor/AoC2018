package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines := parseLines("day3/input.txt")

	fmt.Printf("Day 3 Part 1: %d square inches of fabric are within two or more claims\n", partOne(lines))
	fmt.Printf("Day 3 Part 2: %d is the ID of the only claim that doesn't overlap\n", partTwo(lines))
}

func partOne(input []string) int {
	var re = regexp.MustCompile(`(?m)#(\d*) @ (\d{1,4}),(\d{1,4}): (\d{1,4})x(\d{1,4})`)

	claimedFabric := make(map[string]int)
	for _, line := range input {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			x, _ := strconv.Atoi(match[2])
			y, _ := strconv.Atoi(match[3])
			w, _ := strconv.Atoi(match[4])
			h, _ := strconv.Atoi(match[5])

			for i := x; i < x+w; i++ {
				for j := y; j < y+h; j++ {
					coords := fmt.Sprintf("%d,%d", i, j)
					claimedFabric[coords]++
				}
			}
		}
	}

	doubleClaimed := 0
	for _, claims := range claimedFabric {
		if claims >= 2 {
			doubleClaimed++
		}
	}

	return doubleClaimed
}

func partTwo(input []string) int {
	var re = regexp.MustCompile(`(?m)#(\d*) @ (\d{1,4}),(\d{1,4}): (\d{1,4})x(\d{1,4})`)

	idsOverlapped := make(map[int]bool)
	claimedFabric := make(map[string]int)
	for _, line := range input {
		for _, match := range re.FindAllStringSubmatch(line, -1) {

			id, _ := strconv.Atoi(match[1])
			x, _ := strconv.Atoi(match[2])
			y, _ := strconv.Atoi(match[3])
			w, _ := strconv.Atoi(match[4])
			h, _ := strconv.Atoi(match[5])

			idsOverlapped[id] = false
			for i := x; i < x+w; i++ {
				for j := y; j < y+h; j++ {
					coords := fmt.Sprintf("%d,%d", i, j)

					// Similar as before, but now we first check if there is already a claim for these coords
					// if thats the case, both the current claim and the claim of that position is invalidated
					if overlapsId, exists := claimedFabric[coords]; exists {
						idsOverlapped[id] = true
						idsOverlapped[overlapsId] = true
					}

					claimedFabric[coords] = id
				}
			}
		}
	}

	for id, isOverlapped := range idsOverlapped {
		if !isOverlapped {
			return id
		}
	}
	return 0
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
