package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(openStdinOrFile())

	var freq0, freq1 []int
	var initialised bool
	var bitWidth int
	var gammaStr, epsilonStr string

	for scanner.Scan() {
		line := scanner.Text()
		if !initialised {
			bitWidth = len(scanner.Text())
			freq0 = make([]int, bitWidth)
			freq1 = make([]int, bitWidth)
			initialised = true
		}

		for i, c := range line {
			if c == '0' {
				freq0[i] += 1
			} else {
				freq1[i] += 1
			}
		}
	}

	fmt.Printf("The number of bits is %v\n", bitWidth)

	for i := 0; i < bitWidth; i++ {
		if freq0[i] > freq1[i] {
			gammaStr = gammaStr + "0"
			epsilonStr = epsilonStr + "1"
		} else {
			gammaStr = gammaStr + "1"
			epsilonStr = epsilonStr + "0"
		}
	}

	gammaRate, err := strconv.ParseUint(gammaStr, 2, bitWidth)
	if err != nil {
		log.Fatal(err)
	}
	epsilonRate, err := strconv.ParseUint(epsilonStr, 2, bitWidth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The Gamma Rate is %v = %v\n", gammaStr, gammaRate)
	fmt.Printf("The Epsilon Rate is %v = %v\n", epsilonStr, epsilonRate)
	fmt.Printf("The Power Consumption is %v\n", epsilonRate*gammaRate)
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
