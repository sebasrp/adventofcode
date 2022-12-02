package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part2()
}

func part1() {
	// A: Rock, B: Paper, C: Scissors
	// X: Rock, Y: Paper, Z: Scissors
	roundValues := map[string]int{
		"AX": 3,
		"AY": 6,
		"AZ": 0,
		"BX": 0,
		"BY": 3,
		"BZ": 6,
		"CX": 6,
		"CY": 0,
		"CZ": 3,
	}
	shapePoints := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	file, _ := os.Open("day2.input")

	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		hands := strings.Join(strings.Fields(scanner.Text()), "")
		round := roundValues[hands] + shapePoints[hands[1:2]]
		score = score + round
		fmt.Printf("%v, %v, %v, %v\n", hands[0:1], hands[1:2], roundValues[hands], shapePoints[hands[1:2]])
	}
	fmt.Printf("total score: %v\n", score)
}

func part2() {
	// A: Rock, B: Paper, C: Scissors
	roundValues := map[string]string{
		"AX": "C",
		"AY": "A",
		"AZ": "B",
		"BX": "A",
		"BY": "B",
		"BZ": "C",
		"CX": "B",
		"CY": "C",
		"CZ": "A",
	}
	shapePoints := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	// X: lose, Y: draw, Z: win
	outcomePoint := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	file, _ := os.Open("day2.input")

	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		entry := strings.Join(strings.Fields(scanner.Text()), "")
		score = score + outcomePoint[entry[1:2]] + shapePoints[roundValues[entry]]
		//fmt.Printf("%v, %v, %v\n", entry, outcomePoint[entry[1:2]], shapePoints[roundValues[entry]])
	}
	fmt.Printf("total score: %v\n", score)
}
