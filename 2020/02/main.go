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

	var validPasswords, totalPasswords int

	for scanner.Scan() {
		totalPasswords++
		if isValidPassword(scanner.Text()) {
			validPasswords++
		}
	}

	fmt.Printf("There wew %v valid password of %v total passwords in the stream", validPasswords, totalPasswords)

}

func isValidPassword(input string) bool {
	minCount, err := strconv.Atoi(input[0:strings.Index(input, "-")])
	if err != nil {
		log.Println(err)
		return false
	}
	maxCount, err := strconv.Atoi(input[strings.Index(input, "-")+1 : strings.Index(input, " ")])
	if err != nil {
		log.Println(err)
		return false
	}
	character := input[strings.Index(input, " ")+1 : strings.Index(input, ":")]

	charCount := strings.Count(input[strings.Index(input, ":")+1:], character)

	return charCount >= minCount && charCount <= maxCount
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
