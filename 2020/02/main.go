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

	var validPasswords, validPasswordsTwo, totalPasswords int

	for scanner.Scan() {
		password := scanner.Text()
		totalPasswords++
		if isValidPassword(password) {
			validPasswords++
		}
		if isValidPasswordTwo(password) {
			validPasswordsTwo++
		}
	}

	fmt.Printf("There wew %v valid password of %v total passwords in the stream [Policy 1]", validPasswords, totalPasswords)
	fmt.Printf("There wew %v valid password of %v total passwords in the stream [Policy 2]", validPasswordsTwo, totalPasswords)

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

func isValidPasswordTwo(input string) bool {
	firstIndex, err := strconv.Atoi(input[0:strings.Index(input, "-")])
	if err != nil {
		log.Println(err)
		return false
	}
	secondIndex, err := strconv.Atoi(input[strings.Index(input, "-")+1 : strings.Index(input, " ")])
	if err != nil {
		log.Println(err)
		return false
	}
	character := input[strings.Index(input, " ")+1]
	password := input[strings.Index(input, ":")+1:]

	firstChar := password[firstIndex]
	secondChar := password[secondIndex]

	return firstChar != secondChar && (firstChar == character || secondChar == character)

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
