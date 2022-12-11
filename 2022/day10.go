package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

type Operation struct {
	instruction string
	increase    int
}

type Screen []rune

func (crt Screen) String() string {
	var builder strings.Builder
	for y := 0; y < 6; y++ { // screen is 6 pixels height
		for x := 0; x < 40; x++ { // screen is 40 pixels wide
			builder.WriteRune(crt[y*40+x])
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

func part1() {
	file, _ := os.Open("day10.input")
	scanner := bufio.NewScanner(file)

	instructions := []Operation{}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		instruction := line[0]
		increase := 0
		if instruction == "addx" {
			increase, _ = strconv.Atoi(line[1])
		}
		instructions = append(instructions, Operation{instruction, increase})
	}

	answer := 0

	x := 1
	i := 0
	midAdd := false

	for cycle := 1; i < len(instructions); cycle++ {
		fmt.Printf("[Cycle: %v] start\n", cycle)
		if cycle%20 == 0 {
			if cycle == 20 || (cycle-20)%40 == 0 {
				answer += x * cycle
				fmt.Printf("[Cycle: %v] Signal Strength: %v\n", cycle, x*cycle)
			}
		}
		if !midAdd {
			switch instructions[i].instruction {
			case "noop":
				i++
			case "addx":
				midAdd = true
			}
		} else {
			midAdd = false
			x += instructions[i].increase
			i++
		}

	}
	fmt.Printf("answer: %v\n", answer)
}

func part2() {
	file, _ := os.Open("day10.input")
	scanner := bufio.NewScanner(file)

	instructions := []Operation{}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		instruction := line[0]
		increase := 0
		if instruction == "addx" {
			increase, _ = strconv.Atoi(line[1])
		}
		instructions = append(instructions, Operation{instruction, increase})
	}

	x := 1
	i := 0
	midAdd := false

	screen := make(Screen, 6*40)
	for i := 0; i < len(screen); i++ {
		screen[i] = '.'
	}

	for cycle := 0; i < len(instructions); cycle++ {
		fmt.Printf("[Cycle: %v] start\n", cycle)
		if x == cycle%40 || x-1 == cycle%40 || x+1 == cycle%40 {
			screen[cycle] = '#'
		}
		if !midAdd {
			switch instructions[i].instruction {
			case "noop":
				i++
			case "addx":
				midAdd = true
			}
		} else {
			midAdd = false
			x += instructions[i].increase
			i++
		}

	}
	fmt.Printf("%v", screen.String())
}
