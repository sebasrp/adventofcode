package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type efile struct {
	name     string
	size     int
	parent   *efile
	children map[string]*efile
}

func (f *efile) Size() (size int) {
	if f.children == nil {
		return f.size
	}
	for _, d := range f.children {
		size += d.Size()
	}
	return
}

func main() {
	part2()
}

func part1() {
	file, _ := os.Open("day7.input")

	scanner := bufio.NewScanner(file)

	var current *efile

	tree := []*efile{}
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		if len(line) > 2 { // a command
			if line[2] == ".." {
				current = current.parent
			} else if line[2] == "/" { // root folder
				current = &efile{name: "/", size: 0, parent: nil, children: make(map[string]*efile)}
			} else {
				current = current.children[line[2]]
			}
		} else if line[0] == "dir" {
			current.children[line[1]] = &efile{name: line[1], size: 0, parent: current, children: make(map[string]*efile)}
			tree = append(tree, current.children[line[1]])
		} else if line[0] != "$" {
			size, _ := strconv.Atoi(line[0])
			current.children[line[1]] = &efile{name: line[1], size: size, parent: current, children: nil}
		}
	}

	result := 0

	for _, dir := range tree {
		size := dir.Size()
		fmt.Printf("Dir %v, size: %v\n", dir.name, size)
		if size <= 100000 {
			result += size
		}
	}

	fmt.Println(result)
}

func part2() {
	file, _ := os.Open("day7.input")

	scanner := bufio.NewScanner(file)

	var current *efile

	tree := []*efile{}
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		if len(line) > 2 { // a command
			if line[2] == ".." {
				current = current.parent
			} else if line[2] == "/" { // root folder
				current = &efile{name: "/", size: 0, parent: nil, children: make(map[string]*efile)}
				tree = append(tree, current)
			} else {
				current = current.children[line[2]]
			}
		} else if line[0] == "dir" {
			current.children[line[1]] = &efile{name: line[1], size: 0, parent: current, children: make(map[string]*efile)}
			tree = append(tree, current.children[line[1]])
		} else if line[0] != "$" {
			size, _ := strconv.Atoi(line[0])
			current.children[line[1]] = &efile{name: line[1], size: size, parent: current, children: nil}
		}
	}

	toFree := 30000000 - (70000000 - tree[0].Size())
	result := tree[0].Size() // we start with root

	fmt.Printf("Need to free: %v\n", toFree)
	for _, dir := range tree {
		size := dir.Size()
		fmt.Printf("%v,%v\n", size, dir.name)
		if size > toFree && (size < result) {
			result = size
		}
	}
	fmt.Println(result)
}
