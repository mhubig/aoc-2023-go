package main

import (
	_ "embed"
	"testing"
)

func TestFindAllMatchesOfNumbers_ShouldFindTwoMatches(t *testing.T) {
	matches := findAllMatchesOfNumber(6, "6", "66")

	if len(matches) != 2 {
		t.Errorf("Expected %d, got %d", 2, len(matches))
	}
}

func TestFindAllMatchesOfNumbers_ShouldFindNoMatches(t *testing.T) {
	matches := findAllMatchesOfNumber(6, "6", "77")

	if len(matches) != 0 {
		t.Errorf("Expected %d, got %d", 0, len(matches))
	}
}

func TestGetSortedNumbers_GivenNonOverlappingNumbers_ShouldFindTwoMatches(t *testing.T) {
	line := "oneeight"
	matches := getSortedNumbers(line)

	if len(matches) != 2 {
		t.Errorf("Expected %d, got %d", 2, len(matches))
	}

	if matches[0].Number != 1 {
		t.Errorf("Expected %d, got %d", 1, matches[0].Number)
	}

	if matches[0].Index != 0 {
		t.Errorf("Expected %d, got %d", 0, matches[0].Index)
	}

	if matches[1].Number != 8 {
		t.Errorf("Expected %d, got %d", 8, matches[1].Number)
	}

	if matches[1].Index != 3 {
		t.Errorf("Expected %d, got %d", 3, matches[1].Index)
	}
}

func TestGetSortedNumbers_GivenOverlappingNumbers_ShouldFindTwoMatches(t *testing.T) {
	line := "eightwo"
	matches := getSortedNumbers(line)

	if len(matches) != 2 {
		t.Errorf("Expected %d, got %d", 2, len(matches))
	}

	if matches[0].Number != 8 {
		t.Errorf("Expected %d, got %d", 8, matches[0].Number)
	}

	if matches[0].Index != 0 {
		t.Errorf("Expected %d, got %d", 0, matches[0].Index)
	}

	if matches[1].Number != 2 {
		t.Errorf("Expected %d, got %d", 2, matches[1].Number)
	}

	if matches[1].Index != 4 {
		t.Errorf("Expected %d, got %d", 3, matches[1].Index)
	}
}

func TestGetSortedNumbers_GivenTwoSortedDigits_ShouldFindTwoMatches(t *testing.T) {
	line := "28"
	matches := getSortedNumbers(line)

	if len(matches) != 2 {
		t.Errorf("Expected %d, got %d", 2, len(matches))
	}

	if matches[0].Number != 2 {
		t.Errorf("Expected %d, got %d", 2, matches[0].Number)
	}

	if matches[0].Index != 0 {
		t.Errorf("Expected %d, got %d", 0, matches[0].Index)
	}

	if matches[1].Number != 8 {
		t.Errorf("Expected %d, got %d", 8, matches[1].Number)
	}

	if matches[1].Index != 1 {
		t.Errorf("Expected %d, got %d", 1, matches[1].Index)
	}
}

func TestGetSortedNumbers_GivenTwoUnsortedDigits_ShouldFindTwoMatches(t *testing.T) {
	line := "82"
	matches := getSortedNumbers(line)

	if len(matches) != 2 {
		t.Errorf("Expected %d, got %d", 2, len(matches))
	}

	if matches[0].Number != 8 {
		t.Errorf("Expected %d, got %d", 8, matches[0].Number)
	}

	if matches[0].Index != 0 {
		t.Errorf("Expected %d, got %d", 0, matches[0].Index)
	}

	if matches[1].Number != 2 {
		t.Errorf("Expected %d, got %d", 2, matches[1].Number)
	}

	if matches[1].Index != 1 {
		t.Errorf("Expected %d, got %d", 1, matches[1].Index)
	}
}

func TestGetSortedNumbers_GivenStringWithDigitsandNumbers_ShouldFindFiveMatches(t *testing.T) {
	line := "cq57sixeightwosvx"
	matches := getSortedNumbers(line)

	if len(matches) != 5 {
		t.Errorf("Expected %d, got %d", 2, len(matches))
	}

	if matches[0].Number != 5 {
		t.Errorf("Expected %d, got %d", 8, matches[0].Number)
	}

	if matches[0].Index != 2 {
		t.Errorf("Expected %d, got %d", 2, matches[0].Index)
	}

	if matches[1].Number != 7 {
		t.Errorf("Expected %d, got %d", 7, matches[1].Number)
	}

	if matches[1].Index != 3 {
		t.Errorf("Expected %d, got %d", 3, matches[1].Index)
	}

	if matches[2].Number != 6 {
		t.Errorf("Expected %d, got %d", 6, matches[1].Number)
	}

	if matches[2].Index != 4 {
		t.Errorf("Expected %d, got %d", 4, matches[1].Index)
	}

	if matches[3].Number != 8 {
		t.Errorf("Expected %d, got %d", 8, matches[1].Number)
	}

	if matches[3].Index != 7 {
		t.Errorf("Expected %d, got %d", 7, matches[1].Index)
	}

	if matches[4].Number != 2 {
		t.Errorf("Expected %d, got %d", 2, matches[1].Number)
	}

	if matches[4].Index != 11 {
		t.Errorf("Expected %d, got %d", 11, matches[1].Index)
	}
}

func TestGetSortedNumbers_GivenTwoTimesTheSameDigit_ShouldFindTwoMatches(t *testing.T) {
	line := "66"
	matches := getSortedNumbers(line)

	if len(matches) != 2 {
		t.Errorf("Expected %d, got %d", 2, len(matches))
	}

	if matches[0].Number != 6 {
		t.Errorf("Expected %d, got %d", 6, matches[0].Number)
	}

	if matches[0].Index != 0 {
		t.Errorf("Expected %d, got %d", 0, matches[1].Index)
	}

	if matches[1].Number != 6 {
		t.Errorf("Expected %d, got %d", 6, matches[0].Number)
	}

	if matches[1].Index != 1 {
		t.Errorf("Expected %d, got %d", 1, matches[1].Index)
	}
}

func TestComputeCalibrationValue_GivenTwoNumber_ConcatBoth(t *testing.T) {
	given := "onetwo"
	expected := 12

	result := computeCalibrationValue(given)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestComputeCalibrationValue_GivenThreeNumbers_ConcatFirstAndLast(t *testing.T) {
	given := "onetwothree"
	expected := 13

	result := computeCalibrationValue(given)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestComputeCalibrationValueSum(t *testing.T) {
	given := "onetwothree\nfourfivesix"
	expected := 13 + 46

	result := computeCalibrationValueSum(given)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

//go:embed data-test.txt
var testData string

func TestFindCalibrationValue_WithTestData_1(t *testing.T) {
	expected := 281
	result := computeCalibrationValueSum(testData)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
