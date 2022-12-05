package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1(9, "day5.input")
	part2(9, "day5.input")
}

func part1(stackCount int, inputFile string) {
	file, _ := os.Open(inputFile)

	scanner := bufio.NewScanner(file)
	stacks := map[int][]string{} // key is stack number, valye is list of crates

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if line[0] != 'm' { // ugly to find stack entry vs move... but it works
			for i := 0; i < stackCount; i++ {
				crateString := line[i*4 : i*4+3]
				crate, found := crateName(crateString)
				if found {
					stacks[i+1] = append(stacks[i+1], crate)
					fmt.Printf("crate: %v in stack %v\n", crate, i+1)
				}
			}
		} else { // these are the crate moves. i know. it's ugly
			cratesCount, from, to := extractMoves(line)
			for i := 0; i < cratesCount; i++ {
				if len(stacks[from]) == 0 {
					continue
				} else {
					crateToMove := stacks[from][0]
					stacks[from] = stacks[from][1:]
					stacks[to] = append([]string{crateToMove}, stacks[to]...)
				}
			}
			fmt.Printf("move: %v from %v to %v\n", cratesCount, from, to)
		}
	}
	fmt.Printf("stacks: %v\n", stacks)
	topCratesString := ""
	for i := 1; i < stackCount+1; i++ {
		topCratesString += stacks[i][0]
	}
	fmt.Printf("message: %v\n", topCratesString)
}

func part2(stackCount int, inputFile string) {
	file, _ := os.Open(inputFile)

	scanner := bufio.NewScanner(file)
	stacks := map[int][]string{} // key is stack number, valye is list of crates

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if line[0] != 'm' { // ugly to find stack entry vs move... but it works
			for i := 0; i < stackCount; i++ {
				crateString := line[i*4 : i*4+3]
				crate, found := crateName(crateString)
				if found {
					stacks[i+1] = append(stacks[i+1], crate)
				}
			}
		} else { // these are the crate moves. i know. it's ugly
			cratesCount, from, to := extractMoves(line)
			if len(stacks[from]) == 0 {
				continue
			} else {
				cpyFrom := make([]string, len(stacks[from]))
				copy(cpyFrom, stacks[from])
				stacks[to] = append(stacks[from][0:cratesCount], stacks[to]...)
				stacks[from] = cpyFrom[cratesCount:]
			}
		}
	}
	fmt.Printf("stacks: %v\n", stacks)
	topCratesString := ""
	for i := 1; i < stackCount+1; i++ {
		topCratesString += stacks[i][0]
	}
	fmt.Printf("message: %v\n", topCratesString)
}

func crateName(crateString string) (result string, found bool) {
	s := strings.Index(crateString, "[")
	e := strings.Index(crateString, "]")
	if s == -1 || e == -1 {
		return result, false
	}
	return crateString[s+1 : e], true
}

func extractMoves(movesString string) (cratesCount int, from int, to int) {
	words := strings.Fields(movesString)
	cratesCount, _ = strconv.Atoi(words[1])
	from, _ = strconv.Atoi(words[3])
	to, _ = strconv.Atoi(words[5])
	return
}
