package camel_cards

type Hand struct {
	Cards [5]Card
	Bid   int

	Part int
}

type HandClassification int

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (hand Hand) Classification() HandClassification {
	switch {
	case hand.isFiveOfAKind():
		return FiveOfAKind
	case hand.isFourOfAKind():
		return FourOfAKind
	case hand.isFullHouse():
		return FullHouse
	case hand.isThreeOfAKind():
		return ThreeOfAKind
	case hand.isTwoPair():
		return TwoPair
	case hand.isOnePair():
		return OnePair
	default:
		return HighCard
	}
}

func (hand Hand) isFiveOfAKind() bool {
	counts := hand.GroupCards()
	jokers := 0
	if hand.Part == 2 {
		jokers = counts["J"]
		delete(counts, "J")
	}
	for _, count := range counts {
		if count+jokers == 5 {
			return true
		}
	}

	return jokers == 5
}

func (hand Hand) isFourOfAKind() bool {
	counts := hand.GroupCards()
	jokers := 0
	if hand.Part == 2 {
		jokers = counts["J"]
		delete(counts, "J")
	}
	for _, count := range counts {
		if count+jokers == 4 {
			return true
		}
	}
	return jokers == 4
}

func (hand Hand) isFullHouse() bool {
	counts := hand.GroupCards()
	jokers := 0
	if hand.Part == 2 {
		jokers = counts["J"]
		delete(counts, "J")
	}

	triples := 0
	pairs := 0
	for _, count := range counts {
		if count == 3 {
			triples++
		}
		if count == 2 {
			pairs++
		}
	}

	return (triples == 1 && pairs == 1) ||
		(triples == 1 && pairs == 0 && jokers == 1) ||
		(triples == 0 && pairs == 2 && jokers == 1) ||
		(triples == 0 && pairs == 1 && jokers == 2) ||
		(triples == 0 && pairs == 0 && jokers == 3)
}

func (hand Hand) isThreeOfAKind() bool {
	counts := hand.GroupCards()
	jokers := 0
	if hand.Part == 2 {
		jokers = counts["J"]
		delete(counts, "J")
	}
	for _, count := range counts {
		if count+jokers == 3 {
			return true
		}
	}
	return jokers == 3
}

func (hand Hand) isTwoPair() bool {
	counts := hand.GroupCards()
	pairs := 0
	for _, count := range counts {
		if count == 2 {
			pairs++
		}
	}
	return pairs == 2
	// interesting Jokers don't matter
	// 0 jokers: would need two actual pairs
	// 1 jokers: would need one real pair... but then it'd be a 3OAK
	// 2 jokers: ... any card becomes a 3OAK
	// ...
}

func (hand Hand) isOnePair() bool {
	counts := hand.GroupCards()
	jokers := 0
	if hand.Part == 2 {
		jokers = counts["J"]
		delete(counts, "J")
	}
	pairs := 0
	for _, count := range counts {
		if count == 2 {
			pairs++
		}
	}
	return pairs == 1 || jokers == 1
}

func (hand Hand) StrongerThan(other Hand) bool {
	classificationDifference := hand.Classification() - other.Classification()
	if classificationDifference == 0 {
		for i := range hand.Cards {
			if hand.Cards[i].Rank(hand.Part) == other.Cards[i].Rank(hand.Part) {
				continue
			}
			return hand.Cards[i].Rank(hand.Part) > other.Cards[i].Rank(hand.Part)
		}
		panic("Identical hands!")
	}
	return classificationDifference > 0
}

func (hand Hand) GroupCards() (groups map[Card]int) {
	groups = make(map[Card]int)
	for _, card := range hand.Cards {
		groups[card] += 1
	}
	return
}

type ByStrength []Hand

func (hands ByStrength) Len() int           { return len(hands) }
func (hands ByStrength) Less(i, j int) bool { return !hands[i].StrongerThan(hands[j]) }
func (hands ByStrength) Swap(i, j int)      { hands[i], hands[j] = hands[j], hands[i] }
