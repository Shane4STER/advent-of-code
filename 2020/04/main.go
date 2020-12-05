package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type height struct {
	value int
	unit  string
}

type idCard struct {
	byr int
	iyr int
	eyr int
	hgt height
	ecl string
	hcl string
	pid string
	cid int
}

type idCardCheck struct {
	validPassport      bool
	validNorthPoleCard bool
	byr                bool
	iyr                bool
	eyr                bool
	hgt                bool
	ecl                bool
	hcl                bool
	pid                bool
	cid                bool
}

func main() {
	var validPassports int
	currentIDData := make([]string, 0, 8)
	scanner := bufio.NewScanner(openStdinOrFile())

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			currentIDData = append(currentIDData, line)
			continue
		} else if len(currentIDData) > 0 {
			if testChunk(currentIDData, true) {
				validPassports++
			}
			currentIDData = currentIDData[:0]
		}
	}
	if testChunk(currentIDData, true) {
		validPassports++
	}

	fmt.Printf("Found %v valid passports\n", validPassports)
}

func testChunk(chunk []string, allowIDCard bool) bool {
	idData := strings.Join(chunk, " ")
	idCard := parseIDCard(idData)

	result := isValid(idCard)

	return result.validPassport || allowIDCard && result.validNorthPoleCard
}

func isValid(pp idCard) idCardCheck {
	validEyeColour := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	validHairColour := regexp.MustCompile("^#[0-9a-f]{6}$")
	validPassportNumber := regexp.MustCompile("^[0-9]{9}$")
	check := idCardCheck{
		false,
		false,
		pp.byr >= 1920 && pp.byr <= 2002,
		pp.iyr >= 2010 && pp.iyr <= 2020,
		pp.eyr >= 2020 && pp.eyr <= 2030,
		(pp.hgt.unit == "in" && pp.hgt.value >= 59 && pp.hgt.value <= 76) ||
			(pp.hgt.unit == "cm" && pp.hgt.value >= 150 && pp.hgt.value <= 193),
		validEyeColour.MatchString(pp.ecl),
		validHairColour.MatchString(pp.hcl),
		validPassportNumber.MatchString(pp.pid),
		pp.cid > 0,
	}
	check.validNorthPoleCard = check.byr &&
		check.iyr &&
		check.eyr &&
		check.hgt &&
		check.ecl &&
		check.hcl &&
		check.pid
	check.validPassport = check.validNorthPoleCard && check.cid

	return check
}

func parseIDCard(input string) idCard {
	var parsed idCard
	idFields := strings.Fields(input)

	for _, field := range idFields {
		var err error
		data := strings.Split(field, ":")
		fieldName := data[0]
		fieldValue := data[1]
		switch fieldName {
		case "byr":
			parsed.byr, err = strconv.Atoi(fieldValue)
			if err != nil {
				parsed.byr = 0
			}
		case "iyr":
			parsed.iyr, err = strconv.Atoi(fieldValue)
			if err != nil {
				parsed.iyr = 0
			}
		case "eyr":
			parsed.eyr, err = strconv.Atoi(fieldValue)
			if err != nil {
				parsed.eyr = 0
			}
		case "ecl":
			parsed.ecl = fieldValue
		case "hcl":
			parsed.hcl = fieldValue
		case "pid":
			parsed.pid = fieldValue
		case "cid":
			parsed.cid, err = strconv.Atoi(fieldValue)
			if err != nil {
				parsed.cid = 0
			}
		case "hgt":
			heightValue, err := strconv.Atoi(fieldValue[:len(fieldValue)-2])
			if err != nil {
				parsed.cid = 0
			}
			parsed.hgt = height{
				heightValue,
				fieldValue[len(fieldValue)-2:],
			}
		default:
			fmt.Printf("Found unknown field %v\n", fieldName)
		}
	}

	return parsed
}

func (id idCardCheck) toString() string {
	return fmt.Sprintf(`
	isValidPassport: %v,
	isValidNorthPoleCard: %v,
	hasValidBirthYear: %v,
	hasValidIssueYear: %v,
	hasValidExpiryYear: %v,
	hasValidHeight: %v,
	hasValidEyeColour: %v,
	hasValidHairColour: %v,
	hasValidPassportId: %v,
	hasValidCountryId: %v
	`, id.validPassport, id.validNorthPoleCard, id.byr, id.iyr, id.eyr, id.hgt, id.ecl, id.hcl, id.pid, id.cid)
}

func (id idCard) toString() string {
	return fmt.Sprintf(`
	Birth Year: %v,
	Issue Year: %v,
	Expiry Year: %v,
	Height: %v,
	Eye Colour: %v,
	Hair Colour: %v,
	Passport ID: %v,
	Country Id: %v`, id.byr, id.iyr, id.eyr, id.hgt, id.ecl, id.hcl, id.pid, id.cid)
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
