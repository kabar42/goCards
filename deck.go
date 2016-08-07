package main

const DEFAULT_SIZE uint32 = 52

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
