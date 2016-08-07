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
	cards := make([]Card, 0)
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

func GenAllHands(deck *Deck, hand *Hand, allHands *[]Hand) {
	if hand.Full() {
		newHand := hand.Copy()
		*allHands = append(*allHands, newHand)
	} else if len(deck.Cards) != 0 {
		nextCard := deck.Cards[0]
		deckCopy := deck.Copy()
		deckCopy.Remove(nextCard)

		newHand := hand.Copy()

		// Check all paths that do NOT include this card
		GenAllHands(&deckCopy, &newHand, allHands)

		// Check all paths that DO include this card
		newHand.Append(nextCard)
		GenAllHands(&deckCopy, &newHand, allHands)
	}
}