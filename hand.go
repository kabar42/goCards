package main

import (
	_ "fmt"
)

const DEFAULT_HAND_SIZE int = 5

type Hand struct {
	Cards []Card
	Size  int
}

func NewHand() Hand {
	cards := make([]Card, 0, DEFAULT_HAND_SIZE)
	newHand := Hand{cards, DEFAULT_HAND_SIZE}
	return newHand
}

func (h *Hand) Append(c Card) {
	h.Cards = append(h.Cards, c)
}

func (h *Hand) Copy() Hand {
	newHand := NewHand()

	for _, c := range h.Cards {
		newHand.Cards = append(newHand.Cards, c)
	}

	return newHand
}

func (h *Hand) Full() bool {
	return len(h.Cards) >= h.Size
}
