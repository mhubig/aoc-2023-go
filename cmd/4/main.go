package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type ScratchCard struct {
	CardNumber     int
	WinningNumbers []int
	GivenNumbers   []int
	Wins           int
	Points         int
	Copies         int
}

func (sc *ScratchCard) String() string {
	return fmt.Sprintf("Card %d (%d): %v | %v", sc.CardNumber, sc.Copies, sc.WinningNumbers, sc.GivenNumbers)
}

func (sc *ScratchCard) UnmarshalText(text []byte) error {
	parts := strings.Split(string(text), ":")

	var err error
	sc.CardNumber, err = strconv.Atoi(strings.Fields(parts[0])[1])
	if err != nil {
		return err
	}

	anumbers := strings.Split(parts[1], "|")
	wnumbers := strings.Fields(anumbers[0])
	gnumbers := strings.Fields(anumbers[1])

	for i := range wnumbers {
		number, err := strconv.Atoi(wnumbers[i])
		if err != nil {
			return err
		}

		sc.WinningNumbers = append(sc.WinningNumbers, number)
	}

	for i := range gnumbers {
		number, err := strconv.Atoi(gnumbers[i])
		if err != nil {
			return err
		}

		sc.GivenNumbers = append(sc.GivenNumbers, number)
	}

	return nil
}

func setWins(sc *ScratchCard) {
	var luckyNumbers []int

	for i := range sc.WinningNumbers {
		if slices.Contains(sc.GivenNumbers, sc.WinningNumbers[i]) {
			luckyNumbers = append(luckyNumbers, sc.WinningNumbers[i])
		}
	}

	sc.Wins = len(luckyNumbers)
}

func setPoints(sc *ScratchCard) {
	if sc.Wins == 0 {
		sc.Points = 0
	} else {
		sc.Points = int(math.Pow(2, float64(sc.Wins-1)))
	}
}

func winScatchCards(cards []ScratchCard) (result int) {
	for i := range cards {
		cards[i].Copies += 1

		for j := 1; j <= cards[i].Wins; j++ {
			cards[i+j].Copies += cards[i].Copies
		}
	}

	for i := range cards {
		result += cards[i].Copies
	}

	return result
}

func readScratchCards(data string) (cards []ScratchCard) {
	lines := strings.Split(data, "\n")

	for i := range lines {
		card := ScratchCard{}
		err := card.UnmarshalText([]byte(lines[i]))
		if err != nil {
			fmt.Println(err)
			return cards
		}

		setWins(&card)
		setPoints(&card)

		cards = append(cards, card)
	}

	return cards
}

func calculatePoints(cards []ScratchCard) (points int) {
	for i := range cards {
		points += cards[i].Points
	}

	return points
}

//go:embed data.txt
var data string

func main() {
	cards := readScratchCards(data)
	points := calculatePoints(cards)
	result := winScatchCards(cards)

	fmt.Println("Total Points:", points)
	fmt.Println("Total Cards: ", result)
}
