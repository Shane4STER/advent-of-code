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
	var data []string
	var initialised bool
	var bitWidth int
	var gammaStr, epsilonStr string

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
		if !initialised {
			bitWidth = len(line)
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

	oxyStr := filterPrefix(true, data)
	co2Str := filterPrefix(false, data)

	oxyRate, err := strconv.ParseUint(oxyStr, 2, bitWidth)
	if err != nil {
		log.Fatal(err)
	}
	co2Rate, err := strconv.ParseUint(co2Str, 2, bitWidth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The Oxygen Generator Rating is %v = %v\n", oxyStr, oxyRate)
	fmt.Printf("The CO2 Scrubber Rating is %v = %v\n", co2Str, co2Rate)
	fmt.Printf("The Life Support Rating is %v\n", oxyRate*co2Rate)
}

func filterPrefix(useFrequent bool, haystack []string) string {
	var diff int
	var zeroStack []string
	var oneStack []string

	for _, s := range haystack {
		if s[0] == '0' {
			zeroStack = append(zeroStack, s[1:])
			diff++
		} else {
			oneStack = append(oneStack, s[1:])
			diff--
		}
	}

	if useFrequent {
		if diff > 0 {
			if len(zeroStack) > 1 {
				return "0" + filterPrefix(useFrequent, zeroStack)
			} else {
				return "0" + zeroStack[0]
			}
		} else {
			if len(oneStack) > 1 {
				return "1" + filterPrefix(useFrequent, oneStack)
			} else {
				return "1" + oneStack[0]
			}
		}
	} else {
		if diff <= 0 {
			if len(zeroStack) > 1 {
				return "0" + filterPrefix(useFrequent, zeroStack)
			} else {
				return "0" + zeroStack[0]
			}
		} else {
			if len(oneStack) > 1 {
				return "1" + filterPrefix(useFrequent, oneStack)
			} else {
				return "1" + oneStack[0]
			}
		}
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
