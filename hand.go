package main

import (
	_ "fmt"
)

const DEFAULT_HAND_SIZE int = 5

const (
	RoyalFlush = iota
	StraightFlush
	FourOfAKind
	FullHouse
	Flush
	Straight
	ThreeOfAKind
	TwoPair
	OnePair
	NoPair
	LastHandType
)

func HandToString(handType int) string {
	var value string

	switch handType {
	case RoyalFlush:
        value = "RoyalFlush"
	case StraightFlush:
        value = "StraightFlush"
	case FourOfAKind:
        value = "FourOfAKind"
	case FullHouse:
        value = "FullHouse"
	case Flush:
        value = "Flush"
	case Straight:
        value = "Straight"
	case ThreeOfAKind:
        value = "ThreeOfAKind"
	case TwoPair:
        value = "TwoPair"
	case OnePair:
        value = "OnePair"
	case NoPair:
        value = "NoPair"
	}

	return value
}

type Hand struct {
	Cards [DEFAULT_HAND_SIZE]Card
	Size  int
}

func NewHand() Hand {
	return Hand{}
}

func (h *Hand) Append(c Card) {
	if !h.Full() {
		h.Cards[h.Size] = c
		h.Size = h.Size + 1
	}
}

func (h *Hand) Copy() Hand {
	newHand := NewHand()

	for i, c := range h.Cards {
		newHand.Cards[i] = c
		newHand.Size = h.Size
	}

	return newHand
}

func (h *Hand) Full() bool {
	return h.Size >= DEFAULT_HAND_SIZE
}
