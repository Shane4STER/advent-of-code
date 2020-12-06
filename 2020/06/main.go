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
			total += processPaxGroup(currentChunk, false)
			totalP2 += processPaxGroup(currentChunk, true)
			currentChunk = currentChunk[:0]
		}
	}
	total += processPaxGroup(currentChunk, false)

	totalP2 += processPaxGroup(currentChunk, true)

	fmt.Printf("The total is %v\n", total)
	fmt.Printf("The total is %v\n", totalP2)
}

func processPaxGroup(chunk []string, countGroupAnswersOnly bool) int {
	yesAnswers := make(map[rune]int)
	for _, s := range chunk {
		for _, c := range s {
			yesAnswers[c]++
		}
	}
	if countGroupAnswersOnly {
		return mapValueIsTargetValue(yesAnswers, len(chunk))
	} else {
		return len(yesAnswers)
	}
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
