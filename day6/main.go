package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := parseLines("day6/input.txt")
	demo := parseLines("day6/demo.txt")

	fmt.Printf("Day 6 Part Demo: %d is the size of the largest area that isn't infinite\n", partOne(demo))
	fmt.Printf("Day 6 Part 1: %d is the size of the largest area that isn't infinite\n", partOne(input))
	fmt.Printf("Day 6 Part 2: %d  is the size of the region containing all locations which have a total distance to all given coordinates of less than 10000", partTwo(input))
}

func partOne(input []string) int {
	gridMinX, gridMinY, gridMaxX, gridMaxY, points := parseGrid(input)

	// Loop over the grid
	for x := gridMinX - 1; x <= gridMaxX+1; x++ {
		for y := gridMinY - 1; y <= gridMaxY+1; y++ {
			location := Coords{x, y}

			// Check for everyone how far away they are from this coord
			minDistance := 1000000
			minDistanceId := 0
			minDistanceMultiple := false

			distances := make(map[int]int)
			for _, point := range points {
				distance := distanceCoords(point.X, location.X, point.Y, location.Y)
				distances[point.Id] = distance

				if minDistance > distance {
					minDistanceId = point.Id
					minDistance = distance
					minDistanceMultiple = false
				} else if minDistance == distance {
					minDistanceMultiple = true
				}
			}

			if minDistanceMultiple == false {
				points[minDistanceId].Counter++

				if location.X == gridMinX-1 || location.X == gridMaxX+1 || location.Y == gridMinY-1 || location.Y == gridMaxY+1 {
					points[minDistanceId].IsInfinite = true
				}
			}
		}
	}

	largestAreaSize := 0
	for _, point := range points {
		//fmt.Println(point.Id, point.IsInfinite, point.Counter)
		if !point.IsInfinite && largestAreaSize < point.Counter {
			largestAreaSize = point.Counter
		}
	}

	return largestAreaSize
}

func parseGrid(input []string) (int, int, int, int, map[int]*Pair) {
	gridMinX := 0
	gridMinY := 0
	gridMaxX := 0
	gridMaxY := 0
	points := make(map[int]*Pair)
	for i, line := range input {
		id := i + 1
		split := strings.Split(line, ", ")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		points[id] = &Pair{X: x, Y: y, Id: id, Counter: 0}

		if gridMinX > x || i == 0 {
			gridMinX = x
		}
		if gridMaxX < x || i == 0 {
			gridMaxX = x
		}
		if gridMinY > y || i == 0 {
			gridMinY = y
		}
		if gridMaxY < y || i == 0 {
			gridMaxY = y
		}
	}
	return gridMinX, gridMinY, gridMaxX, gridMaxY, points
}

func distanceCoords(x1, x2, y1, y2 int) int {
	xD := x1 - x2
	if xD < 0 {
		xD *= -1
	}

	yD := y1 - y2
	if yD < 0 {
		yD *= -1
	}

	return xD + yD
}

func FindCoordsSurroundingPoint(coords Coords, distance int) []Coords {
	var output []Coords

	// Top + Bottom
	for x := coords.X - distance; x <= coords.X+distance; x++ {
		output = append(output, Coords{x, coords.Y - distance})
		output = append(output, Coords{x, coords.Y + distance})
	}

	// Left + Right
	for y := coords.Y - (distance - 1); y <= coords.Y+(distance-1); y++ {
		output = append(output, Coords{coords.X - distance, y})
		output = append(output, Coords{coords.X - distance, y})
	}

	return output
}

func partTwo(input []string) int {
	gridMinX, gridMinY, gridMaxX, gridMaxY, points := parseGrid(input)
	regionSize := 0
	// Loop over the grid
	for x := gridMinX - 1; x <= gridMaxX+1; x++ {
		for y := gridMinY - 1; y <= gridMaxY+1; y++ {
			location := Coords{x, y}

			// Check for everyone how far away they are from this coord
			sumDistance := 0
			for _, point := range points {
				sumDistance += distanceCoords(point.X, location.X, point.Y, location.Y)
			}

			if sumDistance < 10000 {
				regionSize++
			}
		}
	}

	return regionSize
}

type Coords struct {
	X int
	Y int
}

type Pair struct {
	X          int
	Y          int
	Id         int
	Counter    int
	IsInfinite bool
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
