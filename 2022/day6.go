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
	file, _ := os.Open("day6.test.input")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := scanner.Text()
		fmt.Printf("index: %v\n", trailUnique(input, 4))
	}
}

func part2() {
	file, _ := os.Open("day6.input")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := scanner.Text()
		fmt.Printf("index: %v\n", trailUnique(input, 14))
	}
}

func trailUnique(input string, uniqueCount int) (index int) {
	index = -1
	runes := []rune(input)
	lastn := ""
	for i := 0; i < len(input); i++ {
		current := string(runes[i])
		if strings.Contains(lastn, current) {
			lastn = string(runes[strings.LastIndex(input[0:i], current)+1 : i+1])
		} else {
			lastn += current
			if len(lastn) > uniqueCount {
				lastn = lastn[uniqueCount:]
			}
		}

		if len(lastn) == uniqueCount {
			index = i + 1
			break
		}
	}
	return
}
