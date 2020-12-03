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

func main() {
	scanner := bufio.NewScanner(openStdinOrFile())

	var currentPos, targetPos coord
	var treeCount int

	targetPos.x = moveX
	targetPos.y = moveY

	for scanner.Scan() {
		if currentPos.y < targetPos.y {
			currentPos.y++
			continue
		}
		currentLine := scanner.Text()
		relativeX := targetPos.x % len(currentLine)
		if currentLine[relativeX] == '#' {
			treeCount++
		}

		currentPos = targetPos

		targetPos.x = currentPos.x + moveX
		targetPos.y = currentPos.y + moveY
		currentPos.y++
	}

	fmt.Printf("The total number of trees encountered is %v", treeCount)

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
