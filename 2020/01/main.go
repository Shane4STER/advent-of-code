package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := make(map[int]bool)

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if n, err := strconv.Atoi(scanner.Text()); err == nil {
			var target = 2020 - n
			if numbers[target] {
				fmt.Printf("Value A: %v, Value B: %v, Product: %v", target, n, (target * n))
				os.Exit(0)
			} else {
				numbers[n] = true
			}
		}
	}
	os.Exit(1)

}
