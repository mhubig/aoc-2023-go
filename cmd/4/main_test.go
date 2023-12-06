package main

import (
	"testing"
)

func TestUnmarshalText_HasCorrectCardNumber(t *testing.T) {
	given := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	expected := &ScratchCard{
		CardNumber:     1,
		GivenNumbers:   []int{41, 48, 83, 86, 17},
		WinningNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53},
	}

	card := &ScratchCard{}
	err := card.UnmarshalText([]byte(given))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if card.CardNumber != expected.CardNumber {
		t.Errorf("Expected %d, got %d", expected.CardNumber, card.CardNumber)
	}
}

func TestUnmarshalText_HasCorrectWinningNumbers(t *testing.T) {
	given := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	expected := &ScratchCard{
		CardNumber:     1,
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
		CardNumber:     1,
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

func TestSetWins_EmptyScratchcard_HasNoWins(t *testing.T) {
	given := &ScratchCard{}

	setWins(given)

	if given.Wins != 0 {
		t.Errorf("Expected %d, got %d", 0, given.Wins)
	}
}

func TestSetWins_ScratchcardWithNoIntersectingNumbers_HasNoWins(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3},
		GivenNumbers:   []int{4, 5, 6},
	}

	setWins(given)

	if given.Wins != 0 {
		t.Errorf("Expected %d, got %d", 0, given.Wins)
	}
}

func TestSetWins_ScratchcardWithOneIntersectingNumber_HasOneWin(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3},
		GivenNumbers:   []int{4, 5, 3},
	}

	setWins(given)

	if given.Wins != 1 {
		t.Errorf("Expected %d, got %d", 1, given.Wins)
	}
}

func TestSetWins_ScratchcardWithTwoIntersectingNumbers_HasTwoWins(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3},
		GivenNumbers:   []int{4, 3, 2},
	}

	setWins(given)

	if given.Wins != 2 {
		t.Errorf("Expected %d, got %d", 2, given.Wins)
	}
}

func TestGetPoints_EmptyScratchcard_HasNoPoints(t *testing.T) {
	given := &ScratchCard{}

	setWins(given)
	setPoints(given)

	if given.Points != 0 {
		t.Errorf("Expected %d, got %d", 0, given.Points)
	}
}

func TestGetPoints_ScratchcardWithIntersectingNumbers_HasNoPoints(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3},
		GivenNumbers:   []int{4, 5, 6},
	}

	setWins(given)
	setPoints(given)

	if given.Points != 0 {
		t.Errorf("Expected %d, got %d", 0, given.Points)
	}
}

func TestGetPoints_ScratchcardWithOneIntersectingNumber_HasOnePoint(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3},
		GivenNumbers:   []int{4, 5, 3},
	}

	setWins(given)
	setPoints(given)

	if given.Points != 1 {
		t.Errorf("Expected %d, got %d", 1, given.Points)
	}
}

func TestGetPoints_ScratchcardWithTwoIntersectingNumbers_HasTwoPoints(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3},
		GivenNumbers:   []int{4, 3, 2},
	}

	setWins(given)
	setPoints(given)

	if given.Points != 2 {
		t.Errorf("Expected %d, got %d", 2, given.Points)
	}
}

func TestGetPoints_ScratchcardWithThreeIntersectingNumbers_HasFourPoints(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3},
		GivenNumbers:   []int{1, 3, 2},
	}

	setWins(given)
	setPoints(given)

	if given.Points != 4 {
		t.Errorf("Expected %d, got %d", 4, given.Points)
	}
}

func TestGetPoints_ScratchcardWithFourIntersectingNumbers_HasEightPoints(t *testing.T) {
	given := &ScratchCard{
		WinningNumbers: []int{1, 2, 3, 4},
		GivenNumbers:   []int{1, 3, 2, 4},
	}

	setWins(given)
	setPoints(given)

	if given.Points != 8 {
		t.Errorf("Expected %d, got %d", 8, given.Points)
	}
}

func TestWinScatchCards_GivenTwoCardsWithOneWin_ResultIsThree(t *testing.T) {
	given := []ScratchCard{
		{
			CardNumber:     1,
			WinningNumbers: []int{1, 2, 3},
			GivenNumbers:   []int{4, 5, 3},
			Wins:           1,
			Points:         1,
		},
		{
			CardNumber:     2,
			WinningNumbers: []int{1, 2, 3},
			GivenNumbers:   []int{4, 5, 6},
			Wins:           0,
			Points:         0,
		},
	}
	expected := 3

	result := winScatchCards(given)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
