package main

import (
	"sort"
)

type HandData struct {
	suitCount map[SuitT]int
	rankCount map[RankT]int
}

func GenAllHands(deck *Deck) []Hand {
	h := NewHand()
	allHands := make([]Hand, 0, 2600000)

	genHandsRecursive(deck.Cards, &h, &allHands)

	return allHands
}

func genHandsRecursive(deck []Card, hand *Hand, allHands *[]Hand) {
	if hand.Full() {
		newHand := hand.Copy()
		*allHands = append(*allHands, newHand)
	} else if len(deck) != 0 {
		nextCard := deck[0]

		newHand := hand.Copy()

		// Check all paths that do NOT include this card
		genHandsRecursive(deck[1:], &newHand, allHands)

		// Check all paths that DO include this card
		newHand.Append(nextCard)
		genHandsRecursive(deck[1:], &newHand, allHands)
	}
}

func CountHandTypes(allHands *[]Hand) map[string]int {
	typeCounts := make(map[string]int)

	for _, hand := range *allHands {
		// Get the counts for this hand
		data := getHandData(hand)

		// Classify the hand
		handType := determineHandType(data)

		// Add the data to typeCounts
		if val, ok := typeCounts[handType]; ok == true {
			typeCounts[handType] = val + 1
		} else {
			typeCounts[handType] = 1
		}
	}

	return typeCounts
}

func getHandData(h Hand) HandData {
	s := make(map[SuitT]int)
	r := make(map[RankT]int)
	data := HandData{s, r}

	for _, c := range h.Cards {
		val, ok := data.suitCount[c.Suit]
		if ok {
			data.suitCount[c.Suit] = val + 1
		} else {
			data.suitCount[c.Suit] = 1
		}

		val, ok = data.rankCount[c.Rank]
		if ok {
			data.rankCount[c.Rank] = val + 1
		} else {
			data.rankCount[c.Rank] = 1
		}
	}

	return data
}

func determineHandType(data HandData) string {
	var handType string = "no_pair"
	ranksPresent := getRanksPresent(data.rankCount)

	for _, count := range data.suitCount {
		if count == 5 {
			if len(ranksPresent) == 5 {
				if data.rankCount[Ten] == 1 &&
					data.rankCount[Jack] == 1 &&
					data.rankCount[Queen] == 1 &&
					data.rankCount[King] == 1 &&
					data.rankCount[Ace] == 1 {
					handType = "royal_flush"
					break
				}

				if ranksAreSequential(ranksPresent) {
					handType = "straight_flush"
					break
				}
			}

			handType = "flush"
			break
		}
	}

	return handType
}

func getRanksPresent(counts map[RankT]int) []RankT {
	ranks := make([]RankT, 0)

	for r, _ := range counts {
		found := false
		for _, val := range ranks {
			if val == r {
				found = true
				break
			}
		}
		if !found {
			ranks = append(ranks, r)
		}
	}

	return ranks
}

func ranksAreSequential(ranks RankArr) bool {
	areSeq := true
	sort.Sort(ranks)

	if len(ranks) < 2 {
		areSeq = false
		return areSeq
	}

	for i, val := range ranks {
		if i > 0 {
			if ranks[i-1] != val-1 {
				areSeq = false
				break
			}
		}
	}

	return areSeq
}
