package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const moveX = 3
const moveY = 1

type coord struct {
	x int
	y int
}

type slope struct {
	dx int
	dy int
}

func main() {
	scanner := bufio.NewScanner(openStdinOrFile())

	var mapTemplate []string

	product := 1

	for scanner.Scan() {
		mapTemplate = append(mapTemplate, scanner.Text())
	}

	slopes := []slope{
		slope{1, 1},
		slope{3, 1},
		slope{5, 1},
		slope{7, 1},
		slope{1, 2},
	}
	for _, chosenSlope := range slopes {
		treeCount := countTrees(mapTemplate, chosenSlope)
		fmt.Printf("The total number of trees encountered for %v is %v\n", chosenSlope, treeCount)
		product = product * treeCount
	}
	fmt.Printf("The total product of all tree counts is %v\n", product)

}

func countTrees(mapTemplate []string, chosenSlope slope) int {
	var treeCount int
	var targetPos coord

	targetPos = nextPos(targetPos, chosenSlope)

	for y, line := range mapTemplate {
		if y < targetPos.y {
			continue
		}

		relativeX := targetPos.x % len(line)

		if line[relativeX] == '#' {
			treeCount++
		}

		targetPos = nextPos(targetPos, chosenSlope)
	}

	return treeCount
}

func nextPos(currentPos coord, delta slope) coord {
	return coord{
		currentPos.x + delta.dx,
		currentPos.y + delta.dy,
	}
}

func openStdinOrFile() io.Reader {
	var err error
	r := os.Stdin
	if len(os.Args) > 1 {
		r, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	return r
}
