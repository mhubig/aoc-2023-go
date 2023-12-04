package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Symbol struct {
	Symbol   string
	Position Coordinate
}

type SymbolList []Symbol

func findSymbols(lineNumber int, line string) (found SymbolList) {
	re := regexp.MustCompile(`[^a-zA-Z0-9.]`)
	index := re.FindAllStringIndex(line, -1)

	for i := range index {
		found = append(found, Symbol{
			Symbol: line[index[i][0]:index[i][1]],
			Position: Coordinate{
				X: index[i][0],
				Y: lineNumber,
			},
		})
	}

	return found
}

type Number struct {
	Number   int
	Length   int
	Position Coordinate
}

type NumberList []Number

func findNumbers(lineNumber int, line string) (found NumberList) {
	re := regexp.MustCompile(`\d+`)
	index := re.FindAllStringIndex(line, -1)

	for i := range index {
		length := len(line[index[i][0]:index[i][1]])
		number, err := strconv.Atoi(line[index[i][0]:index[i][1]])
		if err != nil {
			fmt.Println(err)
			return found
		}

		found = append(found, Number{
			Number: number,
			Length: length,
			Position: Coordinate{
				X: index[i][0],
				Y: lineNumber,
			},
		})
	}

	return found
}

func getCoordinatesOfSymbols(data string) (symbols SymbolList) {
	lines := strings.Split(data, "\n")

	for i := range lines {
		symbols = append(symbols, findSymbols(i, lines[i])...)
	}

	return symbols
}

func getCoordinatesOfNumbers(data string) (numbers NumberList) {
	lines := strings.Split(data, "\n")

	for i := range lines {
		numbers = append(numbers, findNumbers(i, lines[i])...)
	}

	return numbers
}

func filterNumbers(numbers NumberList, symbol Symbol) (filtered NumberList) {
	x := symbol.Position.X
	y := symbol.Position.Y

	for n := range numbers {
		if numbers[n].Position.X >= x-numbers[n].Length &&
			numbers[n].Position.X <= x+1 &&
			numbers[n].Position.Y >= y-1 &&
			numbers[n].Position.Y <= y+1 {
			filtered = append(filtered, numbers[n])
		}
	}

	return filtered
}

//go:embed data.txt
var data string

func main() {
	symbols := getCoordinatesOfSymbols(data)
	numbers := getCoordinatesOfNumbers(data)

	var partNumbers NumberList
	for s := range symbols {
		partNumbers = append(partNumbers, filterNumbers(numbers, symbols[s])...)
	}

	fmt.Println(partNumbers)

	var sumOfPartNumbers int
	for n := range partNumbers {
		sumOfPartNumbers += partNumbers[n].Number
	}

	fmt.Println(sumOfPartNumbers)
}
