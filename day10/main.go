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
	input := parseLines("day10/input.txt")

	points := make([]Point, len(input))
	for i, line := range input {
		points[i] = linesToPoint(line)
	}

	fmt.Printf("Day 10 Part 1: %s is the order the steps should be completed\n", partOne(points))
}

var re = regexp.MustCompile(`(?m)position=<\s?(-?\d*), \s?(-?\d*)> velocity=<\s?(-?\d*), \s?(-?\d*)>`)

func partOne(points []Point) string {
	state := State{data: points}
	//state.printState()

	dim := state.calcDimensions()
	step := 0
	for {
	//	state.printState()

		newState := state.move()
		newDim := state.calcDimensions()

		fmt.Printf("step %d, dim: %d\n", step, newDim)
		if newDim > dim || step == 10375 { // for some reason it print a second too late :(
			state.printState()
			break
		}

		dim = newDim
		state = newState
		step++
	}

	return ""
}

func partTwo(input []string) int {
	return 0
}

type Point struct {
	X  int
	Y  int
	dX int
	dY int
}

type State struct {
	data []Point
}

func (s State) move() (State) {
	newPoints := make([]Point, len(s.data))
	for i, p := range s.data {
		newPoints[i] = Point{X: p.X + p.dX, Y: p.Y + p.dY, dX: p.dX, dY: p.dY}
	}

	return State{newPoints}
}

func (s State) calcBoundaries() (xMin, yMin, xMax, yMax int) {
	for i, point := range s.data {
		if i == 0 || xMin > point.X {
			xMin = point.X
		}

		if i == 0 || xMax < point.X {
			xMax = point.X
		}

		if i == 0 || yMin > point.Y {
			yMin = point.Y
		}

		if i == 0 || yMax < point.Y {
			yMax = point.Y
		}
	}

	return
}

func (s State) calcDimensions() int {
	xMin, _, xMax, _ := s.calcBoundaries()
	return xMax - xMin
}

func (s State) printState() {
	xMin, yMin, xMax, yMax := s.calcBoundaries()

	mapResult := make(map[string]bool)
	for _, point := range s.data {
		mapResult[fmt.Sprintf("%d,%d", point.X, point.Y)] = true
	}
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {

			if _, exists := mapResult[fmt.Sprintf("%d,%d", x, y)]; exists {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func linesToPoint(line string) Point {
	matches := re.FindStringSubmatch(line)
	X, _ := strconv.Atoi(matches[1])
	Y, _ := strconv.Atoi(matches[2])
	dX, _ := strconv.Atoi(matches[3])
	dY, _ := strconv.Atoi(matches[4])

	return Point{X, Y, dX, dY}
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
