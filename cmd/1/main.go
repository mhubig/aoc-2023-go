package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	_ "embed"
)

type Match struct {
	Number int
	Index  int
}

type MatchList []Match

func (m MatchList) Len() int {
	return len(m)
}

func (m MatchList) Less(i, j int) bool {
	return m[i].Index < m[j].Index
}

func (m MatchList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func findAllMatchesOfNumber(number int, pattern string, text string) (matches MatchList) {
	regex := regexp.MustCompile(pattern)
	index := regex.FindAllIndex([]byte(text), -1)

	for i := range index {
		matches = append(matches, Match{Number: number, Index: index[i][0]})
	}

	return matches
}

func getSortedNumbers(line string) (matches MatchList) {
	matches = append(matches, findAllMatchesOfNumber(1, "1", line)...)
	matches = append(matches, findAllMatchesOfNumber(2, "2", line)...)
	matches = append(matches, findAllMatchesOfNumber(3, "3", line)...)
	matches = append(matches, findAllMatchesOfNumber(4, "4", line)...)
	matches = append(matches, findAllMatchesOfNumber(5, "5", line)...)
	matches = append(matches, findAllMatchesOfNumber(6, "6", line)...)
	matches = append(matches, findAllMatchesOfNumber(7, "7", line)...)
	matches = append(matches, findAllMatchesOfNumber(8, "8", line)...)
	matches = append(matches, findAllMatchesOfNumber(9, "9", line)...)

	matches = append(matches, findAllMatchesOfNumber(1, "one", line)...)
	matches = append(matches, findAllMatchesOfNumber(2, "two", line)...)
	matches = append(matches, findAllMatchesOfNumber(3, "three", line)...)
	matches = append(matches, findAllMatchesOfNumber(4, "four", line)...)
	matches = append(matches, findAllMatchesOfNumber(5, "five", line)...)
	matches = append(matches, findAllMatchesOfNumber(6, "six", line)...)
	matches = append(matches, findAllMatchesOfNumber(7, "seven", line)...)
	matches = append(matches, findAllMatchesOfNumber(8, "eight", line)...)
	matches = append(matches, findAllMatchesOfNumber(9, "nine", line)...)

	sort.Sort(matches)
	return matches
}

func computeCalibrationValue(line string) int {
	matches := getSortedNumbers(line)

	fist := matches[0]
	last := matches[len(matches)-1]

	result, err := strconv.Atoi(fmt.Sprintf("%d%d", fist.Number, last.Number))
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func computeCalibrationValueSum(data string) (result int) {
	lines := strings.Split(data, "\n")

	for i := range lines {
		result += computeCalibrationValue(lines[i])
	}

	return result
}

//go:embed data.txt
var data string

func main() {
	fmt.Println(computeCalibrationValueSum(data))
}
