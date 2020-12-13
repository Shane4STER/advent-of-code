package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(openStdinOrFile())
	var minTime, bestRoute int

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
		waitingTime := (((target / route) + 1) * route) - target
		if waitingTime < minTime || minTime == 0 {
			minTime = waitingTime
			bestRoute = route
		}
	}
	part1 := minTime * bestRoute
	fmt.Println(part1)

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
