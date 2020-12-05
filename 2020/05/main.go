package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type seat struct {
	row    int64
	column int64
}

func main() {
	var largestID int64
	var smallestID int64

	scanner := bufio.NewScanner(openStdinOrFile())

	allPasses := make(map[int64]seat)
	smallestID = math.MaxInt64

	for scanner.Scan() {
		seat := findSeat(scanner.Text())
		allPasses[seat.getID()] = seat
		if seat.getID() > largestID {
			largestID = seat.getID()
		}
		if seat.getID() < smallestID {
			smallestID = seat.getID()
		}
	}

	firstSeat := allPasses[smallestID]
	lastSeat := allPasses[largestID]

	for r := firstSeat.row; r <= lastSeat.row; r++ {
		for c := 0; c < 8; c++ {
			testSeat := seat{int64(r), int64(c)}
			testSeatID := testSeat.getID()
			_, exists := allPasses[testSeatID]
			_, hasIDBefore := allPasses[testSeatID-1]
			_, hasIDAfter := allPasses[testSeatID+1]
			if !exists && hasIDBefore && hasIDAfter {
				fmt.Printf("Your seat is (%v): %v", testSeatID, testSeat)
			}
		}
	}

	fmt.Printf("The largest seat ID is %v\n", largestID)
}

func findSeat(seatString string) seat {
	return seat{
		getRow(seatString),
		getColumn(seatString),
	}
}

func getRow(seatString string) int64 {
	rowString := seatString[:7]

	rowStringBin := strings.Map(func(r rune) rune {
		if r == 'B' {
			return '1'
		} else {
			return '0'
		}
	}, rowString)

	rowNumber, err := strconv.ParseInt(rowStringBin, 2, 0)
	if err != nil {
		log.Fatal(err)
	}

	return rowNumber
}

func getColumn(seatString string) int64 {
	columnString := seatString[7:]

	columnStringBin := strings.Map(func(r rune) rune {
		if r == 'R' {
			return '1'
		} else {
			return '0'
		}
	}, columnString)

	columnNumber, err := strconv.ParseInt(columnStringBin, 2, 0)
	if err != nil {
		log.Fatal(err)
	}

	return columnNumber
}

func (s seat) getID() int64 {
	return (s.row * 8) + s.column
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
