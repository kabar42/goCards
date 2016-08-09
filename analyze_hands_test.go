package main

import (
    "testing"
)


func getTypeFromCards(cards [DEFAULT_HAND_SIZE]Card) int {
    hand := Hand{ cards, len(cards) }
    data := GetHandData(&hand)
    handType := DetermineHandType(&data)

    return handType
}

func assertHandTypeEquals(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected: %v, Got: %v\n", HandToString(expected), HandToString(actual))
    }
}

func TestRoyalFlush(t *testing.T) {
    cards := [DEFAULT_HAND_SIZE]Card{
        Card{Ace,Spades},
        Card{King,Spades},
        Card{Queen,Spades},
        Card{Jack,Spades},
        Card{Ten,Spades},
        }
    handType := getTypeFromCards(cards)
    expected := RoyalFlush
    assertHandTypeEquals(t, expected, handType)
}

func TestStraightFlush(t *testing.T) {
    cards := [DEFAULT_HAND_SIZE]Card{
        Card{King,Spades},
        Card{Queen,Spades},
        Card{Jack,Spades},
        Card{Ten,Spades},
        Card{Nine,Spades},
        }
    handType := getTypeFromCards(cards)
    expected := StraightFlush
    assertHandTypeEquals(t, expected, handType)
}

func TestFourOfAKind(t *testing.T) {
    cards := [DEFAULT_HAND_SIZE]Card{
        Card{King,Spades},
        Card{King,Diamonds},
        Card{King,Clubs},
        Card{King,Hearts},
        Card{Nine,Spades},
        }
    handType := getTypeFromCards(cards)
    expected := FourOfAKind
    assertHandTypeEquals(t, expected, handType)
}

func TestFullHouse(t *testing.T) {
    cards := [DEFAULT_HAND_SIZE]Card{
        Card{King,Spades},
        Card{King,Diamonds},
        Card{King,Clubs},
        Card{Nine,Hearts},
        Card{Nine,Spades},
        }
    handType := getTypeFromCards(cards)
    expected := FullHouse
    assertHandTypeEquals(t, expected, handType)
}

func TestFlush(t *testing.T) {
    cards := [DEFAULT_HAND_SIZE]Card{
        Card{King,Hearts},
        Card{Jack,Hearts},
        Card{Ten,Hearts},
        Card{Nine,Hearts},
        Card{Two,Hearts},
        }
    handType := getTypeFromCards(cards)
    expected := Flush
    assertHandTypeEquals(t, expected, handType)
}

func TestStraight(t *testing.T) {
    cards := [DEFAULT_HAND_SIZE]Card{
        Card{Queen,Hearts},
        Card{Jack,Hearts},
        Card{Ten,Clubs},
        Card{Nine,Hearts},
        Card{Eight,Hearts},
        }
    handType := getTypeFromCards(cards)
    expected := Straight
    assertHandTypeEquals(t, expected, handType)
}

func TestThreeOfAKind(t *testing.T) {
    cards := [DEFAULT_HAND_SIZE]Card{
        Card{Queen,Hearts},
        Card{Queen,Spades},
        Card{Queen,Clubs},
        Card{Nine,Hearts},
        Card{Eight,Hearts},
        }
    handType := getTypeFromCards(cards)
    expected := ThreeOfAKind
    assertHandTypeEquals(t, expected, handType)
}

func TestTwoPair(t *testing.T) {
    cards := [DEFAULT_HAND_SIZE]Card{
        Card{Queen,Hearts},
        Card{Queen,Spades},
        Card{Nine,Clubs},
        Card{Nine,Hearts},
        Card{Eight,Hearts},
        }
    handType := getTypeFromCards(cards)
    expected := TwoPair
    assertHandTypeEquals(t, expected, handType)
}

func TestOnePair(t *testing.T) {
    cards := [DEFAULT_HAND_SIZE]Card{
        Card{Queen,Hearts},
        Card{Jack,Spades},
        Card{Nine,Clubs},
        Card{Nine,Hearts},
        Card{Eight,Hearts},
        }
    handType := getTypeFromCards(cards)
    expected := OnePair
    assertHandTypeEquals(t, expected, handType)
}

func TestNoPair(t *testing.T) {
    cards := [DEFAULT_HAND_SIZE]Card{
        Card{Queen,Hearts},
        Card{Jack,Hearts},
        Card{Nine,Clubs},
        Card{Eight,Hearts},
        Card{Seven,Hearts},
        }
    handType := getTypeFromCards(cards)
    expected := NoPair
    assertHandTypeEquals(t, expected, handType)
}
