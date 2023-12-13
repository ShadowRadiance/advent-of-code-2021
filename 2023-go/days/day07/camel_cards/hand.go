package camel_cards

type Hand struct {
	Cards [5]Card
	Bid   int
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
	for _, count := range counts {
		if count == 5 {
			return true
		}
	}
	return false
}
func (hand Hand) isFourOfAKind() bool {
	counts := hand.GroupCards()
	for _, count := range counts {
		if count == 4 {
			return true
		}
	}
	return false
}
func (hand Hand) isFullHouse() bool {
	counts := hand.GroupCards()
	hasThreeOfAKind := false
	for _, count := range counts {
		if count == 3 {
			hasThreeOfAKind = true
		}
	}
	if hasThreeOfAKind {
		for _, count := range counts {
			if count == 2 {
				return true // 3 of one, 2 of another
			}
		}
		return false // its just a 3OAK
	}
	return false // not even 3OAK
}
func (hand Hand) isThreeOfAKind() bool {
	counts := hand.GroupCards()
	hasThreeOfAKind := false
	for _, count := range counts {
		if count == 3 {
			hasThreeOfAKind = true
		}
	}
	if hasThreeOfAKind {
		for _, count := range counts {
			if count == 2 {
				return false // it is a full house
			}
		}
		return true // its a 3OAK
	}
	return false // not even a 3OAK
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
}
func (hand Hand) isOnePair() bool {
	counts := hand.GroupCards()
	pairs := 0
	for _, count := range counts {
		if count == 2 {
			pairs++
		}
	}
	return pairs == 1
}

func (hand Hand) StrongerThan(other Hand) bool {
	classificationDifference := hand.Classification() - other.Classification()
	if classificationDifference == 0 {
		for i := range hand.Cards {
			if hand.Cards[i].Rank() == other.Cards[i].Rank() {
				continue
			}
			return hand.Cards[i].Rank() > other.Cards[i].Rank()
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
