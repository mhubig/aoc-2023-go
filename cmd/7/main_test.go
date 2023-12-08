package main

import (
	"sort"
	"testing"
)

func TestCamelCards_Len(t *testing.T) {
	type fields struct {
		Hands []*Hand
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Zero", fields{[]*Hand{}}, 0},
		{"One", fields{[]*Hand{{}}}, 1},
		{"Two", fields{[]*Hand{{}, {}}}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CamelHands{
				Hands: tt.fields.Hands,
			}
			if got := c.Len(); got != tt.want {
				t.Errorf("CamelCards.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamelCards_Less(t *testing.T) {
	given := &CamelHands{
		Hands: []*Hand{{Rank: 1}, {Rank: 2}},
	}

	if !given.Less(0, 1) {
		t.Errorf("CamelCards.Less() = %v, want %v", false, true)
	}

	if given.Less(1, 0) {
		t.Errorf("CamelCards.Less() = %v, want %v", true, false)
	}
}

func TestCamelCards_Swap(t *testing.T) {
	given := &CamelHands{
		Hands: []*Hand{{Rank: 1}, {Rank: 2}},
	}

	given.Swap(0, 1)

	if rank := given.Hands[0].Rank; rank != 2 {
		t.Errorf("CamelCards.Swap() = %v, want %v", rank, 2)
	}

	if rank := given.Hands[1].Rank; rank != 1 {
		t.Errorf("CamelCards.Swap() = %v, want %v", rank, 1)
	}
}

func TestCamelCards_GivenDifferentRanks_SortByRank(t *testing.T) {
	given := &CamelHands{
		Hands: []*Hand{{Rank: 2}, {Rank: 1}},
	}

	sort.Sort(given)

	if rank := given.Hands[0].Rank; rank != 1 {
		t.Errorf("Got %v, want %v", rank, 1)
	}

	if rank := given.Hands[1].Rank; rank != 2 {
		t.Errorf("Got %v, want %v", rank, 2)
	}
}

func TestCamelCards_GivenSameRanksOneCard_SortByStrength(t *testing.T) {
	given := &CamelHands{
		Hands: []*Hand{{Cards: "A", Rank: 1}, {Cards: "K", Rank: 1}},
	}

	sort.Sort(given)

	if cards := given.Hands[0].Cards; cards != "K" {
		t.Errorf("Got %s, want %s", cards, "K")
	}

	if cards := given.Hands[1].Cards; cards != "A" {
		t.Errorf("Got %s, want %s", cards, "A")
	}
}

func TestCamelCards_GivenSameRanksTwoCards_SortByStrength(t *testing.T) {
	given := &CamelHands{
		Hands: []*Hand{{Cards: "AA", Rank: 1}, {Cards: "AK", Rank: 1}},
	}

	sort.Sort(given)

	if cards := given.Hands[0].Cards; cards != "AK" {
		t.Errorf("Got %s, want %s", cards, "AK")
	}

	if cards := given.Hands[1].Cards; cards != "AA" {
		t.Errorf("Got %s, want %s", cards, "AA")
	}
}

func TestCamelCards_GivenSameRankesFiveCards_SortByStrength(t *testing.T) {
	given := &CamelHands{
		Hands: []*Hand{
			{Cards: "32T3K", Rank: 2},
			{Cards: "T55J5", Rank: 4},
			{Cards: "KK677", Rank: 3},
			{Cards: "KTJJT", Rank: 3},
			{Cards: "QQQJA", Rank: 4},
		},
	}

	sort.Sort(given)

	if cards := given.Hands[0].Cards; cards != "32T3K" {
		t.Errorf("Got %s, want %s", cards, "32T3K")
	}

	if cards := given.Hands[1].Cards; cards != "KTJJT" {
		t.Errorf("Got %s, want %s", cards, "KTJJT")
	}

	if cards := given.Hands[2].Cards; cards != "KK677" {
		t.Errorf("Got %s, want %s", cards, "KK677")
	}

	if cards := given.Hands[3].Cards; cards != "T55J5" {
		t.Errorf("Got %s, want %s", cards, "T55J5")
	}

	if cards := given.Hands[4].Cards; cards != "QQQJA" {
		t.Errorf("Got %s, want %s", cards, "QQQJA")
	}
}

func TestSetRank_GivenHandWithFiveDifferentCards_RankIsHighCard(t *testing.T) {
	given := &Hand{Cards: "23456"}

	setRank(given)

	if rank := given.Rank; rank != 1 {
		t.Errorf("Got %v, want %v", rank, 1)
	}
}

func TestSetRank_GivenHandWithOnePair_RankIsOnePair(t *testing.T) {
	given := &Hand{Cards: "A23A4"}

	setRank(given)

	if rank := given.Rank; rank != 2 {
		t.Errorf("Got %v, want %v", rank, 2)
	}
}

func TestSetRank_GivenHandWithTwoPairs_RankIsTwoPairs(t *testing.T) {
	given := &Hand{Cards: "23432"}

	setRank(given)

	if rank := given.Rank; rank != 3 {
		t.Errorf("Got %v, want %v", rank, 3)
	}
}

func TestSetRank_GivenHandWithThreeOfAKind_RankIsThreeOfAKind(t *testing.T) {
	given := &Hand{Cards: "TTT98"}

	setRank(given)

	if rank := given.Rank; rank != 4 {
		t.Errorf("Got %v, want %v", rank, 4)
	}
}

func TestSetRank_GivenHandWithFullHouse_RankIsFullHouse(t *testing.T) {
	given := &Hand{Cards: "23332"}

	setRank(given)

	if rank := given.Rank; rank != 5 {
		t.Errorf("Got %v, want %v", rank, 5)
	}
}

func TestSetRank_GivenHandWithFourOfAKind_RankIsFourOfAKind(t *testing.T) {
	given := &Hand{Cards: "AA8AA"}

	setRank(given)

	if rank := given.Rank; rank != 6 {
		t.Errorf("Got %v, want %v", rank, 6)
	}
}

func TestSetRank_GivenHandWithFiveOfAKind_RankIsFiveOfAKind(t *testing.T) {
	given := &Hand{Cards: "AAAAA"}

	setRank(given)

	if rank := given.Rank; rank != 7 {
		t.Errorf("Got %v, want %v", rank, 7)
	}
}
