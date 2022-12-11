package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

var commands = map[string]Point{
	"U": {X: 0, Y: 1},
	"R": {X: 1, Y: 0},
	"L": {X: -1, Y: 0},
	"D": {X: 0, Y: -1},
}

func main() {
	part2()
}

func part1() {
	file, _ := os.Open("day9.input")
	scanner := bufio.NewScanner(file)

	head := Point{X: 0, Y: 0}
	tail := Point{X: 0, Y: 0}
	visited := Set[Point]{}
	visited.Add(tail)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		command := line[0]
		steps, _ := strconv.Atoi(line[1])
		fmt.Printf("command: %v, steps: %v\n", command, steps)

		for i := 0; i < steps; i += 1 {
			head.X += commands[command].X
			head.Y += commands[command].Y
			if !isAdjacent(head, tail) {
				if head.X > tail.X {
					tail.X++
				}
				if head.X < tail.X {
					tail.X--
				}
				if head.Y > tail.Y {
					tail.Y++
				}
				if head.Y < tail.Y {
					tail.Y--
				}

			}
			fmt.Printf("head: %v, tail: %v\n", head, tail)
			visited.Add(tail)
		}
	}
	//fmt.Printf("positions visited: %v\n", visited.List())
	fmt.Printf("position count: %v\n", len(visited.List()))
}

func part2() {
	file, _ := os.Open("day9.input")
	scanner := bufio.NewScanner(file)

	rope := []Point{}
	for i := 0; i < 10; i++ {
		rope = append(rope, Point{X: 0, Y: 0})
	}

	visited := Set[Point]{}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		command := line[0]
		steps, _ := strconv.Atoi(line[1])
		fmt.Printf("command: %v, steps: %v\n", command, steps)

		for i := 0; i < steps; i += 1 {
			rope[0].X += commands[command].X
			rope[0].Y += commands[command].Y
			for seg := 0; seg < len(rope)-1; seg++ {
				leader := rope[seg]
				follower := rope[seg+1]
				if !isAdjacent(leader, follower) {
					if leader.X > follower.X {
						follower.X++
					}
					if leader.X < follower.X {
						follower.X--
					}
					if leader.Y > follower.Y {
						follower.Y++
					}
					if leader.Y < follower.Y {
						follower.Y--
					}

				}
				rope[seg] = leader
				rope[seg+1] = follower
			}

			visited.Add(rope[len(rope)-1])
			fmt.Printf("tail added: %v\n", rope[len(rope)-1])
		}
	}
	//fmt.Printf("positions visited: %v\n", visited.List())
	fmt.Printf("position count: %v\n", len(visited.List()))
}

func isAdjacent(head Point, tail Point) bool {
	xDistance := head.X - tail.X
	if xDistance < 0 {
		xDistance = -xDistance
	}
	if xDistance > 1 {
		return false
	}
	yDistance := head.Y - tail.Y
	if yDistance < 0 {
		yDistance = -yDistance
	}
	return yDistance <= 1

}

// todo: move to its own
type Set[T comparable] map[T]struct{}

func (s Set[T]) String() string {
	items := []string{}
	for k := range s {
		items = append(items, fmt.Sprintf("%v", k))
	}
	return fmt.Sprintf("{%s}", strings.Join(items, ", "))
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Remove(val T) {
	delete(s, val)
}

func (s Set[T]) Contains(val T) (contains bool) {
	_, contains = s[val]
	return
}

func (s Set[T]) Equal(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for k := range s {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	newSet := Set[T]{}
	if len(other) < len(s) {
		s, other = other, s
	}
	for k := range s {
		if other.Contains(k) {
			newSet[k] = struct{}{}
		}
	}
	return newSet
}

func NewSet[T comparable](values []T) Set[T] {
	set := Set[T]{}
	for _, v := range values {
		set[v] = struct{}{}
	}
	return set
}

func SetFromValues[T comparable](values ...T) Set[T] {
	set := Set[T]{}
	for _, v := range values {
		set[v] = struct{}{}
	}
	return set
}

func (s Set[T]) List() []T {
	result := []T{}
	for x := range s {
		result = append(result, x)
	}
	return result
}
