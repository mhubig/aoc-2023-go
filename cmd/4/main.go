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
	CardNumber     string
	WinningNumbers []int
	GivenNumbers   []int
}

func (sc *ScratchCard) UnmarshalText(text []byte) error {
	parts := strings.Split(string(text), ":")
	sc.CardNumber = strings.Fields(parts[0])[1]

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

func (sc *ScratchCard) GetLuckyNumbers() []int {
	var luckyNumbers []int

	for i := range sc.WinningNumbers {
		if slices.Contains(sc.GivenNumbers, sc.WinningNumbers[i]) {
			luckyNumbers = append(luckyNumbers, sc.WinningNumbers[i])
		}
	}

	return luckyNumbers
}

func (sc *ScratchCard) GetPoints() (points int) {
	luckyNumbers := sc.GetLuckyNumbers()
	numberOfLuckyNumbers := len(luckyNumbers)

	if numberOfLuckyNumbers == 0 {
		return 0
	}

	return int(math.Pow(2, float64(numberOfLuckyNumbers-1)))
}

//go:embed data.txt
var data string

func main() {
	lines := strings.Split(data, "\n")

	var cards []*ScratchCard
	for i := range lines {
		card := &ScratchCard{}
		err := card.UnmarshalText([]byte(lines[i]))
		if err != nil {
			fmt.Println(err)
			return
		}

		cards = append(cards, card)
	}

	var points int
	for i := range cards {
		points += cards[i].GetPoints()
	}

	fmt.Println("Total Points:", points)
}
