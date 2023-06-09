package blackjack

import "github.com/ccr5/blackjack-go/deck"

func checkAs(countA, sum int) int {
	if countA == 1 && sum <= 10 {
		return sum + 11
	} else {
		return sum + 1*countA
	}
}

func CheckWinner(hand1, hand2 []deck.Card) (string, int, int) {
	h1 := 0
	h2 := 0
	countH1 := 0
	countH2 := 0

	for _, card := range hand1 {
		if card.Ticker == "A" {
			countH1 += 1
		} else {
			h1 += card.Height[0]
		}
	}

	for _, card := range hand2 {
		if card.Ticker == "A" {
			countH2 += 1
		} else {
			h2 += card.Height[0]
		}
	}

	verifyH1 := checkAs(countH1, h1)
	verifyH2 := checkAs(countH2, h2)

	if verifyH1 > 21 && verifyH2 > 21 {
		return "no win", verifyH1, verifyH2
	} else if verifyH1 <= 21 && verifyH2 > 21 {
		return "h1", verifyH1, verifyH2
	} else if h1 > 21 && h2 <= 21 {
		return "h2", verifyH1, verifyH2
	} else if h1 == h2 {
		return "draw", verifyH1, verifyH2
	} else if h1 == 21 {
		return "h1", verifyH1, verifyH2
	} else if h2 == 21 {
		return "h2", verifyH1, verifyH2
	} else if h1 > h2 {
		return "h1", verifyH1, verifyH2
	} else if h2 > h1 {
		return "h2", verifyH1, verifyH2
	} else {
		panic("Error to find a winner")
	}
}
