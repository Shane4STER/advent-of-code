package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	scanner := bufio.NewScanner(openStdinOrFile())

	var total, totalP2 int

	var currentChunk []string

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			currentChunk = append(currentChunk, line)
			continue
		} else {
			total += processChunk(currentChunk)
			totalP2 += processChunkPart2(currentChunk)
			currentChunk = currentChunk[:0]
		}
	}
	total += processChunk(currentChunk)

	totalP2 += processChunkPart2(currentChunk)

	fmt.Printf("The total is %v\n", total)
	fmt.Printf("The total is %v\n", totalP2)
}

func processChunk(chunk []string) int {
	yesAnswers := make(map[rune]bool)
	for _, s := range chunk {
		for _, c := range s {
			yesAnswers[c] = true
		}
	}
	return len(yesAnswers)
}

func processChunkPart2(chunk []string) int {
	yesAnswers := make(map[rune]int)
	for _, s := range chunk {
		for _, c := range s {
			yesAnswers[c]++
		}
	}
	return mapValueIsTargetValue(yesAnswers, len(chunk))
}

func mapValueIsTargetValue(data map[rune]int, targetValue int) int {
	var r int
	for _, v := range data {
		if v == targetValue {
			r++
		}
	}

	return r
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
