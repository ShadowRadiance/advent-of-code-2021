package camel_cards

import (
	"testing"
)

var (
	handAAAAA = Hand{Cards: [5]Card{"A", "A", "A", "A", "A"}, Bid: 1000}
	hand26222 = Hand{Cards: [5]Card{"2", "6", "2", "2", "2"}, Bid: 1000}
	handQQQJA = Hand{Cards: [5]Card{"Q", "Q", "Q", "J", "A"}, Bid: 483}
	handT55J5 = Hand{Cards: [5]Card{"T", "5", "5", "J", "5"}, Bid: 684}
	handKK677 = Hand{Cards: [5]Card{"K", "K", "6", "7", "7"}, Bid: 28}
	handKTJJT = Hand{Cards: [5]Card{"K", "T", "J", "J", "T"}, Bid: 220}
	hand32T3K = Hand{Cards: [5]Card{"3", "2", "T", "3", "K"}, Bid: 765}
	handJ2379 = Hand{Cards: [5]Card{"J", "2", "3", "7", "9"}, Bid: 10}
)

func TestHand_Classification_Part1(t *testing.T) {
	type example struct {
		input    Hand
		expected HandClassification
	}

	examples := []example{
		{handAAAAA, FiveOfAKind},
		{hand26222, FourOfAKind},
		{handQQQJA, ThreeOfAKind},
		{handT55J5, ThreeOfAKind},
		{handKK677, TwoPair},
		{handKTJJT, TwoPair},
		{hand32T3K, OnePair},
		{handJ2379, HighCard},
	}

	for i, example := range examples {
		example.input.Part = 1
		examples[i] = example
	}

	for _, example := range examples {
		actual := example.input.Classification()
		if example.expected != actual {
			t.Errorf("Expected: %v, got: %v", example.expected, actual)
		}
	}
}

func TestHand_Classification_Part2(t *testing.T) {
	type example struct {
		input    Hand
		expected HandClassification
	}

	examples := []example{
		{handAAAAA, FiveOfAKind},
		{hand26222, FourOfAKind},
		{handQQQJA, FourOfAKind},
		{handT55J5, FourOfAKind},
		{handKK677, TwoPair},
		{handKTJJT, FourOfAKind},
		{hand32T3K, OnePair},
		{handJ2379, OnePair},
	}

	for i, example := range examples {
		example.input.Part = 2
		examples[i] = example
	}

	for _, example := range examples {
		actual := example.input.Classification()
		if example.expected != actual {
			t.Errorf("With %v, expected: %v, got: %v", example.input.Cards, example.expected, actual)
		}
	}
}

func TestHand_StrongerThan_Part1(t *testing.T) {
	type example struct {
		stronger Hand
		weaker   Hand
	}

	examples := []example{
		{handAAAAA, hand26222},
		{hand26222, handQQQJA},
		{handQQQJA, handT55J5},
		{handT55J5, handKK677},
		{handKK677, handKTJJT},
		{handKTJJT, hand32T3K},
		{hand32T3K, handJ2379},
	}

	for i, example := range examples {
		example.stronger.Part = 1
		example.weaker.Part = 1
		examples[i] = example
	}

	for _, example := range examples {
		actual := example.stronger.StrongerThan(example.weaker)
		if !actual {
			t.Errorf("Expected: %v to be stronger than: %v", example.stronger, example.weaker)
		}
	}

}

func TestHand_StrongerThan_Part2(t *testing.T) {
	type example struct {
		stronger Hand
		weaker   Hand
	}

	examples := []example{
		{handAAAAA, hand26222},
		{handQQQJA, hand26222},
		{handQQQJA, handT55J5},
		{handT55J5, handKK677},
		{handKTJJT, handKK677},
		{handKTJJT, hand32T3K},
		{hand32T3K, handJ2379}, // 3 > J
	}

	for i, example := range examples {
		example.stronger.Part = 2
		example.weaker.Part = 2
		examples[i] = example
	}

	for _, example := range examples {
		actual := example.stronger.StrongerThan(example.weaker)
		if !actual {
			t.Errorf("Expected: %v to be stronger than: %v", example.stronger, example.weaker)
		}
	}
}
