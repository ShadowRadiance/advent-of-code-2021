package camel_cards

import (
	"testing"
)

var (
	hand5oaK      = Hand{Cards: [5]Card{"2", "2", "2", "2", "2"}, Bid: 1000}
	hand4oaK      = Hand{Cards: [5]Card{"2", "6", "2", "2", "2"}, Bid: 1000}
	hand3oaKHi    = Hand{Cards: [5]Card{"Q", "Q", "Q", "J", "A"}, Bid: 483}
	hand3oaKLo    = Hand{Cards: [5]Card{"T", "5", "5", "J", "5"}, Bid: 684}
	handTwoPairHi = Hand{Cards: [5]Card{"K", "K", "6", "7", "7"}, Bid: 28}
	handTwoPairLo = Hand{Cards: [5]Card{"K", "T", "J", "J", "T"}, Bid: 220}
	handOnePair   = Hand{Cards: [5]Card{"3", "2", "T", "3", "K"}, Bid: 765}
	handNothing   = Hand{Cards: [5]Card{"2", "3", "7", "8", "9"}, Bid: 10}
)

func TestHand_Classification(t *testing.T) {
	type example struct {
		input    Hand
		expected HandClassification
	}

	examples := []example{
		{hand5oaK, FiveOfAKind},
		{hand4oaK, FourOfAKind},
		{hand3oaKHi, ThreeOfAKind},
		{hand3oaKLo, ThreeOfAKind},
		{handTwoPairHi, TwoPair},
		{handTwoPairLo, TwoPair},
		{handOnePair, OnePair},
		{handNothing, HighCard},
	}

	for _, example := range examples {
		actual := example.input.Classification()
		if example.expected != actual {
			t.Errorf("Expected: %v, got: %v", example.expected, actual)
		}
	}
}

func TestHand_StrongerThan(t *testing.T) {
	type example struct {
		stronger Hand
		weaker   Hand
	}

	examples := []example{
		{hand5oaK, hand4oaK},
		{hand4oaK, hand3oaKHi},
		{hand3oaKHi, hand3oaKLo},
		{hand3oaKLo, handTwoPairHi},
		{handTwoPairHi, handTwoPairLo},
		{handTwoPairLo, handOnePair},
		{handOnePair, handNothing},
	}

	for _, example := range examples {
		actual := example.stronger.StrongerThan(example.weaker)
		if !actual {
			t.Errorf("Expected: %v to be stronger than: %v", example.stronger, example.weaker)
		}
	}
}
