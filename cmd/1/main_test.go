package main

import (
	_ "embed"
	"testing"
)

//go:embed data-test.txt
var testData string

func TestFindCalibrationValue(t *testing.T) {
	expected := 142
	result := computeCalibrationValueSum(testData)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
