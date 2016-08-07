package main

import (
	"bytes"
	"fmt"
)

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	cards := make([]Card, 0)
	newDeck := Deck{cards}
	return newDeck
}

func NewStdDeck() Deck {
	newDeck := NewDeck()

	for _, s := range Suits {
		for _, r := range Ranks {
			newDeck.Append(Card{r, s})
		}
	}

	return newDeck
}

func (d *Deck) Append(c Card) {
	d.Cards = append(d.Cards, c)
}

func (d *Deck) Remove(c Card) {
	found := false
	index := -1

	for i, v := range d.Cards {
		if v == c {
			found = true
			index = i
			break
		}
	}

	if found {
		d.Cards = append(d.Cards[:index], d.Cards[index+1:]...)
	}
}

func (d *Deck) Copy() Deck {
	newDeck := NewDeck()

	for _, c := range d.Cards {
		newDeck.Cards = append(newDeck.Cards, c)
	}

	return newDeck
}

func (d *Deck) String() string {
	var buffer bytes.Buffer
	var sep string = ","

	for i, c := range d.Cards {
		if i > 0 {
			buffer.WriteString(sep)
		}
		buffer.WriteString(fmt.Sprintf("%v%v", c.Rank, c.Suit))
	}

	return buffer.String()
}
