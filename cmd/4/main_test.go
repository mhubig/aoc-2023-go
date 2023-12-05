package main

import (
	"testing"
)

func TestUnmarshalText_HasCorrectCardNumber(t *testing.T) {
	given := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	expected := &ScratchCard{
		CardNumber:     "1",
		GivenNumbers:   []int{41, 48, 83, 86, 17},
		WinningNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53},
	}

	card := &ScratchCard{}
	err := card.UnmarshalText([]byte(given))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if card.CardNumber != expected.CardNumber {
		t.Errorf("Expected %s, got %s", expected.CardNumber, card.CardNumber)
	}
}

func TestUnmarshalText_HasCorrectWinningNumbers(t *testing.T) {
	given := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	expected := &ScratchCard{
		CardNumber:     "1",
		WinningNumbers: []int{41, 48, 83, 86, 17},
		GivenNumbers:   []int{83, 86, 6, 31, 17, 9, 48, 53},
	}

	card := &ScratchCard{}
	err := card.UnmarshalText([]byte(given))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(card.WinningNumbers) != len(expected.WinningNumbers) {
		t.Errorf("Expected %d, got %d", len(expected.WinningNumbers), len(card.WinningNumbers))
	}
}

func TestUnmarshalText_HasCorrectGivenNumbers(t *testing.T) {
	given := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	expected := &ScratchCard{
		CardNumber:     "1",
		WinningNumbers: []int{41, 48, 83, 86, 17},
		GivenNumbers:   []int{83, 86, 6, 31, 17, 9, 48, 53},
	}

	card := &ScratchCard{}
	err := card.UnmarshalText([]byte(given))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(card.GivenNumbers) != len(expected.GivenNumbers) {
		t.Errorf("Expected %d, got %d", len(expected.GivenNumbers), len(card.GivenNumbers))
	}
}

func TestGetLuckyNumbers_EmptyScratchcard_HasNoLuckyNumbers(t *testing.T) {
	given := &ScratchCard{}

	if len(given.GetLuckyNumbers()) != 0 {
		t.Errorf("Expected %d, got %d", 0, len(given.GetLuckyNumbers()))
	}
}

func TestGetLuckyNumbers_ScratchcardWithNoLuckyNumbers_HasNoLuckyNumbers(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3},
		GivenNumbers:   []int{4, 5, 6},
	}

	if len(given.GetLuckyNumbers()) != 0 {
		t.Errorf("Expected %d, got %d", 0, len(given.GetLuckyNumbers()))
	}
}

func TestGetLuckyNumbers_ScratchcardWithOneLuckyNumber_HasOneLuckyNumber(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3},
		GivenNumbers:   []int{4, 5, 3},
	}

	if len(given.GetLuckyNumbers()) != 1 {
		t.Errorf("Expected %d, got %d", 1, len(given.GetLuckyNumbers()))
	}

	if given.GetLuckyNumbers()[0] != 3 {
		t.Errorf("Expected %d, got %d", 3, given.GetLuckyNumbers()[0])
	}
}

func TestGetLuckyNumbers_ScratchcardWithTwoLuckyNumbers_HasTwoLuckyNumbers(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3},
		GivenNumbers:   []int{4, 3, 2},
	}

	if len(given.GetLuckyNumbers()) != 2 {
		t.Errorf("Expected %d, got %d", 2, len(given.GetLuckyNumbers()))
	}

	if given.GetLuckyNumbers()[0] != 2 {
		t.Errorf("Expected %d, got %d", 2, given.GetLuckyNumbers()[0])
	}

	if given.GetLuckyNumbers()[1] != 3 {
		t.Errorf("Expected %d, got %d", 3, given.GetLuckyNumbers()[1])
	}
}

func TestGetPoints_EmptyScratchcard_HasNoPoints(t *testing.T) {
	given := &ScratchCard{}

	if given.GetPoints() != 0 {
		t.Errorf("Expected %d, got %d", 0, given.GetPoints())
	}
}

func TestGetPoints_ScratchcardWithNoLuckyNumbers_HasNoPoints(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3},
		GivenNumbers:   []int{4, 5, 6},
	}

	if given.GetPoints() != 0 {
		t.Errorf("Expected %d, got %d", 0, given.GetPoints())
	}
}

func TestGetPoints_ScratchcardWithOneLuckyNumber_HasOnePoint(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3},
		GivenNumbers:   []int{4, 5, 3},
	}

	if given.GetPoints() != 1 {
		t.Errorf("Expected %d, got %d", 1, given.GetPoints())
	}
}

func TestGetPoints_ScratchcardWithTwoLuckyNumbers_HasTwoPoints(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3},
		GivenNumbers:   []int{4, 3, 2},
	}

	if given.GetPoints() != 2 {
		t.Errorf("Expected %d, got %d", 2, given.GetPoints())
	}
}

func TestGetPoints_ScratchcardWithThreeLuckyNumbers_HasFourPoints(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3},
		GivenNumbers:   []int{1, 3, 2},
	}

	if given.GetPoints() != 4 {
		t.Errorf("Expected %d, got %d", 4, given.GetPoints())
	}
}

func TestGetPoints_ScratchcardWithFourLuckyNumbers_HasEightPoints(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3, 4},
		GivenNumbers:   []int{1, 3, 2, 4},
	}

	if given.GetPoints() != 8 {
		t.Errorf("Expected %d, got %d", 8, given.GetPoints())
	}
}
