package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type height struct {
	value int
	unit  string
}

type passport struct {
	byr int
	iyr int
	eyr int
	hgt height
	ecl string
	hcl string
	pid int
	cid int
}

func main() {
	var validPassports int
	currentPassportData := make([]string, 0, 8)
	scanner := bufio.NewScanner(openStdinOrFile())

	for scanner.Scan() {
		currentLine := scanner.Text()
		if len(currentLine) > 0 {
			currentPassportData = append(currentPassportData, scanner.Text())
		} else {
			currentPassport := parsePassport(strings.Join(currentPassportData, " "))
			if isValid(currentPassport, true) {
				validPassports++
			}
			currentPassportData = make([]string, 0, 8)
		}
	}

	fmt.Printf("Found %v valid passports\n", validPassports)
}

func isValid(pp passport, allowNPC bool) bool {
	return pp.byr > 0 &&
		pp.iyr > 0 &&
		pp.eyr > 0 &&
		pp.hgt.value > 0 &&
		len(pp.ecl) > 0 &&
		len(pp.hcl) > 0 &&
		pp.pid > 0 &&
		(pp.cid > 0 || allowNPC)
}

func parsePassport(input string) passport {
	var parsed passport
	ppFields := strings.Fields(input)

	for _, field := range ppFields {
		var err error
		data := strings.Split(field, ":")
		fieldName := data[0]
		fieldValue := data[1]
		switch fieldName {
		case "byr":
			parsed.byr, err = strconv.Atoi(fieldValue)
			if err != nil {
				fmt.Println("Invalid passport, returning empty.")
				return passport{}
			}
		case "iyr":
			parsed.iyr, err = strconv.Atoi(fieldValue)
			if err != nil {
				fmt.Println("Invalid passport, returning empty.")
				return passport{}
			}
		case "eyr":
			parsed.eyr, err = strconv.Atoi(fieldValue)
			if err != nil {
				fmt.Println("Invalid passport, returning empty.")
				return passport{}
			}
		case "ecl":
			parsed.ecl = fieldValue
		case "hcl":
			parsed.hcl = fieldValue
		case "pid":
			parsed.pid, err = strconv.Atoi(fieldValue)
			if err != nil {
				fmt.Println("Invalid passport, returning empty.")
				return passport{}
			}
		case "cid":
			parsed.cid, err = strconv.Atoi(fieldValue)
			if err != nil {
				fmt.Println("Invalid passport, returning empty.")
				return passport{}
			}
		case "hgt":
			parsed.hgt.unit = fieldValue[len(fieldValue)-2:]
			parsed.hgt.value, err = strconv.Atoi(fieldValue[:len(fieldValue)-2])
			if err != nil {
				fmt.Println("Invalid passport, returning empty.")
				return passport{}
			}
		default:
			fmt.Printf("Found unknown field %v\n", fieldName)
		}
	}

	return parsed
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
