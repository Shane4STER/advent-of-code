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
	index := -1
	scanner := bufio.NewScanner(openStdinOrFile())

	data := make([]int, 0, 25)
	alldata := make([]int, 0, 1000)

	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		alldata = append(alldata, v)

		if len(data) < 25 {
			data = append(data, v)
		} else {
			index = (index + 1) % 25
			if !hasTargetSum(data, v) {
				fmt.Printf("Invalid number detected: %v\n", v)
				fmt.Printf("Weakness number is: %v\n", findContiguousSet(alldata, v))
				os.Exit(0)
			} else {
				data[index] = v
			}
		}
	}
}

func findContiguousSet(haystack []int, needle int) int {
	for i, v := range haystack {
		total := v
		for j := (i + 1); total < needle; j++ {
			total += haystack[j]
			if total == needle {
				return productSmallestLargest(haystack[i:j])
			}
		}
	}
	return 0
}

func productSmallestLargest(data []int) int {
	smallest := data[0]
	largest := data[0]

	for _, v := range data {
		if v < smallest {
			smallest = v
		}
		if v > largest {
			largest = v
		}
	}

	return smallest + largest
}

func hasTargetSum(haystack []int, needle int) bool {
	for i, v := range haystack {
		target := needle - v
		for j, w := range haystack {
			if i == j {
				continue
			}
			if w == target {
				return true
			}
		}
	}
	return false
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
