package main

import (
	"testing"
)

func TestFindSymbols_GivenOneSymbol_ReturnOneSymbol(t *testing.T) {
	given := "783.*.234"
	expected := SymbolList{
		{Symbol: "*", Position: Coordinate{4, 0}},
	}

	found := findSymbols(0, given)

	if len(found) != len(expected) {
		t.Errorf("Expected %d, got %d", len(expected), len(found))
	}

	if found[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected[0], found[0])
	}
}

func TestFindSymbols_GivenTwoSymbols_ReturnTwoSymbols(t *testing.T) {
	given := "...$.*...."
	expected := SymbolList{
		{Symbol: "$", Position: Coordinate{3, 0}},
		{Symbol: "*", Position: Coordinate{5, 0}},
	}

	found := findSymbols(0, given)

	if len(found) != len(expected) {
		t.Errorf("Expected %d, got %d", len(expected), len(found))
	}

	if found[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected[0], found[0])
	}

	if found[1] != expected[1] {
		t.Errorf("Expected %v, got %v", expected[1], found[1])
	}
}

func TestFindNumbers_GivenOneNumber_ReturnOneNumber(t *testing.T) {
	given := "783"
	expected := NumberList{
		{Number: 783, Length: 3, Position: Coordinate{0, 0}},
	}

	found := findNumbers(0, given)

	if len(found) != len(expected) {
		t.Errorf("Expected %d, got %d", len(expected), len(found))
	}

	if found[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected[0], found[0])
	}
}

func TestFindNumbers_GivenOneNumberWithLengthTwo_ReturnOneNumber(t *testing.T) {
	given := "78"
	expected := NumberList{
		{Number: 78, Length: 2, Position: Coordinate{0, 0}},
	}

	found := findNumbers(0, given)

	if len(found) != len(expected) {
		t.Errorf("Expected %d, got %d", len(expected), len(found))
	}

	if found[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected[0], found[0])
	}
}

func TestFindNumbers_GivenTwoDifferentNumbers_ReturnTwoNumbers(t *testing.T) {
	given := ".111..222."
	expected := NumberList{
		{Number: 111, Length: 3, Position: Coordinate{1, 0}},
		{Number: 222, Length: 3, Position: Coordinate{6, 0}},
	}

	found := findNumbers(0, given)

	if len(found) != len(expected) {
		t.Errorf("Expected %d, got %d", len(expected), len(found))
	}

	if found[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected[0], found[0])
	}

	if found[1] != expected[1] {
		t.Errorf("Expected %v, got %v", expected[1], found[1])
	}
}

func TestFindNumbers_GivenTheSameTwoNumbers_ReturnTwoNumbers(t *testing.T) {
	given := ".111..111."
	expected := NumberList{
		{Number: 111, Length: 3, Position: Coordinate{1, 0}},
		{Number: 111, Length: 3, Position: Coordinate{6, 0}},
	}

	found := findNumbers(0, given)

	if len(found) != len(expected) {
		t.Errorf("Expected %d, got %d", len(expected), len(found))
	}

	if found[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected[0], found[0])
	}

	if found[1] != expected[1] {
		t.Errorf("Expected %v, got %v", expected[1], found[1])
	}
}

func TestGetCoordinatesOfSymbols_GivenTwoLinesWithOneSymbolEach_ReturnTwoSymbols(t *testing.T) {
	given := "*\n*"
	expected := SymbolList{
		{Symbol: "*", Position: Coordinate{0, 0}},
		{Symbol: "*", Position: Coordinate{0, 1}},
	}

	found := getCoordinatesOfSymbols(given)

	if len(found) != len(expected) {
		t.Errorf("Expected %d, got %d", len(expected), len(found))
	}

	if found[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected[0], found[0])
	}

	if found[1] != expected[1] {
		t.Errorf("Expected %v, got %v", expected[1], found[1])
	}
}

func TestGetCoordinatesOfSymbols_GivenTwoLinesWithSymbolsAndNumbers_ReturnTwoSymbols(t *testing.T) {
	given := "111.*.111\n222.*.222"
	expected := SymbolList{
		{Symbol: "*", Position: Coordinate{4, 0}},
		{Symbol: "*", Position: Coordinate{4, 1}},
	}

	found := getCoordinatesOfSymbols(given)

	if len(found) != len(expected) {
		t.Errorf("Expected %d, got %d", len(expected), len(found))
	}

	if found[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected[0], found[0])
	}

	if found[1] != expected[1] {
		t.Errorf("Expected %v, got %v", expected[1], found[1])
	}
}

func TestGetCoordinatesOfNumbers_GivenTwoLinesWithOneNumberEach_ReturnTwoNumbers(t *testing.T) {
	given := "111\n222"
	expected := NumberList{
		{Number: 111, Length: 3, Position: Coordinate{0, 0}},
		{Number: 222, Length: 3, Position: Coordinate{0, 1}},
	}

	found := getCoordinatesOfNumbers(given)

	if len(found) != len(expected) {
		t.Errorf("Expected %d, got %d", len(expected), len(found))
	}

	if found[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected[0], found[0])
	}

	if found[1] != expected[1] {
		t.Errorf("Expected %v, got %v", expected[1], found[1])
	}
}

func TestGetCoordinatesOfNumberss_GivenTwoLinesWithSymbolsAndNumbers_ReturnFourNumbers(t *testing.T) {
	given := "111.*.111\n222.*.222"
	expected := NumberList{
		{Number: 111, Length: 3, Position: Coordinate{0, 0}},
		{Number: 111, Length: 3, Position: Coordinate{6, 0}},
		{Number: 222, Length: 3, Position: Coordinate{0, 1}},
		{Number: 222, Length: 3, Position: Coordinate{6, 1}},
	}

	found := getCoordinatesOfNumbers(given)

	if len(found) != len(expected) {
		t.Errorf("Expected %d, got %d", len(expected), len(found))
	}

	if found[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected[0], found[0])
	}

	if found[1] != expected[1] {
		t.Errorf("Expected %v, got %v", expected[1], found[1])
	}

	if found[2] != expected[2] {
		t.Errorf("Expected %v, got %v", expected[2], found[2])
	}

	if found[3] != expected[3] {
		t.Errorf("Expected %v, got %v", expected[3], found[3])
	}
}

func TestFilterNumbers_GivenOneSymbolAndOneAdjacentNumber_ReturnTheNumber(t *testing.T) {
	givenList := []string{
		"111....\n...*...\n.......",
		".111...\n...*...\n.......",
		"..111..\n...*...\n.......",
		"...111.\n...*...\n.......",
		"....111\n...*...\n.......",
		".......\n111*...\n.......",
		".......\n...*111\n.......",
		".......\n...*...\n111....",
		".......\n...*...\n.111...",
		".......\n...*...\n..111..",
		".......\n...*...\n...111.",
		".......\n...*...\n....111",
	}

	for i := range givenList {
		data := givenList[i]
		symbols := getCoordinatesOfSymbols(data)
		numbers := getCoordinatesOfNumbers(data)
		filtered := filterNumbers(numbers, symbols[0])

		if len(symbols) != 1 {
			t.Errorf("Expected %d, got %d", 1, len(symbols))
		}

		if symbols[0].Symbol != "*" {
			t.Errorf("Expected %v, got %v", "*", symbols[0].Symbol)
		}

		if symbols[0].Position.X != 3 {
			t.Errorf("Expected %d, got %d", 3, symbols[0].Position.X)
		}

		if symbols[0].Position.Y != 1 {
			t.Errorf("Expected %d, got %d", 1, symbols[0].Position.Y)
		}

		if len(numbers) != 1 {
			t.Errorf("Expected %d, got %d", 1, len(numbers))
		}

		if len(filtered) != 1 {
			t.Errorf("Expected %d, got %d", 1, len(filtered))
		}

		if filtered[0] != numbers[0] {
			t.Errorf("Expected %v, got %v", numbers[0], filtered[0])
		}

		if filtered[0].Number != 111 {
			t.Errorf("Expected %d, got %d", 111, filtered[0].Number)
		}
	}
}

func TestFilterNumbers_GivenOneSymbolAndOneNotAdjacentNumber_ReturnNothing(t *testing.T) {
	givenList := "11.....\n...*...\n......."
	symbols := getCoordinatesOfSymbols(givenList)
	numbers := getCoordinatesOfNumbers(givenList)
	filtered := filterNumbers(numbers, symbols[0])

	if len(symbols) != 1 {
		t.Errorf("Expected %d, got %d", 1, len(symbols))
	}

	if symbols[0].Symbol != "*" {
		t.Errorf("Expected %v, got %v", "*", symbols[0].Symbol)
	}

	if symbols[0].Position.X != 3 {
		t.Errorf("Expected %d, got %d", 3, symbols[0].Position.X)
	}

	if symbols[0].Position.Y != 1 {
		t.Errorf("Expected %d, got %d", 1, symbols[0].Position.Y)
	}

	if len(numbers) != 1 {
		t.Errorf("Expected %d, got %d", 1, len(numbers))
	}

	if len(filtered) != 0 {
		t.Errorf("Expected %d, got %d", 0, len(filtered))
	}
}
