package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part2("day8.input")
}

func part1(inputFile string) {
	treematrix := ParseTreeMatrix(inputFile)
	var matrixSize = len(treematrix)

	// lets find now how many tress are found
	visible := matrixSize*4 - 4 // all the edges
	for i := 1; i < matrixSize-1; i += 1 {
		for j := 1; j < matrixSize-1; j += 1 {
			left, up, right, down := true, true, true, true
			for m := 0; m < j; m += 1 {
				if treematrix[i][m] >= treematrix[i][j] {
					left = false
					break
				}
			}
			for m := matrixSize - 1; m > j; m -= 1 {
				if treematrix[i][m] >= treematrix[i][j] {
					right = false
					break
				}
			}
			for m := 0; m < i; m += 1 {
				if treematrix[m][j] >= treematrix[i][j] {
					down = false
					break
				}
			}
			for m := matrixSize - 1; m > i; m -= 1 {
				if treematrix[m][j] >= treematrix[i][j] {
					up = false
					break
				}
			}
			if left || down || right || up {
				visible += 1
			}
		}
	}

	fmt.Printf("visible: %v\n", visible)
}

func part2(inputFile string) {
	treematrix := ParseTreeMatrix(inputFile)
	var matrixSize = len(treematrix)

	highestScore := 0
	for i := 1; i < matrixSize-1; i += 1 {
		for j := 1; j < matrixSize-1; j += 1 {
			left, up, right, down := 0, 0, 0, 0
			for m := i - 1; m >= 0; m -= 1 {
				up += 1
				if treematrix[m][j] >= treematrix[i][j] {
					break
				}
			}
			for m := j - 1; m >= 0; m -= 1 {
				left += 1
				if treematrix[i][m] >= treematrix[i][j] {
					break
				}
			}

			for m := i + 1; m < matrixSize; m += 1 {
				down += 1
				if treematrix[m][j] >= treematrix[i][j] {
					break
				}
			}
			for m := j + 1; m < matrixSize; m += 1 {
				right += 1
				if treematrix[i][m] >= treematrix[i][j] {
					break
				}
			}
			score := left * up * right * down
			if score > highestScore {
				highestScore = score
			}
		}
	}

	fmt.Printf("highestScore: %v\n", highestScore)
}

func ParseTreeMatrix(inputFile string) (treematrix [][]int) {
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)

	rowNumber := 0
	for scanner.Scan() {
		line := scanner.Text()

		if treematrix == nil {
			treematrix = make([][]int, len(line))
		}

		chars := []rune(line)
		strArray := []int{}
		for i := 0; i < len(chars); i++ {
			intChar, _ := strconv.Atoi(string(chars[i]))
			strArray = append(strArray, intChar)
		}
		treematrix[rowNumber] = strArray
		rowNumber += 1
	}
	return
}
