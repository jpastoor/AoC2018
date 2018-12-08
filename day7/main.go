package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

func main() {
	input := parseLines("day7/input.txt")
	demo := parseLines("day7/demo.txt")

	fmt.Printf("Day 7 Part Demo: %s is the order the steps should be completed\n", partOne(demo))
	fmt.Printf("Day 7 Part 1: %s is the order the steps should be completed\n", partOne(input))
	fmt.Printf("Day 7 Part 2: %d will it take to complete all of the steps", partTwo(input))
}

var re = regexp.MustCompile(`(?m)Step (\S) must be finished before step (\S) can begin.`)

func partOne(input []string) string {
	// Parse input
	var parsedInput []InputLine
	for _, line := range input {
		matches := re.FindStringSubmatch(line)
		parsedInput = append(parsedInput, InputLine{matches[2], matches[1]})
	}

	// Make a map that stores for every step the list of prerequisite steps
	stepPrereqs := make(map[string]map[string]bool)
	for _, parsedInputLine := range parsedInput {
		addPrereq(stepPrereqs, parsedInputLine.Prereq, parsedInputLine.Step)
	}

	route := ""
	for {
		// We are done when all steps are done
		if len(stepPrereqs) == 0 {
			break
		}

		// Find all steps without prereqs
		stepsWithoutPrereqs := []string{}
		for step, stepPrereq := range stepPrereqs {
			if len(stepPrereq) == 0 {
				stepsWithoutPrereqs = append(stepsWithoutPrereqs, step)
			}
		}

		// Sort the steps
		sort.Strings(stepsWithoutPrereqs)
		nextStep := stepsWithoutPrereqs[0]
		delete(stepPrereqs, nextStep)
		removeStepFromPrereqs(stepPrereqs, nextStep)
		route += nextStep
	}

	return route
}

func partTwo(input []string) int {
	// Parse input
	var parsedInput []InputLine
	for _, line := range input {
		matches := re.FindStringSubmatch(line)
		parsedInput = append(parsedInput, InputLine{matches[2], matches[1]})
	}

	// Make a map that stores for every step the list of prerequisite steps
	stepPrereqs := make(map[string]map[string]bool)
	for _, parsedInputLine := range parsedInput {
		addPrereq(stepPrereqs, parsedInputLine.Prereq, parsedInputLine.Step)
	}

	workers := []*Worker{
		&Worker{Id: 1},
		&Worker{Id: 2},
		&Worker{Id: 3},
		&Worker{Id: 4},
		&Worker{Id: 5},
	}

	route := ""
	t := 0
	for {
		// Check if any steps can be marked as done
		stepsDone := []string{}
		for _, worker := range workers {
			if worker.DoneAt == t {
				// TODO I might not have good support for 2 workers done in the same turn
				stepsDone = append(stepsDone, worker.Step)
				worker.Step = ""
				worker.DoneAt = 0
			}
		}

		sort.Strings(stepsDone)
		for _, step := range stepsDone {
			delete(stepPrereqs, step)
			removeStepFromPrereqs(stepPrereqs, step)
			route += step
		}

		// We are done when all steps are done
		if len(stepPrereqs) == 0 {
			break
		}

		// Find all steps without prereqs
		var stepsWithoutPrereqs []string
		for step, stepPrereq := range stepPrereqs {
			if len(stepPrereq) == 0 {
				stepsWithoutPrereqs = append(stepsWithoutPrereqs, step)
			}
		}

		// Sort the steps
		sort.Strings(stepsWithoutPrereqs)

		// Loop over the steps
		for _, step := range stepsWithoutPrereqs {
			// Make sure no workers already on it
			isWIP := false
			for _, worker := range workers {
				if worker.Step == step {
					isWIP = true
				}
			}

			if isWIP {
				continue
			}

			// Find an available worker
			for _, worker := range workers {
				if worker.Step == "" {
					worker.Step = step
					jobTime := int(step[0]) - 64 + 60
					//fmt.Printf("%s takes %d seconds\n", step, jobTime)
					worker.DoneAt = t + jobTime
					break
				}
			}
		}

		fmt.Printf("%02d %01s %01s %01s %01s %01s %s\n", t, workers[0].Step, workers[1].Step, workers[2].Step, workers[3].Step, workers[4].Step, route)

		t++
	}

	return t
}

func addPrereq(data map[string]map[string]bool, prereq string, step string) {
	if _, exists := data[step]; !exists {
		data[step] = make(map[string]bool)
	}

	data[step][prereq] = true

	// Also make sure any prereqs are added as a possible step
	if _, exists := data[prereq]; !exists {
		data[prereq] = make(map[string]bool)
	}
}

func removeStepFromPrereqs(data map[string]map[string]bool, stepToRemove string) {
	for step, _ := range data {
		delete(data[step], stepToRemove)
	}
}

type InputLine struct {
	Step   string
	Prereq string
}

type Worker struct {
	Id     int
	Step   string
	DoneAt int
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
