package main

import (
	"fmt"
)

func main() {

	// TestPowerLevel
	fmt.Println(calculatePowerLevel(3, 5, 8))
	fmt.Println(calculatePowerLevel(122, 79, 57))
	fmt.Println(calculatePowerLevel(217, 196, 39))
	fmt.Println(calculatePowerLevel(101, 153, 71))

	demoX, demoY, demoPower:= partOne(18, 300)
	fmt.Printf("Day 11 Demo: (%d,%d) is the coordinate the top-left fuel cell of the 3x3 square with the largest total power (%d)\n", demoX, demoY, demoPower)
	x, y, power := partOne(2694, 300)
	fmt.Printf("Day 11 Part 1: (%d,%d) is the coordinate the top-left fuel cell of the 3x3 square with the largest total power (%d)\n", x, y, power)
	x2, y2, size2, power2 := partTwo(2694, 300)
	fmt.Printf("Day 11 Part 2: (%d,%d,%d) is the coordinate the top-left fuel cell of the 3x3 square with the largest total power (%d)\n", x2, y2, size2, power2)

}

func partOne(gridSerialNumber int, gridSize int) (clusterX, clusterY, clusterPower int) {
	squareSize := 3

	for y := 1; y <= gridSize-squareSize; y++ {
		for x := 1; x <= gridSize-squareSize; x++ {

			powerSum := 0
			for xS := 0; xS < squareSize; xS++ {
				for yS := 0; yS < squareSize; yS++ {
					powerSum += calculatePowerLevel(x+xS, y+yS, gridSerialNumber)
				}
			}

			if powerSum > clusterPower {
				clusterX, clusterY = x, y
				clusterPower = powerSum
			}
		}
	}

	return clusterX, clusterY, clusterPower
}

func partTwo(gridSerialNumber int, gridSize int) (clusterX, clusterY, squareSize int, clusterPower int) {
	for size := 1; size <= gridSize; size++ {
		for y := 1; y <= gridSize-size; y++ {
			for x := 1; x <= gridSize-size; x++ {

				powerSum := 0
				for xS := 0; xS < size; xS++ {
					for yS := 0; yS < size; yS++ {
						powerSum += calculatePowerLevel(x+xS, y+yS, gridSerialNumber)
					}
				}

				if powerSum > clusterPower {
					clusterX, clusterY = x, y
					clusterPower = powerSum
					squareSize = size
				}
			}
		}
	}

	return clusterX, clusterY, squareSize, clusterPower
}

func calculatePowerLevel(x, y, gridSerialNumber int) int {
	rackId := x + 10
	powerLevel := rackId * y
	powerLevel += gridSerialNumber
	powerLevel *= rackId
	powerLevel = (powerLevel / 100) % 10
	return powerLevel - 5
}
