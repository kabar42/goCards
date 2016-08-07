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

func (r RankT) String() string {
	var value string

	switch r {
	case Ace:
		value = "Ace"
	case Two:
		value = "Two"
	case Three:
		value = "Three"
	case Four:
		value = "Four"
	case Five:
		value = "Five"
	case Six:
		value = "Six"
	case Seven:
		value = "Seven"
	case Eight:
		value = "Eight"
	case Nine:
		value = "Nine"
	case Ten:
		value = "Ten"
	case Jack:
		value = "Jack"
	case Queen:
		value = "Queen"
	case King:
		value = "King"
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
		value = "Hearts"
	case Clubs:
		value = "Clubs"
	case Diamonds:
		value = "Diamonds"
	case Spades:
		value = "Spades"
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
