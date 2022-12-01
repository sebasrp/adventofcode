package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	part2()
}

func part1() {
	file, err := os.Open("day1.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var current_max int64 = 0
	var current_sum int64 = 0
	for scanner.Scan() {
		var line = scanner.Text()
		if line == "" {
			if current_sum > current_max {
				current_max = current_sum
			}
			current_sum = 0
		} else {
			var lineInt, _ = strconv.ParseInt(line, 10, 32)
			current_sum = current_sum + lineInt
		}
	}

	fmt.Printf("max: %v\n", current_max)
}

func part2() {
	file, err := os.Open("day1.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sums := []int{}
	var current_sum int64 = 0
	for scanner.Scan() {
		var line = scanner.Text()
		if line == "" {
			sums = append(sums, int(current_sum))
			fmt.Printf("%v\n", current_sum)
			current_sum = 0
		} else {
			var lineInt, _ = strconv.ParseInt(line, 10, 32)
			current_sum = current_sum + lineInt
		}
	}
	sums = append(sums, int(current_sum)) //add the last item at the end

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	top3 := sums[0] + sums[1] + sums[2]

	fmt.Printf("top3: %v\n", top3)
}
