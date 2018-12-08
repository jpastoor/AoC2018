package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := readFile("day8/input.txt")
	demo := readFile("day8/demo.txt")

	fmt.Printf("Day 8 Part Demo: %d is the sum of all metadata entries\n", partOne(demo))
	fmt.Printf("Day 8 Part 1: %d is the sum of all metadata entries\n", partOne(input))
	fmt.Printf("Day 8 Part Demo 2: %d is the value of the root node\n", partTwo(demo))
	fmt.Printf("Day 8 Part 2: %d  is the value of the root node\n", partTwo(input))
}

func partOne(input string) int {
	// Parse input
	codeStrings := strings.Fields(input)
	codes := make([]int, len(codeStrings))
	for i, code := range codeStrings {
		codes[i], _ = strconv.Atoi(code)
	}

	_, sumMetadata := parseNode(codes)

	return sumMetadata
}

func parseNode(input []int) (int, int) {
	metadataSum := 0
	quantChildNodes := input[0]
	quantMetaEntries := input[1]
	marker := 2
	for iChild := 1; iChild <= quantChildNodes; iChild++ {
		childMarker, childMetadataSum := parseNode(input[marker:])
		metadataSum += childMetadataSum
		marker += childMarker
	}

	for iMeta := 0; iMeta < quantMetaEntries; iMeta++ {
		metadataSum += input[marker+iMeta]
	}

	return marker+quantMetaEntries, metadataSum
}


func partTwo(input string) int {
	// Parse input
	codeStrings := strings.Fields(input)
	codes := make([]int, len(codeStrings))
	for i, code := range codeStrings {
		codes[i], _ = strconv.Atoi(code)
	}

	_, sumMetadata := parseNode2(codes)

	return sumMetadata
}

func parseNode2(input []int) (int, int) {
	nodeScore := 0
	quantChildNodes := input[0]
	quantMetaEntries := input[1]
	marker := 2

	// Keep the scores of all childs in case they are references later
	childValues := make(map[int]int)
	for iChild := 1; iChild <= quantChildNodes; iChild++ {
		childMarker, childScore := parseNode2(input[marker:])
		childValues[iChild] = childScore
		marker += childMarker
	}

	for iMeta := 0; iMeta < quantMetaEntries; iMeta++ {
		// If there are no childnodes, use this value for the score
		if quantChildNodes == 0 {
			nodeScore += input[marker+iMeta]
		} else if childValue, exists := childValues[input[marker+iMeta]]; exists {
			// If the childNode being references exists, use that value
			nodeScore += childValue
		}
	}

	return marker+quantMetaEntries, nodeScore
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
