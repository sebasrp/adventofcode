package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	part2()
}

func part1() {
	file, _ := os.Open("day3.input")

	scanner := bufio.NewScanner(file)
	items := []rune{}
	for scanner.Scan() {
		ruckstack := scanner.Text()
		sizeR := len(ruckstack)
		first := ruckstack[0:(sizeR / 2)]
		second := ruckstack[sizeR/2 : sizeR]
		fmt.Printf("first: %v, second: %v\n", first, second)

		firstMap := map[string]int{}
		for _, letter := range first {
			firstMap[string(letter)] = 1 // value does not matter
		}

		wrongItemsRuckstack := map[rune]int{}
		for _, letter := range second {
			if _, ok := firstMap[string(letter)]; ok {
				wrongItemsRuckstack[letter] = 1
				fmt.Printf("problematic item found: %v\n", string(letter))
			}
		}
		for k := range wrongItemsRuckstack {
			items = append(items, k)
		}
	}
	fmt.Printf("problematic items: %v\n", items)
	sumPriorities := 0
	for _, item := range items {
		sumPriorities += alphaNum(item)
		fmt.Printf("item %v (priority %v)\n", string(item), alphaNum(item))
	}
	fmt.Printf("sum prios: %v\n", sumPriorities)
}

func part2() {
	file, _ := os.Open("day3.input")

	scanner := bufio.NewScanner(file)
	group := []string{}
	sumPriorities := 0

	for scanner.Scan() {
		ruckstack := scanner.Text()
		fmt.Printf("ruckstack: %v\n", ruckstack)
		group = append(group, ruckstack)

		if len(group) == 3 {
			fmt.Printf("looking for badge in group: %v\n", group)

			firstMap := map[string]int{}
			for _, letter := range group[0] {
				firstMap[string(letter)] = 1 // value does not matter
			}

			firstSecond := map[rune]int{}
			for _, letter := range group[1] {
				if _, ok := firstMap[string(letter)]; ok {
					firstSecond[letter] = 1
				}
			}

			var badge rune
			for _, letter := range group[2] {
				if _, ok := firstSecond[letter]; ok {
					badge = letter
				}
			}
			fmt.Printf("badge: %v\n", string(badge))
			sumPriorities += alphaNum(badge)
			group = []string{} // we restart
		}
	}
	fmt.Printf("sum prios: %v\n", sumPriorities)

}

func alphaNum(letter rune) int {
	// check ascii codes http://sticksandstones.kstrom.com/appen.html
	char := int(letter)
	if unicode.IsLower(letter) {
		char -= 96
	} else {
		char -= 38
	}
	return char
}
