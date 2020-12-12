package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var adapterJoltages []int
	var delta1, delta3 int

	scanner := bufio.NewScanner(openStdinOrFile())

	adapterJoltages = append(adapterJoltages, 0)

	for scanner.Scan() {
		joltage, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		adapterJoltages = append(adapterJoltages, joltage)
	}

	sort.Ints(adapterJoltages)

	adapterJoltages = append(adapterJoltages, (adapterJoltages[len(adapterJoltages)-1] + 3))

	for i, j := range adapterJoltages {
		if i+1 >= len(adapterJoltages) {
			continue
		}
		if adapterJoltages[i+1]-j == 1 {
			delta1++
		} else if adapterJoltages[i+1]-j == 3 {
			delta3++
		}
	}

	fmt.Printf("delta 1: %v\ndelta 3: %v\nproduct: %v\n", delta1, delta3, (delta1 * delta3))

	fmt.Printf("Combos: %v\n", validCombo(adapterJoltages))
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

func lookahead(adapters []int) []int {
	current := adapters[0]
	for i, v := range adapters {
		if i == 0 {
			continue
		}
		if v-current > 3 {
			return adapters[1:i]
		}
	}
	return adapters[1:1]
}

func validCombo(adapters []int) int {
	var r int
	validCombos := make([]int, len(adapters), len(adapters))
	validCombos[0] = 1

	for i, v := range adapters {
		if i == 0 {
			continue
		}
		for j := 1; j <= 3; j++ {
			if i-j >= 0 {
				if v-adapters[i-j] <= 3 {
					validCombos[i] += validCombos[i-j]
				}
			}
		}
		r = validCombos[i]
	}
	return r
}

func sum(a []int) int {
	sum := 1
	for _, v := range a {
		if v != 1 {
			sum = sum * v
		}
	}
	return sum
}
