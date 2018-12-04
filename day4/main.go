package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	lines := parseLines("day4/input.txt")

	fmt.Printf("Day 4 Part 1: %d is the ID of the guard multiplied by the minute he slept most\n", partOne(lines))
	fmt.Printf("Day 4 Part 2: %d is the ID of the guard multiplied by the minute he slept most\n", partTwo(lines))
}

var (
	re  = regexp.MustCompile(`(?m)\[(\d{4}-\d{2}-\d{2}\s\d{2}:\d{2})]\s(.*)`)
	re2 = regexp.MustCompile(`(?m)Guard #(\d{1,4}) begins shift`)
)

func partOne(input []string) int {
	// First we parse the timestamp and store the string with it
	parsedLines := make(map[time.Time]string)
	for _, line := range input {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			text := match[2]

			parsedTime, _ := time.Parse("2006-01-02 15:04", match[1])
			parsedLines[parsedTime] = text
		}
	}

	// In Go maps are unsorted... so we have to sort it by using an intermediate array
	// To store the keys in slice in sorted order
	var keys []time.Time
	for k := range parsedLines {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i].Unix() < keys[j].Unix() })

	// Looping over the sorted keys
	activeGuard := 0
	startSleepMinute := 0
	minutesAsleepPerGuard := make(map[int]map[int]int)
	for _, k := range keys {
		line := parsedLines[k]
		if line == "wakes up" {
			stopSleepMinute := k.Minute()
			for i := startSleepMinute; i < stopSleepMinute; i++ {
				minutesAsleepPerGuard[activeGuard][i]++
			}
		}

		if line == "falls asleep" {
			startSleepMinute = k.Minute()
		}

		if match := re2.FindStringSubmatch(line); len(match) > 0 {
			activeGuard, _ = strconv.Atoi(match[1])

			// Create submap for this guard if not already present
			if _, exists := minutesAsleepPerGuard[activeGuard]; !exists {
				minutesAsleepPerGuard[activeGuard] = make(map[int]int)
			}
		}
	}

	// Find which guard has the most minutes
	guardMostMinutesId := 0
	guardMostMinutesCount := 0
	guardMostMinutesMinute := 0
	for guardId, minutesData := range minutesAsleepPerGuard {
		totalMinutes := 0
		mostMinutesMinute := 0
		mostMinutesCount := 0
		for minute, count := range minutesData {
			totalMinutes += count
			if count > mostMinutesCount {
				mostMinutesCount = count
				mostMinutesMinute = minute
			}
		}

		// If this guard sleeps longer then the previous max sleeper
		if totalMinutes > guardMostMinutesCount {
			// Store the state
			guardMostMinutesId = guardId
			guardMostMinutesCount = totalMinutes
			guardMostMinutesMinute = mostMinutesMinute
		}
	}

	return guardMostMinutesId * guardMostMinutesMinute
}

func partTwo(input []string) int {
	// First we parse the timestamp and store the string with it
	parsedLines := make(map[time.Time]string)
	for _, line := range input {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			text := match[2]

			parsedTime, _ := time.Parse("2006-01-02 15:04", match[1])
			parsedLines[parsedTime] = text
		}
	}

	// In Go maps are unsorted... so we have to sort it by using an intermediate array
	// To store the keys in slice in sorted order
	var keys []time.Time
	for k := range parsedLines {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i].Unix() < keys[j].Unix() })

	// Looping over the sorted keys
	activeGuard := 0
	startSleepMinute := 0
	minutesAsleepPerGuard := make(map[int]map[int]int)
	for _, k := range keys {
		line := parsedLines[k]
		if line == "wakes up" {
			stopSleepMinute := k.Minute()
			for i := startSleepMinute; i < stopSleepMinute; i++ {
				minutesAsleepPerGuard[activeGuard][i]++
			}
		}

		if line == "falls asleep" {
			startSleepMinute = k.Minute()
		}

		if match := re2.FindStringSubmatch(line); len(match) > 0 {
			activeGuard, _ = strconv.Atoi(match[1])

			// Create submap for this guard if not already present
			if _, exists := minutesAsleepPerGuard[activeGuard]; !exists {
				minutesAsleepPerGuard[activeGuard] = make(map[int]int)
			}
		}
	}

	// Find which guard has the highest count on any given minute
	guardMostMinutesId := 0
	guardMostMinutesCount := 0
	guardMostMinutesMinute := 0
	for guardId, minutesData := range minutesAsleepPerGuard {
		mostMinutesMinute := 0
		mostMinutesCount := 0
		for minute, count := range minutesData {
			if count > mostMinutesCount {
				mostMinutesCount = count
				mostMinutesMinute = minute
			}
		}

		// If this guard sleeps longer then the previous max sleeper
		if mostMinutesCount > guardMostMinutesCount {
			// Store the state
			guardMostMinutesId = guardId
			guardMostMinutesCount = mostMinutesCount
			guardMostMinutesMinute = mostMinutesMinute
		}
	}

	return guardMostMinutesId * guardMostMinutesMinute
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
