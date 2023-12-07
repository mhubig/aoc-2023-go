package main

import (
	"testing"
)

func TestCalculateWinningStrategies(t *testing.T) {
	given := &Race{Time: 7, Distance: 9}
	strategies := calculateWinningStrategies(given)
	expected := 4

	if strategies != expected {
		t.Errorf("Expected %d strategies, got %d", expected, strategies)
	}
}

func TestCalculateWinningStrategiesX(t *testing.T) {
	given := &Race{Time: 71530, Distance: 940200}
	strategies := calculateWinningStrategies(given)
	expected := 71503

	if strategies != expected {
		t.Errorf("Expected %d strategies, got %d", expected, strategies)
	}
}

func TestCalculateWinningStrategiesY(t *testing.T) {
	given := &Race{Time: 40828492, Distance: 233101111101487}
	strategies := calculateWinningStrategies(given)
	expected := 27102791

	if strategies != expected {
		t.Errorf("Expected %d strategies, got %d", expected, strategies)
	}
}
