package blackjack

import (
	"testing"

	"github.com/ccr5/blackjack-go/deck"
	"github.com/ccr5/blackjack-go/player"
)

func TestCheckAs(t *testing.T) {
	sum := checkAs(3, 7)
	if sum != 10 {
		t.Errorf("checkAs FAILED. Expect %v, got %v", 10, sum)
	}

	sum = checkAs(3, 10)
	if sum != 13 {
		t.Errorf("checkAs FAILED. Expect %v, got %v", 13, sum)
	}

	sum = checkAs(3, 15)
	if sum != 18 {
		t.Errorf("checkAs FAILED. Expect %v, got %v", 18, sum)
	}

	sum = checkAs(1, 7)
	if sum != 18 {
		t.Errorf("checkAs FAILED. Expect %v, got %v", 18, sum)
	}

	sum = checkAs(1, 10)
	if sum != 21 {
		t.Errorf("checkAs FAILED. Expect %v, got %v", 21, sum)
	}

	sum = checkAs(1, 15)
	if sum != 16 {
		t.Errorf("checkAs FAILED. Expect %v, got %v", 16, sum)
	}
}

func TestCheckWinner(t *testing.T) {
	player1 := player.Human{
		Player: player.Player{
			Name:    "Player1",
			Hand:    make([]deck.Card, 0),
			Balance: 0.0,
			Wins:    0,
			Defeats: 0,
		},
	}

	player2 := player.Human{
		Player: player.Player{
			Name:    "Player2",
			Hand:    make([]deck.Card, 0),
			Balance: 0.0,
			Wins:    0,
			Defeats: 0,
		},
	}

	cardA := deck.Card{Name: "As", Height: []int{1, 11}, Ticker: "A"}
	cardTen := deck.Card{Name: "Ten", Height: []int{10}, Ticker: "10"}
	cardFive := deck.Card{Name: "Five", Height: []int{5}, Ticker: "5"}
	cardValet := deck.Card{Name: "Valet", Height: []int{10}, Ticker: "J"}

	player1.AddCardToHand(cardA)
	player1.AddCardToHand(cardTen)

	player2.AddCardToHand(cardA)
	player2.AddCardToHand(cardValet)

	winner, verifyH1, verifyH2 := CheckWinner(player1.Hand, player2.Hand)

	if winner != "draw" || verifyH1 != 21 || verifyH2 != 21 {
		t.Errorf("CheckWinner FAILED. Expect %v, %v, %v; got %v, %v, %v", "no win", 21, 21, winner, verifyH1, verifyH2)
	}

	player1.ClearHand()
	player2.ClearHand()

	player1.AddCardToHand(cardA)
	player1.AddCardToHand(cardFive)

	player2.AddCardToHand(cardTen)
	player2.AddCardToHand(cardValet)

	winner, verifyH1, verifyH2 = CheckWinner(player1.Hand, player2.Hand)

	if winner != "h2" || verifyH1 != 16 || verifyH2 != 20 {
		t.Errorf("CheckWinner FAILED. Expect %v, %v, %v; got %v, %v, %v", "h2", 16, 20, winner, verifyH1, verifyH2)
	}

	player1.ClearHand()
	player2.ClearHand()

	player1.AddCardToHand(cardTen)
	player1.AddCardToHand(cardValet)

	player2.AddCardToHand(cardA)
	player2.AddCardToHand(cardFive)

	winner, verifyH1, verifyH2 = CheckWinner(player1.Hand, player2.Hand)

	if winner != "h1" || verifyH1 != 20 || verifyH2 != 16 {
		t.Errorf("CheckWinner FAILED. Expect %v, %v, %v; got %v, %v, %v", "h1", 20, 16, winner, verifyH1, verifyH2)
	}

	player1.ClearHand()
	player2.ClearHand()

	player1.AddCardToHand(cardTen)
	player1.AddCardToHand(cardValet)
	player1.AddCardToHand(cardFive)

	player2.AddCardToHand(cardTen)
	player2.AddCardToHand(cardValet)
	player2.AddCardToHand(cardFive)

	winner, verifyH1, verifyH2 = CheckWinner(player1.Hand, player2.Hand)

	if winner != "no win" || verifyH1 != 25 || verifyH2 != 25 {
		t.Errorf("CheckWinner FAILED. Expect %v, %v, %v; got %v, %v, %v", "no win", 25, 25, winner, verifyH1, verifyH2)
	}
}
