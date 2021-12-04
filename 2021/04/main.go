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

type BingoCard struct {
	hasWon  bool
	values  [][]int
	matches [][]bool
}

func main() {
	var drawStr []string
	var drawNumbers []int
	var currentCard []string
	var bingoCards []BingoCard

	scanner := bufio.NewScanner(openStdinOrFile())

	for scanner.Scan() {
		line := scanner.Text()

		if len(drawStr) == 0 {
			drawStr = strings.Split(line, ",")
			drawNumbers = make([]int, len(drawStr))
			for i, str := range drawStr {
				num, err := strconv.Atoi(str)
				if err != nil {
					log.Fatal(err)
				}
				drawNumbers[i] = num
			}
			continue
		}

		if len(line) == 0 {
			if len(currentCard) == 5 {
				bingoCards = append(bingoCards, parseCard(currentCard))
				currentCard = make([]string, 0)
			}
			continue
		}

		currentCard = append(currentCard, line)
	}

	for _, num := range drawNumbers {
		for i, card := range bingoCards {
			if !card.hasWon {
				card.callNumber(num)
				if card.isWinner() {
					bingoCards[i].hasWon = true
					fmt.Printf("BINGO! The Card Score was %v\n", card.score(num))
				}
			}
		}
	}

}

func parseCard(cardStr []string) BingoCard {

	card := BingoCard{
		false,
		make([][]int, 5),
		make([][]bool, 5),
	}

	for i, line := range cardStr {
		card.values[i] = make([]int, 5)
		card.matches[i] = make([]bool, 5)

		numbers := strings.Fields(line)

		for j, num := range numbers {
			value, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			card.values[i][j] = value
		}
	}

	return card
}

func (card *BingoCard) callNumber(number int) {
	for r, row := range card.values {
		for c, col := range row {
			if col == number {
				card.matches[r][c] = true
			}
		}
	}
}

func (card *BingoCard) isWinner() bool {
	colWin := []bool{true, true, true, true, true}

	for _, row := range card.matches {
		rowWin := true
		for c, col := range row {
			if !col {
				rowWin = false
				colWin[c] = false
			}
		}
		if rowWin {
			return true
		}
	}
	for _, col := range colWin {
		if col {
			return true
		}
	}

	return false
}

func (card *BingoCard) score(lastCall int) int {
	sum := 0
	for r, row := range card.matches {
		for c, col := range row {
			if !col {
				sum += card.values[r][c]
			}
		}
	}

	return sum * lastCall
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
