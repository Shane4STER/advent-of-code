package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/deanveloper/modmath/v1/bigmod"
)

func main() {
	scanner := bufio.NewScanner(openStdinOrFile())
	var minTime, bestRoute, activeRoutes int

	scanner.Scan()
	target, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	schedule := parseBuses(scanner.Text())

	for _, route := range schedule {
		if route < 0 {
			continue
		}
		activeRoutes++
		waitingTime := (((target / route) + 1) * route) - target
		if waitingTime < minTime || minTime == 0 {
			minTime = waitingTime
			bestRoute = route
		}
	}
	part1 := minTime * bestRoute
	fmt.Printf("Part 1: %v\n", part1)

	var pIndex int
	problem := make([]bigmod.CrtEntry, activeRoutes)

	for i, route := range schedule {
		if route < 0 {
			continue
		}

		problem[pIndex] = bigmod.CrtEntry{
			big.NewInt(int64(((-i) % route) + route)),
			big.NewInt(int64(route)),
		}
		pIndex++
	}
	solution := bigmod.SolveCrtMany(problem)
	fmt.Printf("Part 2: %v\n", solution)
}

func parseBuses(schedule string) []int {
	scheduleArr := strings.Split(schedule, ",")
	var r []int
	for _, s := range scheduleArr {
		if s == "x" {
			r = append(r, -1)
		} else {
			current, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			r = append(r, current)
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
