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
	result := computeCalibrationValueSum(data)
	fmt.Println(result)
}

func computeCalibrationValueSum(data string) int {
	lines := strings.Split(data, "\n")

	var numbers []int
	for i := 0; i < len(lines); i++ {
		numbers = append(numbers, findCalibrationValue(lines[i]))
	}

	var result int
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}

	return result
}

func findCalibrationValue(text string) int {
	re := regexp.MustCompile("[0-9]")
	digits := re.FindAllString(text, -1)

	first := digits[0]
	last := digits[len(digits)-1]

	calibrationValue, err := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
	if err != nil {
		log.Fatal(err)
	}

	return calibrationValue
}
