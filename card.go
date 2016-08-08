package main

import (
	"fmt"
)

type RankT int

const (
	Ace = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var Ranks = []RankT{
	Ace,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Ten,
	Jack,
	Queen,
	King,
}

// Used for sorting
type RankArr []int

func (r RankArr) Len() int {
	return len(r)
}

func (r RankArr) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r RankArr) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r RankT) String() string {
	var value string

	switch r {
	case Ace:
		value = "A"
	case Two:
		value = "2"
	case Three:
		value = "3"
	case Four:
		value = "4"
	case Five:
		value = "5"
	case Six:
		value = "6"
	case Seven:
		value = "7"
	case Eight:
		value = "8"
	case Nine:
		value = "9"
	case Ten:
		value = "10"
	case Jack:
		value = "J"
	case Queen:
		value = "Q"
	case King:
		value = "K"
	}

	return value
}

type SuitT int

const (
	Hearts = iota
	Clubs
	Diamonds
	Spades
)

var Suits = []SuitT{
	Hearts,
	Clubs,
	Diamonds,
	Spades,
}

func (s SuitT) String() string {
	var value string

	switch s {
	case Hearts:
		value = "H"
	case Clubs:
		value = "C"
	case Diamonds:
		value = "D"
	case Spades:
		value = "S"
	}

	return value
}

type Card struct {
	Rank RankT
	Suit SuitT
}

func (c *Card) String() string {
	cardStr := fmt.Sprintf("%v %v", c.Rank, c.Suit)
	return cardStr
}
