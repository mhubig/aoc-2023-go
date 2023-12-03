package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed data.txt
var data string

func main() {
	regexPattern1 := "[1-9]"
	regexPattern2 := "[0-9]|one|two|three|four|five|six|seven|eight|nine"

	result1 := computeCalibrationValueSum(data, regexPattern1)
	fmt.Println("Part 1:", result1)

	result2 := computeCalibrationValueSum(data, regexPattern2)
	fmt.Println("Part 2:", result2)
}

func computeCalibrationValueSum(data string, pattern string) int {
	regex := regexp.MustCompile(pattern)
	lines := strings.Split(data, "\n")

	var numbers []int
	for i := 0; i < len(lines); i++ {
		numbers = append(numbers, findCalibrationValue(lines[i], regex))
	}

	var result int
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}

	return result
}

func findCalibrationValue(line string, regex *regexp.Regexp) int {
	var digits []string
	for i := 0; i < len(line); i++ {
		found := regex.FindString(line[i:])
		if found == "" {
			continue
		}
		digits = append(digits, found)
	}

	sanitizedDigits := sanitizeDigits(digits)

	first := sanitizedDigits[0]
	last := sanitizedDigits[len(sanitizedDigits)-1]

	calibrationValue, err := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(line, digits, sanitizedDigits, calibrationValue)
	return calibrationValue
}

func sanitizeDigits(digits []string) []string {
	var sanitizedDigits []string

	for i := 0; i < len(digits); i++ {
		if len(digits[i]) > 1 {
			sanitizedDigits = append(sanitizedDigits, convertToNumber(digits[i]))
		} else {
			sanitizedDigits = append(sanitizedDigits, digits[i])
		}
	}

	return sanitizedDigits
}

func convertToNumber(number string) string {
	switch number {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return "0"
	}
}
