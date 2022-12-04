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

func part1() {
	file, _ := os.Open("day4.input")

	scanner := bufio.NewScanner(file)

	reconsiderCount := 0
	for scanner.Scan() {
		elfSections := strings.Split(scanner.Text(), ",")
		firstStart, firstEnd := SectionsStartEnd(elfSections[0])
		secondStart, secondEnd := SectionsStartEnd(elfSections[1])
		if (secondStart >= firstStart && secondEnd <= firstEnd) || (firstStart >= secondStart && firstEnd <= secondEnd) {
			reconsiderCount += 1
		}
	}
	fmt.Printf("to reconsider: %v\n", reconsiderCount)
}

func part2() {
	file, _ := os.Open("day4.input")

	scanner := bufio.NewScanner(file)

	reconsiderCount := 0
	for scanner.Scan() {
		elfSections := strings.Split(scanner.Text(), ",")
		firstStart, firstEnd := SectionsStartEnd(elfSections[0])
		secondStart, secondEnd := SectionsStartEnd(elfSections[1])
		if (secondStart >= firstStart && secondStart <= firstEnd) || (firstStart >= secondStart && firstStart <= secondEnd) {
			reconsiderCount += 1
		}
	}
	fmt.Printf("to reconsider: %v\n", reconsiderCount)
}

func SectionsStartEnd(sectionString string) (int, int) {
	sections := strings.Split(sectionString, "-")
	start, _ := strconv.Atoi(sections[0])
	end, _ := strconv.Atoi(sections[1])
	return start, end
}
