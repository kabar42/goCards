package main

import (
	_ "fmt"
	"sort"
)

type HandData struct {
	suitCount []int
	rankCount []int
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
	typeCounts := make(map[string]int, 10)

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
	s := make([]int, len(Suits))
	r := make([]int, len(Ranks))
	data := HandData{s, r}

	for _, c := range h.Cards {
		data.suitCount[int(c.Suit)] = data.suitCount[int(c.Suit)] + 1
		data.rankCount[int(c.Rank)] = data.rankCount[int(c.Rank)] + 1
	}

	return data
}

func determineHandType(data HandData) string {
	var handType string = "no_pair"
	classified := false
	ranksPresent := getRanksPresent(data.rankCount)

	orderedRankCount := make([]int, len(Ranks))
	for i, count := range data.rankCount {
		orderedRankCount[i] = count
	}
	sort.Sort(RankArr(orderedRankCount))

	lenRankCount := len(orderedRankCount)
	lastRankCount := orderedRankCount[lenRankCount-1]

	// All cards have the same suit
	if countArrayContainsValue(data.suitCount, DEFAULT_HAND_SIZE) {
		if data.rankCount[int(Ten)] == 1 &&
			data.rankCount[int(Jack)] == 1 &&
			data.rankCount[int(Queen)] == 1 &&
			data.rankCount[int(King)] == 1 &&
			data.rankCount[int(Ace)] == 1 {
			handType = "royal_flush"
			classified = true
		}

		if !classified && ranksAreSequential(ranksPresent) {
			handType = "straight_flush"
			classified = true
		}

		if !classified {
			handType = "flush"
			classified = true
		}
	}

	if !classified && countArrayContainsValue(data.rankCount, 4) {
		handType = "four_of_a_kind"
		classified = true
	}

	if !classified && ranksAreSequential(ranksPresent) {
		handType = "straight"
		classified = true
	}

	if !classified && lastRankCount == 3 {
		if lenRankCount > 1 && orderedRankCount[lenRankCount-2] == 2 {
			handType = "full_house"
			classified = true
		}

		if !classified {
			handType = "three_of_a_kind"
			classified = true
		}
	}

	if !classified && lastRankCount == 2 {
		if lenRankCount > 1 && orderedRankCount[lenRankCount-2] == 2 {
			handType = "two_pair"
			classified = true
		}

		if !classified {
			handType = "one_pair"
			classified = true
		}
	}

	return handType
}

func getRanksPresent(counts []int) []int {
	ranks := make([]int, 0, DEFAULT_HAND_SIZE)

	for r, c := range counts {
		if c > 0 {
			ranks = append(ranks, r)
		}
	}

	return ranks
}

func countArrayContainsValue(counts []int, val int) bool {
	found := false
	for _, v := range counts {
		if v == val {
			found = true
			break
		}
	}

	return found
}

func ranksAreSequential(ranks RankArr) bool {
	areSeq := true
	sort.Sort(ranks)

	if len(ranks) < 5 {
		areSeq = false
		return areSeq
	}

	if ranks[0] == Ace &&
		ranks[1] == Ten &&
		ranks[2] == Jack &&
		ranks[3] == Queen &&
		ranks[4] == King {
		return areSeq
	}

	for i, val := range ranks {
		if i > 0 && ranks[i-1] != val-1 {
			areSeq = false
			break
		}
	}

	return areSeq
}
