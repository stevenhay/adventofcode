package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.OpenFile("input.txt", os.O_RDONLY, 0755)
	defer input.Close()

	scanner := bufio.NewScanner(input)

	// A = Rock
	// B = Paper
	// C = Scissors
	//
	// X = Rock (1)
	// Y = Paper (2)
	// Z = Scissors (3)

	// X = Should lose
	// Y = Should draw
	// Z = Should win

	// Win = 6
	// Draw = 3
	// Lose = 0

	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		// Should lose
		if tokens[1] == "X" {
			totalScore += 0

			if tokens[0] == "A" {
				totalScore += 3
			} else if tokens[0] == "B" {
				totalScore += 1
			} else if tokens[0] == "C" {
				totalScore += 2
			}
		}

		// Should draw
		if tokens[1] == "Y" {
			totalScore += 3

			if tokens[0] == "A" {
				totalScore += 1
			} else if tokens[0] == "B" {
				totalScore += 2
			} else if tokens[0] == "C" {
				totalScore += 3
			}
		}

		// Should win
		if tokens[1] == "Z" {
			totalScore += 6

			if tokens[0] == "A" {
				totalScore += 2
			} else if tokens[0] == "B" {
				totalScore += 3
			} else if tokens[0] == "C" {
				totalScore += 1
			}
		}
		/*
			if tokens[0] == "A" && tokens[1] == "X" {
				totalScore += 1
				totalScore += 3
				// draw
			}
			if tokens[0] == "A" && tokens[1] == "Y" {
				totalScore += 2
				totalScore += 6
				// win
			}
			if tokens[0] == "A" && tokens[1] == "Z" {
				totalScore += 3
				totalScore += 0
				// lose
			}

			if tokens[0] == "B" && tokens[1] == "X" {
				totalScore += 1
				totalScore += 0
				// lose
			}
			if tokens[0] == "B" && tokens[1] == "Y" {
				totalScore += 2
				totalScore += 3
				// draw
			}
			if tokens[0] == "B" && tokens[1] == "Z" {
				totalScore += 3
				totalScore += 6
				// win
			}

			if tokens[0] == "C" && tokens[1] == "X" {
				totalScore += 1
				totalScore += 6
				// win
			}
			if tokens[0] == "C" && tokens[1] == "Y" {
				totalScore += 2
				totalScore += 0
				// lose
			}
			if tokens[0] == "C" && tokens[1] == "Z" {
				totalScore += 3
				totalScore += 3
				// draw
			}*/
	}

	fmt.Printf("Total Points: %d", totalScore)
}
