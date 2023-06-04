package blackjack

import (
	"testing"

	"github.com/ccr5/blackjack-go/deck"
	"github.com/ccr5/blackjack-go/player"
)

func TestCheckBalance(t *testing.T) {
	bj := Blackjack{}

	player1 := player.Human{
		Player: player.Player{
			Name:    "Player1",
			Hand:    make([]deck.Card, 0),
			Balance: 0.0,
			Wins:    0,
			Defeats: 0,
		},
	}

	result, err := bj.checkBalance(player1.Player)
	if err != nil || !result {
		t.Errorf("CheckBalance FAILED. Expect %v, %v; got %v, %v", true, nil, result, err)
	}

	player1.Deposit(100.0)
	result, err = bj.checkBalance(player1.Player)
	if err != nil || result {
		t.Errorf("CheckBalance FAILED. Expect %v, %v; got %v, %v", false, nil, result, err)
	}
}

func TestCheckResult(t *testing.T) {
	bj := Blackjack{}
	bets := Bets{PlayerBetAmount: 10, ComputerBetAmount: 10}

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
	bj.checkResult(&player1.Player, &player2.Player, bets, winner, verifyH1, verifyH2)

	if player1.Balance != 10 || player2.Balance != 10 {
		t.Errorf("CheckResult FAILED. Expected %v, %v; got %v, %v", 10, 10, player1.Balance, player2.Balance)
	}

	player1.ClearHand()
	player2.ClearHand()

	player1.AddCardToHand(cardA)
	player1.AddCardToHand(cardFive)

	player2.AddCardToHand(cardTen)
	player2.AddCardToHand(cardValet)

	winner, verifyH1, verifyH2 = CheckWinner(player1.Hand, player2.Hand)
	bj.checkResult(&player1.Player, &player2.Player, bets, winner, verifyH1, verifyH2)

	if player1.Balance != 10 || player2.Balance != 30 {
		t.Errorf("CheckResult FAILED. Expected %v, %v; got %v, %v", 10, 30, player1.Balance, player2.Balance)
	}

	player1.ClearHand()
	player2.ClearHand()

	player1.AddCardToHand(cardTen)
	player1.AddCardToHand(cardValet)

	player2.AddCardToHand(cardA)
	player2.AddCardToHand(cardFive)

	winner, verifyH1, verifyH2 = CheckWinner(player1.Hand, player2.Hand)
	bj.checkResult(&player1.Player, &player2.Player, bets, winner, verifyH1, verifyH2)

	if player1.Balance != 30 || player2.Balance != 30 {
		t.Errorf("CheckResult FAILED. Expected %v, %v; got %v, %v", 30, 30, player1.Balance, player2.Balance)
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
	bj.checkResult(&player1.Player, &player2.Player, bets, winner, verifyH1, verifyH2)

	if player1.Balance != 40 || player2.Balance != 40 {
		t.Errorf("CheckResult FAILED. Expected %v, %v; got %v, %v", 30, 30, player1.Balance, player2.Balance)
	}
}

func TestCheckPlayAgain(t *testing.T) {
	bj := Blackjack{}
	if bj.checkPlayAgain("y") != false {
		t.Errorf("CheckPlayAgain FAILED. Expected false, got true")
	}

	if bj.checkPlayAgain("x") != true {
		t.Errorf("CheckPlayAgain FAILED. Expected true, got false")
	}

	if bj.checkPlayAgain("n") != true {
		t.Errorf("CheckPlayAgain FAILED. Expected true, got false")
	}
}
