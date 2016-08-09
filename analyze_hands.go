package main

import (
	_ "fmt"
	"sort"
)

type HandData struct {
	suitCount [LastSuit]int
	rankCount [LastRank]int
}

func GenAllHands(deck *Deck) []Hand {
	h := NewHand()
	allHands := make([]Hand, 0, 2600000)

	genHandsRecursive(deck.Cards, &h, &allHands)

	return allHands
}

func genHandsRecursive(deck []Card, hand *Hand, allHands *[]Hand) {
	if hand.Full() {
		*allHands = append(*allHands, *hand)
	} else if len(deck) != 0 {
		// Check all paths that do NOT include this card
		genHandsRecursive(deck[1:], hand, allHands)

		// Check all paths that DO include this card
		newHand := hand.Copy()
		newHand.Append(deck[0])
		genHandsRecursive(deck[1:], &newHand, allHands)
	}
}

func CountHandTypes(allHands *[]Hand) [LastHandType]int {
	typeCounts := [LastHandType]int{}

	for _, hand := range *allHands {
		data := getHandData(&hand)
		handType := determineHandType(&data)
		typeCounts[handType] = typeCounts[handType] + 1
	}

	return typeCounts
}

func getHandData(h *Hand) HandData {
	// s := make([]int, len(Suits))
	// r := make([]int, len(Ranks))
	data := HandData{}

	for _, c := range h.Cards {
		data.suitCount[int(c.Suit)] = data.suitCount[int(c.Suit)] + 1
		data.rankCount[int(c.Rank)] = data.rankCount[int(c.Rank)] + 1
	}

	return data
}

func determineHandType(data *HandData) int {
	largestRankCount, secondLargestRankCount := 0, 0
	for _, count := range data.rankCount {
		if count > largestRankCount {
			secondLargestRankCount = largestRankCount
			largestRankCount = count
		} else if count > secondLargestRankCount {
			secondLargestRankCount = count
		}
	}

	ranksPresent := getRanksPresent(data.rankCount)

	// All cards have the same suit
	if suitArrayContainsValue(data.suitCount, DEFAULT_HAND_SIZE) {
		if data.rankCount[int(Ten)] == 1 &&
			data.rankCount[int(Jack)] == 1 &&
			data.rankCount[int(Queen)] == 1 &&
			data.rankCount[int(King)] == 1 &&
			data.rankCount[int(Ace)] == 1 {
			return RoyalFlush
		}

		if ranksAreSequential(ranksPresent) {
			return StraightFlush
		}

		return Flush
	}

	if rankArrayContainsValue(data.rankCount, 4) {
		return FourOfAKind
	}

	if ranksAreSequential(ranksPresent) {
		return Straight
	}

	if largestRankCount == 3 {
		if secondLargestRankCount == 2 {
			return FullHouse
		}

		return ThreeOfAKind
	}

	if largestRankCount == 2 {
		if secondLargestRankCount == 2 {
			return TwoPair
		}

		return OnePair
	}

	return NoPair
}

func getRanksPresent(counts [LastRank]int) []int {
	ranks := make([]int, 0, DEFAULT_HAND_SIZE)

	for r, c := range counts {
		if c > 0 {
			ranks = append(ranks, r)
		}
	}

	return ranks
}

func rankArrayContainsValue(counts [LastRank]int, val int) bool {
	for _, v := range counts {
		if v == val {
			return true
		}
	}

	return false
}

func suitArrayContainsValue(counts [LastSuit]int, val int) bool {
	for _, v := range counts {
		if v == val {
			return true
		}
	}

	return false
}

func ranksAreSequential(ranks RankArr) bool {
	if len(ranks) < DEFAULT_HAND_SIZE {
		return false
	}

	sort.Sort(ranks)

	if ranks[0] == Ace &&
		ranks[1] == Ten &&
		ranks[2] == Jack &&
		ranks[3] == Queen &&
		ranks[4] == King {
		return true
	}

	for i, val := range ranks {
		if i > 0 && ranks[i-1] != val-1 {
			return false
		}
	}

	return true
}
