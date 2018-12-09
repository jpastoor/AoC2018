package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Day 9 Games!\n")
	games := []Game{
		{9, 25},
		{10, 1618},
		{13, 7999},
		{17, 1104},
		{21, 6111},
		{30, 5807},
		{452, 7078400},
	}

	highScore := 0
	for _, game := range games {
		start :=  time.Now()
		highScore = playGame(game.players, game.lastMarble)
		end := time.Now()
		fmt.Printf("%d players; last marble is worth %d points: high score is %d. Took %s\n", game.players, game.lastMarble, highScore, end.Sub(start))
	}
}

/**
Im using a (circular linked list structure)
 */
type ListItem struct {
	next *ListItem
	prev *ListItem
	value int
}

func (item *ListItem) insertAfter(elem int) *ListItem {
	newElem := &ListItem{prev: item, next: item.next, value: elem}
	item.next.prev = newElem
	item.next = newElem
	return newElem
}

func (item *ListItem) remove() *ListItem {
	item.next.prev = item.prev
	item.prev.next = item.next
	return item.next
}

type Game struct {
	players int
	lastMarble int
}

func playGame(players, lastMarble int) int {
	// Make a list of remaining marbles to play
	var remainingMarbles []int
	for i := 1; i <= lastMarble; i++ {
		remainingMarbles = append(remainingMarbles, i)
	}

	// Scoreboard
	scores := make(map[int]int)
	for i := 1; i <= players; i++ {
		scores[i] = 0
	}

	first := &ListItem{ value:0}
	current := first
	// This part makes the list circular
	current.prev = current
	current.next = current

	currentPlayer := 1
	for turn := 1; turn <= lastMarble; turn++ {
		nextMarble := remainingMarbles[0]
		specialTurn := nextMarble % 23 == 0

		// Normal case
		if !specialTurn {
			current = current.next.insertAfter(nextMarble)
		} else {
			// Special turn! We go back a few steps
			current = current.prev.prev.prev.prev.prev.prev.prev

			// We need to remember the value of the marble we are removing to add it to the score later
			removedMarble := current.value

			// Remove marble
			current = current.remove()

			// Update scoreboard
			scores[currentPlayer] += nextMarble + removedMarble
		}

		remainingMarbles = remainingMarbles[1:]

		currentPlayer++
		if currentPlayer > players {
			currentPlayer = 1
		}
	}

	// Calculate the highscore
	highScore := 0
	for _, score := range scores {
		if score > highScore {
			highScore = score
		}
	}

	return highScore
}