package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type GameSet struct {
	Red   int
	Green int
	Blue  int
}

// 3 blue
func (gs *GameSet) getNumber(text string) int {
	text = strings.Trim(text, " ")
	parts := strings.Split(text, " ")

	number, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return number
}

// 1 red, 2 green, 3 blue
func (gs *GameSet) UnmarshalText(text []byte) error {
	colors := strings.Split(string(text), ",")

	for i := range colors {
		if strings.Contains(colors[i], "red") {
			gs.Red = gs.getNumber(colors[i])
		} else if strings.Contains(colors[i], "green") {
			gs.Green = gs.getNumber(colors[i])
		} else if strings.Contains(colors[i], "blue") {
			gs.Blue = gs.getNumber(colors[i])
		}
	}
	return nil
}

type Game struct {
	Id   int
	Sets []*GameSet
}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func (g *Game) UnmarshalText(text []byte) error {
	split := strings.Split(string(text), ":")
	gameID := strings.Split(split[0], " ")[1]

	id, err := strconv.Atoi(gameID)
	if err != nil {
		return err
	}

	g.Id = id

	gameSets := strings.Split(split[1], ";")
	for i := range gameSets {
		gameSet := &GameSet{}
		err := gameSet.UnmarshalText([]byte(gameSets[i]))
		if err != nil {
			return err
		}
		g.Sets = append(g.Sets, gameSet)
	}

	return nil
}

func readListOfGames(text string) []*Game {
	var games []*Game

	lines := strings.Split(text, "\n")
	for i := range lines {
		game := &Game{}
		err := game.UnmarshalText([]byte(lines[i]))
		if err != nil {
			fmt.Println(err)
			continue
		}
		games = append(games, game)
	}

	return games
}

func gameIsPossible(game *Game, refSet *GameSet) bool {
	for i := range game.Sets {
		if game.Sets[i].Red > refSet.Red {
			return false
		}

		if game.Sets[i].Green > refSet.Green {
			return false
		}

		if game.Sets[i].Blue > refSet.Blue {
			return false
		}
	}

	return true
}

//go:embed data.txt
var data string

func main() {
	refGameSet := &GameSet{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	games := readListOfGames(data)
	var possibleGames []int
	for i := range games {
		if gameIsPossible(games[i], refGameSet) {
			fmt.Println(games[i].Id)
			possibleGames = append(possibleGames, games[i].Id)
		}
	}

	var result int
	for i := range possibleGames {
		result += possibleGames[i]
	}

	fmt.Println("Sum of possible game ID's:", result)
}
