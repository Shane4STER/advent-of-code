package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type targetWindow struct {
	a int
	b int
	c int
}

func main() {
	previous := 0
	previous2 := 0
	increases := 0

	windows := []targetWindow{}

	scanner := bufio.NewScanner(openStdinOrFile())

	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if depth > previous && previous != 0 {
			increases += 1
		}

		if previous2 != 0 {
			windows = append(windows, targetWindow{previous2, previous, depth})
		}

		previous2 = previous
		previous = depth
	}

	fmt.Printf("The total number of increases is %v\n", increases)

	windowIncreases := 0
	previousSum := 0
	for i, window := range windows {
		currentSum := window.sum()

		if i == 0 {
			previousSum = currentSum
			continue
		}

		if currentSum > previousSum {
			windowIncreases += 1
		}

		previousSum = currentSum
	}

	fmt.Printf("The total number of windowed increases is %v\n", windowIncreases)
}

func (tw *targetWindow) sum() int {
	return tw.a + tw.b + tw.c
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
