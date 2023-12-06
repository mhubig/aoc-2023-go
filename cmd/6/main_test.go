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
