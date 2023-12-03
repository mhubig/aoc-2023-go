package main

import (
	_ "embed"
	"testing"
)

//go:embed data-test-1.txt
var testData1 string

//go:embed data-test-2.txt
var testData2 string

func TestFindCalibrationValue_WithTestData_1(t *testing.T) {
	expected := 142
	regexPattern := "[1-9]"
	result := computeCalibrationValueSum(testData1, regexPattern)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFindCalibrationValue_WithTestData_2(t *testing.T) {
	expected := 281
	regexPattern := "[0-9]|one|two|three|four|five|six|seven|eight|nine"
	result := computeCalibrationValueSum(testData2, regexPattern)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
