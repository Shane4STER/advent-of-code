package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type targetSum struct {
	a int
	b int
}

func main() {
	numbers := []int{}
	targets := make(map[int]targetSum)

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if n, err := strconv.Atoi(scanner.Text()); err == nil {
			var target = 2020 - n
			match, exists := targets[target]
			if exists {
				fmt.Printf("Value A: %v, Value B: %v, Value C: %v Product: %v", match.a, match.b, n, (match.a * match.b * n))
				os.Exit(0)
			} else {
				for _, x := range numbers {
					var sum = x + n
					targets[sum] = targetSum{x, n}
				}
				numbers = append(numbers, n)
			}
		} else {
			log.Println(err)
		}
	}
	os.Exit(1)

}
