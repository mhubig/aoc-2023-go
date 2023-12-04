package main

import (
	_ "embed"
	"testing"
)

//go:embed data-test.txt
var testData []byte

func TestUnmarshalText_GameWithoutGameSets(t *testing.T) {
	given := "Game 1:"
	expected := &Game{
		Id:   1,
		Sets: []*GameSet{},
	}

	game := &Game{}
	err := game.UnmarshalText([]byte(given))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if game.Id != expected.Id {
		t.Errorf("Expected %d, got %d", expected.Id, game.Id)
	}
}

func TestUnmarshalText_GameWithOneGameSet(t *testing.T) {
	given := "Game 1: 3 blue, 4 red"
	expected := &Game{
		Id: 1,
		Sets: []*GameSet{
			{
				Red:  4,
				Blue: 3,
			},
		},
	}

	game := &Game{}
	err := game.UnmarshalText([]byte(given))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if game.Id != expected.Id {
		t.Errorf("Expected %d, got %d", expected.Id, game.Id)
	}

	if len(game.Sets) != len(expected.Sets) {
		t.Errorf("Expected %d, got %d", len(expected.Sets), len(game.Sets))
	}

	if game.Sets[0].Red != expected.Sets[0].Red {
		t.Errorf("Expected %d, got %d", expected.Sets[0].Red, game.Sets[0].Red)
	}

	if game.Sets[0].Blue != expected.Sets[0].Blue {
		t.Errorf("Expected %d, got %d", expected.Sets[0].Blue, game.Sets[0].Blue)
	}
}

func TestUnmarshalText_GameWithTwoGameSets(t *testing.T) {
	given := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue"
	expected := &Game{
		Id: 1,
		Sets: []*GameSet{
			{
				Red:  4,
				Blue: 3,
			},
			{
				Red:   1,
				Green: 2,
				Blue:  6,
			},
		},
	}

	game := &Game{}
	err := game.UnmarshalText([]byte(given))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if game.Id != expected.Id {
		t.Errorf("Expected %d, got %d", expected.Id, game.Id)
	}

	if len(game.Sets) != len(expected.Sets) {
		t.Errorf("Expected %d, got %d", len(expected.Sets), len(game.Sets))
	}

	if game.Sets[0].Red != expected.Sets[0].Red {
		t.Errorf("Expected %d, got %d", expected.Sets[0].Red, game.Sets[0].Red)
	}

	if game.Sets[0].Blue != expected.Sets[0].Blue {
		t.Errorf("Expected %d, got %d", expected.Sets[0].Blue, game.Sets[0].Blue)
	}

	if game.Sets[1].Red != expected.Sets[1].Red {
		t.Errorf("Expected %d, got %d", expected.Sets[1].Red, game.Sets[1].Red)
	}

	if game.Sets[1].Green != expected.Sets[1].Green {
		t.Errorf("Expected %d, got %d", expected.Sets[1].Green, game.Sets[1].Green)
	}

	if game.Sets[1].Blue != expected.Sets[1].Blue {
		t.Errorf("Expected %d, got %d", expected.Sets[1].Blue, game.Sets[1].Blue)
	}
}

func TestUnmarshalText_ListOfGames(t *testing.T) {
	given := string(testData)
	expected := 5

	games := readListOfGames(given)

	if len(games) != expected {
		t.Errorf("Expected %d, got %d", expected, len(games))
	}
}
