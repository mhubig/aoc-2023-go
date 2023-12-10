package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var mapping = map[string]int{
	"J": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7,
	"8": 8, "9": 9, "T": 10, "Q": 11, "K": 12, "A": 13,
}

type Hand struct {
	Cards string
	Bid   int
	Rank  int
}

type CamelHands struct {
	Hands []*Hand
}

func (c *CamelHands) Len() int {
	return len(c.Hands)
}

func (c *CamelHands) Less(i, j int) bool {
	if c.Hands[i].Rank == c.Hands[j].Rank {
		for n := 0; n < 5; n++ {
			if mapping[c.Hands[i].Cards[n:n+1]] == mapping[c.Hands[j].Cards[n:n+1]] {
				continue
			}
			return mapping[c.Hands[i].Cards[n:n+1]] < mapping[c.Hands[j].Cards[n:n+1]]
		}
	}

	return c.Hands[i].Rank < c.Hands[j].Rank
}

func (c *CamelHands) Swap(i, j int) {
	c.Hands[i], c.Hands[j] = c.Hands[j], c.Hands[i]
}

func setRank(hand *Hand) {
	duplicates := make(map[string]int)
	for n := 0; n < 5; n++ {
		card := hand.Cards[n : n+1]
		duplicates[card]++
	}

	// 1,1,1,1,1 => 1, High card        1xJ => 2
	// 2,1,1,1   => 2, One pair         1xJ => 4, 2xJ => 4
	// 2,2,1     => 3, Two pair         1xJ => 5, 2xJ => 6
	// 3,1,1     => 4, Three of a kind  1xJ => 6, 3xJ => 6
	// 3,2       => 5, Full house       2xJ => 7, 3xJ => 7
	// 4,1       => 6, Four of a kind   1xJ => 7, 4xJ => 7
	// 5         => 7, Five of a kind

	var rank int
	switch len(duplicates) {
	case 1:
		rank = 7
	case 2:
	OuterLoop2:
		for _, v := range duplicates {
			switch v {
			case 4, 1:
				rank = 6
				break OuterLoop2
			case 3, 2:
				rank = 5
				break OuterLoop2
			}
		}
	case 3:
	OuterLoop3:
		for _, v := range duplicates {
			switch v {
			case 2:
				rank = 3
				break OuterLoop3
			case 3:
				rank = 4
				break OuterLoop3
			}
		}
	case 4:
		rank = 2
	case 5:
		rank = 1
	}

	if amount, joker := duplicates["J"]; joker {
		switch rank {
		case 1:
			rank = 2
		case 2:
			rank = 4
		case 3:
			if amount == 1 {
				rank = 5
			} else {
				rank = 6
			}
		case 4:
			rank = 6
		case 5:
			rank = 7
		case 6:
			rank = 7
		}
	}

	hand.Rank = rank
}

func readCards(data string) *CamelHands {
	camelHands := &CamelHands{}
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)
		cards := fields[0]
		bid, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		hand := &Hand{Cards: cards, Bid: bid}
		setRank(hand)
		camelHands.Hands = append(camelHands.Hands, hand)
	}

	return camelHands
}

//go:embed data.txt
var data string

func main() {
	tik := time.Now()

	cards := readCards(data)
	sort.Sort(cards)

	var totalWinnings int
	for i := 0; i < len(cards.Hands); i++ {
		totalWinnings += cards.Hands[i].Bid * (i + 1)
		fmt.Printf("%s (%d) %3d => %9d\n", cards.Hands[i].Cards, cards.Hands[i].Rank, cards.Hands[i].Bid, totalWinnings)
	}

	elapsed := time.Since(tik).Seconds()
	fmt.Println("==========================")
	fmt.Printf("Total Winning:   %9d (took %fs)\n", totalWinnings, elapsed)
}
